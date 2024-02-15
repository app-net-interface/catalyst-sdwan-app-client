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

package policy

type Policy struct {
	LastUpdatedBy        string `json:"lastUpdatedBy"`
	MastersAttached      int    `json:"mastersAttached"`
	PolicyName           string `json:"policyName"`
	PolicyDefinition     string `json:"policyDefinition"`
	PolicyDefinitionEdit string `json:"policyDefinitionEdit"`
	CreatedOn            int64  `json:"createdOn"`
	PolicyDescription    string `json:"policyDescription"`
	RID                  int    `json:"@rid"`
	PolicyID             string `json:"policyId"`
	CreatedBy            string `json:"createdBy"`
	DevicesAttached      int    `json:"devicesAttached"`
	PolicyType           string `json:"policyType"`
	LastUpdatedOn        int64  `json:"lastUpdatedOn"`
}

type Input struct {
	PolicyDescription string     `json:"policyDescription"`
	PolicyType        string     `json:"policyType"`
	PolicyName        string     `json:"policyName"`
	PolicyDefinition  Definition `json:"policyDefinition"`
	IsPolicyActivated bool       `json:"isPolicyActivated"`
}

type Definition struct {
	Assemblies []Assembly `json:"assembly"`
	Settings   Setting    `json:"settings"`
}

type Assembly struct {
	DefinitionID string `json:"definitionId"`
	Type         string `json:"type"`
}

type Setting struct {
	ApplicationVisibility bool `json:"appVisibility"`
	FlowVisibility        bool `json:"flowVisibility"`
}
