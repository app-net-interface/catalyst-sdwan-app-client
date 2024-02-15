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

package urlallowlist

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
)

type Client struct {
	client *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) List(ctx context.Context) ([]*URLAllowlistListResp, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/policy/list/urlwhitelist")
	if err != nil {
		return nil, err
	}
	var data common.ResponseData[[]*URLAllowlistListResp]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) Get(ctx context.Context, id string) (*URLAllowlist, error) {
	response, err := c.client.GetRequest(ctx, fmt.Sprintf("/dataservice/template/policy/list/urlwhitelist/%s", id))
	if err != nil {
		return nil, err
	}
	var data *URLAllowlist
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetByName(ctx context.Context, name string) (*URLAllowlist, error) {
	list, err := c.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		if u.Name == name {
			return c.Get(ctx, u.ListId)
		}
	}
	return nil, fmt.Errorf("couldn't find URLAllowlist with name %s", name)
}

func (c *Client) Create(ctx context.Context, urlAllowlist *URLAllowlistInput) (definitionId string, err error) {
	response, err := c.client.PostRequestWithResponse(ctx, "/dataservice/template/policy/list/urlwhitelist", urlAllowlist)
	if err != nil {
		return "", common.OperationError(fmt.Sprintf("could not create url Allowlist: %v", err))
	}
	return common.GetCustomFieldFromResponse(response, "listId")
}

func (c *Client) Delete(ctx context.Context, urlAllowlistId string) error {
	urlAllowlistID := fmt.Sprintf("/dataservice/template/policy/list/urlwhitelist/%s", urlAllowlistId)
	_, err := c.client.DeleteRequest(ctx, urlAllowlistID)
	if err != nil {
		return common.OperationError(fmt.Sprintf("could not delete URLAllowlist: %v", err))
	}
	return nil
}

func (c *Client) DeleteByName(ctx context.Context, urlAllowlistName string) error {
	list, err := c.List(ctx)
	if err != nil {
		return err
	}
	for _, u := range list {
		if u.Name == urlAllowlistName {
			return c.Delete(ctx, u.ListId)
		}
	}
	return fmt.Errorf("couldn't find URLAllowlist with name %s", urlAllowlistName)
}
