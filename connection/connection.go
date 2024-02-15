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

package connection

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

func (c *Client) Create(ctx context.Context, parameters *Parameters) (string, error) {
	if err := common.ValidateCloudType(parameters.CloudType); err != nil {
		return "", err
	}
	mapURL := fmt.Sprintf("/dataservice/multicloud/map?cloudType=%s", parameters.CloudType)
	response, err := c.client.PostRequestWithResponse(ctx, mapURL, parameters)
	if err != nil {
		return "", err
	}
	return common.GetIDFromResponse(response)
}

func (c *Client) GetStatus(ctx context.Context, cloudType string) ([]*Status, error) {
	if err := common.ValidateCloudType(cloudType); err != nil {
		return nil, err
	}
	mapURL := fmt.Sprintf("/dataservice/multicloud/map/status?cloudType=%s", cloudType)
	response, err := c.client.GetRequest(ctx, mapURL)
	if err != nil {
		return nil, err
	}
	var status common.ResponseData[[]*Status]
	if err := json.Unmarshal(response, &status); err != nil {
		return nil, err
	}

	return status.Data, nil
}
