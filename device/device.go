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

package device

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
	"github.com/app-net-interface/catalyst-sdwan-app-client/site"
)

type Client struct {
	client *client.Client
	site   *site.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{
		client: c,
		site:   site.NewClient(c),
	}
}

func (c *Client) Get(ctx context.Context, templateID string) (*Template, error) {
	templateURL := fmt.Sprintf("/dataservice/template/device/object/%s", templateID)
	response, err := c.client.GetRequest(ctx, templateURL)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not get template: %v", err))
	}
	var data Template
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) List(ctx context.Context) ([]*Template, error) {
	templateURL := "/dataservice/template/device/"
	response, err := c.client.GetRequest(ctx, templateURL)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not get template: %v", err))
	}
	var data []*Template
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) Update(ctx context.Context, template *Template) (*UpdateResponse, error) {
	templateURL := fmt.Sprintf("/dataservice/template/device/%s", template.TemplateID)
	response, err := c.client.PutRequest(ctx, templateURL, template)
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not create template: %v", err))
	}
	var data common.ResponseData[UpdateResponse]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return &data.Data, nil
}

func (c *Client) GetAttachedDevices(ctx context.Context, templateID string) ([]*AttachedDevice, error) {
	attachedURL := "/dataservice/template/device/config/attached/" + templateID
	response, err := c.client.GetRequest(ctx, attachedURL)
	if err != nil {
		return nil, err
	}
	var data common.ResponseData[[]*AttachedDevice]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) PushConfiguration(
	ctx context.Context,
	attachedDevices []*AttachedDevice,
	templateID string,
) (string, error) {
	deviceIDs := make([]string, 0, len(attachedDevices))
	for _, attachedDevice := range attachedDevices {
		deviceIDs = append(deviceIDs, attachedDevice.UUID)
	}
	pushData := push{
		TemplateID:     templateID,
		DeviceIDs:      deviceIDs,
		IsEdited:       true,
		IsMasterEdited: true,
	}
	templateURL := "/dataservice/template/device/config/input/"
	response, err := c.client.PostRequestWithResponse(ctx, templateURL, pushData)
	if err != nil {
		return "", err
	}
	var data common.ResponseData[[]map[string]interface{}]
	if err := json.Unmarshal(response, &data); err != nil {
		return "", err
	}

	duplicateIPURL := "/dataservice/template/device/config/duplicateip"
	duplicatePayload := map[string]interface{}{"device": data.Data}
	response, err = c.client.PostRequestWithResponse(ctx, duplicateIPURL, duplicatePayload)
	if err != nil {
		return "", err
	}
	var duplicateData common.ResponseData[[]map[string]interface{}]
	if err := json.Unmarshal(response, &duplicateData); err != nil {
		return "", err
	}
	if len(duplicateData.Data) > 0 {
		return "", common.OperationError(fmt.Sprintf("duplicate IP found: %v", duplicateData.Data))
	}

	for i := range data.Data {
		data.Data[i]["csv-templateId"] = templateID
	}
	attach := attach{
		DeviceTemplateList: []deviceTemplateList{
			{
				IsMasterEdited: false,
				IsEdited:       true,
				Device:         data.Data,
				TemplateID:     templateID,
			},
		},
	}

	url := "/dataservice/template/device/config/attachfeature"
	response, err = c.client.PostRequestWithResponse(ctx, url, attach)
	if err != nil {
		return "", err
	}
	return common.GetIDFromResponse(response)
}

func (c *Client) GetFromSite(ctx context.Context, siteID string) (*Template, error) {
	s, err := c.site.Get(ctx, siteID)
	if err != nil {
		return nil, err
	}
	for _, log := range s.TemplateApplyLog {
		substr := "Configuring device with feature template: "
		if strings.Contains(log, substr) {
			templateName := log[strings.LastIndex(log, substr)+len(substr):]
			return c.GetByName(ctx, templateName)
		}
	}

	return nil, &common.NotFoundError{}
}

func (c *Client) GetByName(ctx context.Context, templateName string) (*Template, error) {
	response, err := c.client.GetRequest(ctx, "/dataservice/template/device/")
	if err != nil {
		return nil, common.OperationError(fmt.Sprintf("could not get template: %v", err))
	}
	var data common.ResponseData[[]Template]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	for _, data := range data.Data {
		if data.TemplateName == templateName {
			return c.Get(ctx, data.TemplateID)
		}
	}
	return nil, &common.NotFoundError{}
}
