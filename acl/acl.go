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

package acl

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
)

// Client provides all ACL methods.
type Client struct {
	client *client.Client
}

// NewClient returns a new Client.
func NewClient(c *client.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) List(ctx context.Context) ([]*ACL, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/policy/definition/acl")
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not list acl: %v", err))
	}
	var data common.ResponseData[[]*ACL]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) Delete(ctx context.Context, aclID string) error {
	aclURL := fmt.Sprintf("/dataservice/template/policy/definition/acl/%s", aclID)
	if _, err := c.client.DeleteRequest(ctx, aclURL); err != nil {
		return common.OperationError(fmt.Sprintf("could not delete ACL: %v", err))
	}
	return nil
}

func (c *Client) Create(ctx context.Context, data *Input) (string, error) {
	response, err := c.client.PostRequestWithResponse(ctx, "/dataservice/template/policy/definition/acl", data)
	if err != nil {
		return "", common.OperationError(fmt.Sprintf("could not create acl: %v", err))
	}
	return common.GetCustomFieldFromResponse(response, "definitionId")
}

func (c *Client) Get(ctx context.Context, aclID string) (*ACL, error) {
	aclURL := fmt.Sprintf("/dataservice/template/policy/definition/acl/%s", aclID)
	response, err := c.client.GetRequest(ctx, aclURL)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not delete ACL: %v", err))
	}
	var data common.ResponseData[ACL]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return &data.Data, nil
}

func (c *Client) GetByName(ctx context.Context, aclName string) (*ACL, error) {
	acls, err := c.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, acl := range acls {
		if acl.Name == aclName {
			return acl, nil
		}
	}
	return nil, &common.NotFoundError{}
}

func (c *Client) Update(ctx context.Context, aclID string, data *Input) (*common.UpdateResponse, error) {
	aclURL := fmt.Sprintf("/dataservice/template/policy/definition/acl/%s", aclID)
	response, err := c.client.PutRequest(ctx, aclURL, data)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not create acl: %v", err))
	}
	var updateResponse common.UpdateResponse
	if err := json.Unmarshal(response, &updateResponse); err != nil {
		return nil, err
	}
	return &updateResponse, nil
}
