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

package vpn

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

func (c *Client) List(ctx context.Context, cloudType string) ([]*Data, error) {
	vpnURL := fmt.Sprintf("/dataservice/multicloud/map/vpns?cloudType=%s", cloudType)
	response, err := c.client.GetRequest(ctx, vpnURL)
	if err != nil {
		return nil, err
	}
	var data common.ResponseData[[]*Data]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}
