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

type ListVPCParameters struct {
	CloudType  string `json:"cloudType"`
	AccountIDs string `json:"accountIds"`
	Region     string `json:"region"`
	Untagged   string `json:"untagged"`
}

type ListVPCTagParameters struct {
	CloudType string `json:"cloudType"`
	Region    string `json:"region"`
	TagName   string `json:"tagName"`
}

type CommonData struct {
	AccountID   string `json:"accountId"`
	AccountName string `json:"accountName"`
	CloudType   string `json:"cloudType"`
	Region      string `json:"region"`
}

type VPC struct {
	CommonData
	HostVPCID       string `json:"hostVPCId"`
	HostVPCName     string `json:"hostVPCName"`
	Tag             string `json:"tag"`
	InterconnectTag string `json:"interconnectTag"`
}

type HostVPC struct {
	CommonData
	HostVPCID   string `json:"hostVpcId"`
	HostVPCName string `json:"hostVpcName"`
	ID          string `json:"id"`
	Label       string `json:"label"`
}

type CreateTagData struct {
	TagName  string    `json:"tagName"`
	HostVPCs []HostVPC `json:"hostVpcs"`
}
