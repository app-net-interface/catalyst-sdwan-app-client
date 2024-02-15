// Copyright (c) 2023 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
)

type LoginError struct{}

func (*LoginError) Error() string {
	return "could not fetch data. Please login again"
}

type Client struct {
	Token  string
	Logger *log.Logger
	url    string
	client *http.Client
}

func NewDefaultClient(url string) (*Client, error) {
	transport := &http.Transport{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Transport: transport,
		Jar:       jar,
	}
	logger := log.New()
	logger.Level = log.InfoLevel
	return NewClient(url, client, logger), nil
}

func NewClient(url string, httpClient *http.Client, logger *log.Logger) *Client {
	return &Client{
		url:    url,
		client: httpClient,
		Logger: logger,
	}
}

func (c *Client) GetRequest(ctx context.Context, url string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url+url, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(ctx, request)
}

func (c *Client) DeleteRequest(ctx context.Context, url string) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodDelete, c.url+url, nil)
	if err != nil {
		return nil, err
	}
	return c.doRequest(ctx, request)
}

func (c *Client) PostRequestWithResponse(ctx context.Context, url string, data any) ([]byte, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.Logger.Debugf("body: %s", string(payload))
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, c.url+url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	return c.doRequest(ctx, request)
}

func (c *Client) PutRequest(ctx context.Context, url string, data any) ([]byte, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.Logger.Debugf("body: %s", string(payload))
	request, err := http.NewRequestWithContext(ctx, http.MethodPut, c.url+url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	return c.doRequest(ctx, request)
}

func (c *Client) doRequest(ctx context.Context, request *http.Request) ([]byte, error) {
	c.Logger.Debugf("%s %s", request.Method, request.URL)
	request.Header.Set("X-XSRF-TOKEN", c.Token)
	request.Header.Set("Content-Type", "application/json")
	response, err := c.client.Do(request.WithContext(ctx))
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			c.Logger.Infof("failed to close response body: %v", err)
		}
	}()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	strBody := string(body)
	c.Logger.Debugf("response: %+v", strBody)
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		c.Logger.Infof("response: %+v", strBody)
		return nil, common.OperationError(fmt.Sprintf("invalid status code: %d", response.StatusCode))
	}
	if strings.Contains(strBody, "<html>") {
		return nil, &LoginError{}
	}
	if strings.Contains(strBody, "error") {
		return nil, common.OperationError(strBody)
	}
	return body, nil
}
