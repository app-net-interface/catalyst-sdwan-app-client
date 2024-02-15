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

type Parameters struct {
	CloudType string   `json:"cloudType"`
	Matrices  []Matrix `json:"connMatrix"`
}

type Matrix struct {
	SourceType      string `json:"srcType"`
	SourceID        string `json:"srcId"`
	DestinationType string `json:"destType"`
	DestinationID   string `json:"destId"`
	Connection      string `json:"conn"`
}

type Status struct {
	SrcType            string         `json:"srcType"`
	SrcId              string         `json:"srcId"`
	DestType           string         `json:"destType"`
	DestId             string         `json:"destId"`
	Mapped             []StatusMapped `json:"mapped"`
	Unmapped           []StatusMapped `json:"unmapped"`
	OutstandingMapping []StatusMapped `json:"outstandingMapping"`
}

type StatusMapped struct {
	SrcType      string `json:"srcType"`
	SrcId        string `json:"srcId"`
	DestType     string `json:"destType"`
	DestId       string `json:"destId"`
	TunnelId     string `json:"tunnelId"`
	DestRegion   string `json:"destRegion"`
	DestTag      string `json:"destTag,omitempty"`
	SourceTag    string `json:"sourceTag,omitempty"`
	SourceRegion string `json:"sourceRegion"`
	CloudType    string `json:"cloudType"`
}
