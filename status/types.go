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

package status

type Status struct {
	Validation Data    `json:"validation"`
	Data       []Data  `json:"data"`
	Summary    Summary `json:"summary"`
}

type Data struct {
	Activity        []string `json:"activity"`
	StatusID        string   `json:"statusId"`
	CurrentActivity string   `json:"currentActivity"`
	Status          string   `json:"status"`
}

type Summary struct {
	Action    string `json:"action"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Status    string `json:"status"`
}

type OperationInProgressError struct{}

func (*OperationInProgressError) Error() string {
	return "operation in progress. Please retry again"
}
