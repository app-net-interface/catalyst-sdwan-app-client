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

package feature

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

func (c *Client) Create(ctx context.Context, template *TemplateInput) (string, error) {
	response, err := c.client.PostRequestWithResponse(ctx, "/dataservice/template/device/feature/", template)
	if err != nil {
		return "", common.OperationError(fmt.Sprintf("could not create template: %v", err))
	}
	return common.GetCustomFieldFromResponse(response, "templateId")
}

func (c *Client) ListByType(ctx context.Context, templateType string) ([]*TemplateInput, error) {
	allTemplates, err := c.List(ctx)
	if err != nil {
		return nil, err
	}
	var templates []*TemplateInput
	for _, template := range allTemplates {
		if template.TemplateType == templateType && template.DevicesAttached > 0 {
			templates = append(templates, template)
		}
	}
	return templates, nil
}

func (c *Client) List(ctx context.Context) ([]*TemplateInput, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/feature/")
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not create template: %v", err))
	}
	var data common.ResponseData[[]*TemplateInput]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) Get(ctx context.Context, templateID string) (*Template, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/feature/object/"+templateID)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not create template: %v", err))
	}
	var data Template
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) Update(ctx context.Context, templateID string, template *Template) (*common.UpdateResponse, error) {
	response, err := c.client.PutRequest(ctx, "/dataservice/template/feature/"+templateID, template)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not create template: %v", err))
	}
	var data common.UpdateResponse
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
