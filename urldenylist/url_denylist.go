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

package urldenylist

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

func (c *Client) List(ctx context.Context) ([]*URLDenylistListResp, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/policy/list/urlblacklist")
	if err != nil {
		return nil, err
	}
	var data common.ResponseData[[]*URLDenylistListResp]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) Get(ctx context.Context, id string) (*URLDenylist, error) {
	response, err := c.client.GetRequest(ctx, fmt.Sprintf("/dataservice/template/policy/list/urlblacklist/%s", id))
	if err != nil {
		return nil, err
	}
	var data *URLDenylist
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) GetByName(ctx context.Context, name string) (*URLDenylist, error) {
	list, err := c.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, u := range list {
		if u.Name == name {
			return c.Get(ctx, u.ListId)
		}
	}
	return nil, fmt.Errorf("couldn't find URLDenylist with name %s", name)
}

func (c *Client) Create(ctx context.Context, urlDenylist *URLDenylistInput) (definitionId string, err error) {
	response, err := c.client.PostRequestWithResponse(ctx, "/dataservice/template/policy/list/urlblacklist", urlDenylist)
	if err != nil {
		return "", common.OperationError(fmt.Sprintf("could not create url Denylist: %v", err))
	}
	return common.GetCustomFieldFromResponse(response, "listId")
}

func (c *Client) Delete(ctx context.Context, urlDenylistId string) error {
	urlDenylistID := fmt.Sprintf("/dataservice/template/policy/list/urlblacklist/%s", urlDenylistId)
	_, err := c.client.DeleteRequest(ctx, urlDenylistID)
	if err != nil {
		return common.OperationError(fmt.Sprintf("could not delete URLDenylist: %v", err))
	}
	return nil
}

func (c *Client) DeleteByName(ctx context.Context, urlDenylistName string) error {
	list, err := c.List(ctx)
	if err != nil {
		return err
	}
	for _, u := range list {
		if u.Name == urlDenylistName {
			return c.Delete(ctx, u.ListId)
		}
	}
	return fmt.Errorf("couldn't find URLDenylist with name %s", urlDenylistName)
}
