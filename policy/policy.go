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

package policy

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
)

// Client provides all policy methods.
type Client struct {
	client *client.Client
}

// NewClient returns a new policy Client.
func NewClient(c *client.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) Get(ctx context.Context, policyID string) (*Input, error) {
	policyURL := fmt.Sprintf("/dataservice/template/policy/vedge/definition/%s", policyID)
	response, err := c.client.GetRequest(ctx, policyURL)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not get policy: %v", err))
	}
	var data Input
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) Delete(ctx context.Context, policyID string) error {
	policyURL := fmt.Sprintf("/dataservice/template/policy/vedge/%s", policyID)
	_, err := c.client.DeleteRequest(ctx, policyURL)
	if err != nil {
		return common.OperationError(fmt.Sprintf("could not get policy: %v", err))
	}
	return nil
}

func (c *Client) List(ctx context.Context) ([]*Policy, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/policy/vedge")
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not list policy: %v", err))
	}
	var data common.ResponseData[[]*Policy]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) Create(ctx context.Context, data *Input) (string, error) {
	response, err := c.client.PostRequestWithResponse(ctx, "/dataservice/template/policy/vedge", data)
	if err != nil {
		return "", common.OperationError(fmt.Sprintf("could not create policy: %v", err))
	}
	return common.GetCustomFieldFromResponse(response, "policyId")
}
