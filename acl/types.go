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

package acl

type ACL struct {
	Name           string `json:"name"`
	DefinitionID   string `json:"definitionId"`
	Type           string `json:"type"`
	Description    string `json:"description"`
	Owner          string `json:"owner"`
	LastUpdated    int64  `json:"lastUpdated"`
	InfoTag        string `json:"infoTag"`
	Mode           string `json:"mode"`
	ReferenceCount int    `json:"referenceCount"`
}

type Input struct {
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	Description   string        `json:"description"`
	DefaultAction DefaultAction `json:"defaultAction"`
	Sequences     []Sequence    `json:"sequences"`
}

type DefaultAction struct {
	Type string `json:"type"`
}

type Entry struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type Match struct {
	Entries []Entry `json:"entries"`
}

type Sequence struct {
	SequenceID     int      `json:"sequenceId"`
	SequenceName   string   `json:"sequenceName"`
	BaseAction     string   `json:"baseAction"`
	SequenceType   string   `json:"sequenceType"`
	SequenceIPType string   `json:"sequenceIpType"`
	Match          Match    `json:"match"`
	Actions        []string `json:"actions"`
}
