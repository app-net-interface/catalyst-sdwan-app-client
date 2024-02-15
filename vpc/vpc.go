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

package vpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
)

const (
	hostVPCPath     = "/dataservice/multicloud/hostvpc"
	hostVPCTagsPath = hostVPCPath + "/tags"
)

type Client struct {
	client *client.Client
}

func NewClient(c *client.Client) *Client {
	return &Client{
		client: c,
	}
}

func (c *Client) Get(ctx context.Context, cloudType, vpcID string) (*VPC, error) {
	vpcs, err := c.List(ctx, cloudType)
	if err != nil {
		return nil, err
	}
	for _, vpc := range vpcs {
		if vpc.HostVPCID == vpcID {
			return vpc, nil
		}
	}
	return nil, &common.NotFoundError{}
}

func (c *Client) GetByName(ctx context.Context, cloudType, vpcName string) (*VPC, error) {
	vpcs, err := c.List(ctx, cloudType)
	if err != nil {
		return nil, err
	}
	for _, vpc := range vpcs {
		if vpc.HostVPCName == vpcName {
			return vpc, nil
		}
	}
	return nil, &common.NotFoundError{}
}

func (c *Client) ListAll(ctx context.Context) ([]*VPC, error) {
	allVPCs := make([]*VPC, 0)
	tmpSupportedClouds := []string{"GCP", "AWS"} // for now only those are supported
	// Temporary fix for unstable vManageAPI
	time.Sleep(2 * time.Second)
	for i, cloudType := range tmpSupportedClouds { // should be for _, cloudType := range common.SupportedClouds {
		vpcs, err := c.ListWithParameters(ctx, &ListVPCParameters{CloudType: cloudType})
		if err != nil {
			return nil, err
		}
		allVPCs = append(allVPCs, vpcs...)
		// Temporary fix for unstable vManageAPI, sleep before each new call
		if i != len(tmpSupportedClouds)-1 {
			time.Sleep(3 * time.Second)
		}
	}
	return allVPCs, nil
}

func (c *Client) List(ctx context.Context, cloudType string) ([]*VPC, error) {
	return c.ListWithParameters(ctx, &ListVPCParameters{CloudType: cloudType})
}

func (c *Client) ListWithParameters(ctx context.Context, parameters *ListVPCParameters) ([]*VPC, error) {
	if err := common.ValidateCloudType(parameters.CloudType); err != nil {
		return nil, err
	}
	vpcURL, err := url.Parse(hostVPCPath)
	if err != nil {
		return nil, err
	}
	common.AddParametersToURL(vpcURL, parameters)
	response, err := c.client.GetRequest(ctx, vpcURL.String())
	if err != nil {
		return nil, err
	}
	var data common.ResponseData[[]*VPC]
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data.Data, nil
}

func (c *Client) ListWithTag(ctx context.Context, cloudType, tag string) ([]*VPC, error) {
	return c.ListWithTagWithParameters(ctx, &ListVPCTagParameters{CloudType: cloudType, TagName: tag})
}

func (c *Client) ListWithTagWithParameters(ctx context.Context, parameters *ListVPCTagParameters) ([]*VPC, error) {
	if err := common.ValidateCloudType(parameters.CloudType); err != nil {
		return nil, err
	}
	vpcURL, err := url.Parse(hostVPCTagsPath)
	if err != nil {
		return nil, err
	}
	common.AddParametersToURL(vpcURL, parameters)
	response, err := c.client.GetRequest(ctx, vpcURL.String())
	if err != nil {
		return nil, err
	}
	var data []*VPC
	if err := json.Unmarshal(response, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Client) CreateVPCTag(ctx context.Context, tagName string, vpc *VPC) (string, error) {
	if err := common.ValidateCloudType(vpc.CloudType); err != nil {
		return "", err
	}
	vpcURL := fmt.Sprintf("%s?cloudType=%s", hostVPCTagsPath, vpc.CloudType)
	data := CreateTagData{
		TagName: tagName,
		HostVPCs: []HostVPC{
			{
				CommonData: CommonData{
					CloudType:   vpc.CloudType,
					AccountID:   vpc.AccountID,
					AccountName: vpc.AccountName,
					Region:      vpc.Region,
				},
				ID:          vpc.HostVPCID,
				Label:       vpc.HostVPCID,
				HostVPCID:   vpc.HostVPCID,
				HostVPCName: vpc.HostVPCName,
			},
		},
	}
	response, err := c.client.PostRequestWithResponse(ctx, vpcURL, data)
	if err != nil {
		return "", err
	}
	return common.GetIDFromResponse(response)
}

func (c *Client) DeleteVPCTag(ctx context.Context, cloudType, tag string) (string, error) {
	if err := common.ValidateCloudType(cloudType); err != nil {
		return "", err
	}
	vpcURL := fmt.Sprintf("%s/%s?cloudType=%s", hostVPCTagsPath, tag, cloudType)
	response, err := c.client.DeleteRequest(ctx, vpcURL)
	if err != nil {
		return "", err
	}
	return common.GetIDFromResponse(response)
}
