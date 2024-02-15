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

type Template struct { //nolint:maligned // this order is required by vManage
	TemplateID                   string            `json:"templateId"`
	TemplateName                 string            `json:"templateName"`
	TemplateDescription          string            `json:"templateDescription"`
	DeviceType                   string            `json:"deviceType"`
	DeviceRole                   string            `json:"deviceRole"`
	ConfigType                   string            `json:"configType"`
	FactoryDefault               bool              `json:"factoryDefault"`
	PolicyID                     string            `json:"policyId"`
	FeatureTemplateUIDRange      []string          `json:"featureTemplateUidRange"`
	DraftMode                    interface{}       `json:"draftMode"`
	ConnectionPreferenceRequired bool              `json:"connectionPreferenceRequired"`
	ConnectionPreference         bool              `json:"connectionPreference"`
	GeneralTemplates             []GeneralTemplate `json:"generalTemplates"`
}

type GeneralTemplate struct {
	TemplateID   string        `json:"templateId"`
	TemplateType string        `json:"templateType"`
	SubTemplates []SubTemplate `json:"subTemplates,omitempty"`
}

type SubTemplate struct {
	TemplateID   string `json:"templateId"`
	TemplateType string `json:"templateType"`
}

type AttachedDevice struct {
	HostName         string   `json:"host-name"`
	DeviceIP         string   `json:"deviceIP"`
	LocalSystemIP    string   `json:"local-system-ip"`
	SiteID           string   `json:"site-id"`
	DeviceGroups     []string `json:"device-groups"`
	UUID             string   `json:"uuid"`
	Personality      string   `json:"personality"`
	ConfigCloudXMode string   `json:"configCloudxMode"`
}

type UpdateResponse struct {
	ProcessID       string            `json:"processId"`
	AttachedDevices []*AttachedDevice `json:"attachedDevices"`
}

type push struct {
	TemplateID     string   `json:"templateId"`
	DeviceIDs      []string `json:"deviceIds"`
	IsEdited       bool     `json:"isEdited"`
	IsMasterEdited bool     `json:"isMasterEdited"`
}

type deviceTemplateList struct {
	TemplateID     string                   `json:"templateId"`
	Device         []map[string]interface{} `json:"device"`
	IsEdited       bool                     `json:"isEdited"`
	IsMasterEdited bool                     `json:"isMasterEdited"`
}

type attach struct {
	DeviceTemplateList []deviceTemplateList `json:"deviceTemplateList"`
}
