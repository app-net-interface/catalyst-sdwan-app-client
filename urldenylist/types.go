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

type URLDenylistListResp struct {
	ListId      string  `json:"listId"`
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Description string  `json:"description"`
	Entries     []Entry `json:"entries"`
}

type URLDenylistInput struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Type        string  `json:"type"`
	Entries     []Entry `json:"entries"`
}

type Entry struct {
	Pattern string `json:"pattern"`
}

type URLDenylist struct {
	ListId              string      `json:"listId"`
	Name                string      `json:"name"`
	Type                string      `json:"type"`
	Description         string      `json:"description"`
	Entries             []Entry     `json:"entries"`
	LastUpdated         int64       `json:"lastUpdated"`
	Owner               string      `json:"owner"`
	ReadOnly            bool        `json:"readOnly"`
	Version             string      `json:"version"`
	InfoTag             string      `json:"infoTag"`
	ReferenceCount      int         `json:"referenceCount"`
	References          []Reference `json:"references"`
	IsActivatedByVsmart bool        `json:"isActivatedByVsmart"`
}

type Reference struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}
