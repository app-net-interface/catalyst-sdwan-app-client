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

package urlfiltering

type URLFilteringListRespStructure struct {
	Name           string      `json:"name"`
	DefinitionID   string      `json:"definitionId"`
	Type           string      `json:"type"`
	Description    string      `json:"description"`
	Owner          string      `json:"owner"`
	LastUpdated    int64       `json:"lastUpdated"`
	InfoTag        string      `json:"infoTag"`
	Mode           string      `json:"mode"`
	Optimized      string      `json:"optimized"`
	ReferenceCount int         `json:"referenceCount"`
	References     []Reference `json:"references"`
}

type Reference struct {
	ID       string `json:"id"`
	Property string `json:"property"`
}

type URLFilteringDefinition struct {
	Name        string     `json:"name"`
	Type        string     `json:"type"`
	Description string     `json:"description"`
	Definition  Definition `json:"definition"`
}

type Definition struct {
	WebCategoriesAction string   `json:"webCategoriesAction"`
	WebCategories       []string `json:"webCategories"`
	WebReputation       string   `json:"webReputation"`
	UrlAllowlist        UrlRef   `json:"urlWhiteList"`
	UrlDenylist         UrlRef   `json:"urlBlackList"`
	BlockPageAction     string   `json:"blockPageAction"`
	BlockPageContents   string   `json:"blockPageContents"`
	EnableAlerts        bool     `json:"enableAlerts"`
	Alerts              []string `json:"alerts"`
	Logging             []string `json:"logging"`
	TargetVpns          []string `json:"targetVpns"`
}

type UrlRef struct {
	Ref string `json:"ref"`
}
