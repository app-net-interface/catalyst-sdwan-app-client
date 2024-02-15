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
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type AuthenticityToken struct {
	Token string
}

const (
	tokenPath = "/dataservice/client/token?json=true"
)

func (c *Client) Login(ctx context.Context, username, password string) error {
	if err := c.setSessionCookie(ctx, username, password); err != nil {
		return err
	}
	if err := c.setToken(ctx); err != nil {
		return err
	}
	return nil
}

func (c *Client) setSessionCookie(ctx context.Context, username, password string) error {
	data := url.Values{
		"j_username": {username},
		"j_password": {password},
	}
	loginURL := c.url + "/j_security_check"

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, loginURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		err := response.Body.Close()
		if err != nil {
			c.Logger.Infof("failed to close response body: %v", err)
		}
	}()
	return nil
}

func (c *Client) setToken(ctx context.Context) error {
	data, err := c.GetRequest(ctx, tokenPath)
	if err != nil {
		return err
	}
	var token AuthenticityToken
	if err := json.Unmarshal(data, &token); err != nil {
		return err
	}
	c.Token = token.Token
	return nil
}
