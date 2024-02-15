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

package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
)

var SupportedClouds = []string{"AWS", "AWS_GOVCLOUD", "GCP", "AZURE"}

var errOperation = errors.New("operation error")

type ResponseData[T any] struct {
	Data T `json:"data"`
}

type UpdateResponse struct {
	ProcessID               string   `json:"processId"`
	MasterTemplatesAffected []string `json:"masterTemplatesAffected"`
}

type NotFoundError struct{}

func (*NotFoundError) Error() string {
	return "object not found"
}

func GetIDFromResponse(response []byte) (string, error) {
	return GetCustomFieldFromResponse(response, "id")
}

func GetCustomFieldFromResponse(response []byte, fieldName string) (string, error) {
	var data map[string]interface{}
	if err := json.Unmarshal(response, &data); err != nil {
		return "", err
	}
	result, ok := data[fieldName].(string)
	if !ok {
		return "", OperationError("field is not string")
	}
	return result, nil
}

func AddParametersToURL(baseURL *url.URL, parameters any) {
	query := baseURL.Query()
	val := reflect.ValueOf(parameters).Elem()
	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get("json")
		fieldVal := val.Field(i)
		if !fieldVal.IsZero() {
			query.Add(tag, fieldVal.String())
		}
	}
	baseURL.RawQuery = query.Encode()
}

func ValidateCloudType(cloudType string) error {
	for _, expected := range SupportedClouds {
		if cloudType == expected {
			return nil
		}
	}
	return OperationError(fmt.Sprintf("unexpected cloud type. Chose one of: %v", SupportedClouds))
}

func OperationError(op string) error {
	return fmt.Errorf("%w: %s", errOperation, op)
}
