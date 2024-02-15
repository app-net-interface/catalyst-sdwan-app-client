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

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
)

const (
	statusInProgress = "in_progress"
	statusFailure    = "failure"
	statusScheduled  = "scheduled"
)

type Client struct {
	client           *client.Client
	longPollDuration time.Duration
	longPollRetries  int
}

func NewClient(c *client.Client, longPollDuration time.Duration, longPollRetries int) *Client {
	return &Client{
		client:           c,
		longPollDuration: longPollDuration,
		longPollRetries:  longPollRetries,
	}
}

func (c *Client) ActionStatusLongPoll(ctx context.Context, id string) error {
	url := fmt.Sprintf("/dataservice/device/action/status/%s", id)
	for i := 0; i < c.longPollRetries; i++ {
		response, err := c.client.GetRequest(ctx, url)
		if err != nil {
			return err
		}
		var status Status
		if err := json.Unmarshal(response, &status); err != nil {
			return err
		}
		if len(status.Data) > 0 {
			if err := c.printStatusAndCheckErrors(status.Data[len(status.Data)-1]); err != nil {
				return err
			}
		} else {
			if err := c.printStatusAndCheckErrors(status.Validation); err != nil {
				return err
			}
		}
		if !statusIsInProgress(status) {
			return nil
		}
		time.Sleep(c.longPollDuration)
	}
	return common.OperationError(fmt.Sprintf("status still in progress after %d retries", c.longPollRetries))
}

func (c *Client) printStatusAndCheckErrors(status Data) error {
	c.client.Logger.Infof("[%s] %s", time.Now().Format("2006-01-02 15:04:05"), status.CurrentActivity)
	fmt.Printf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), status.CurrentActivity)
	if status.StatusID == statusFailure {
		if strings.Contains(status.CurrentActivity, "operation is already in progress") {
			return &OperationInProgressError{}
		}
		return common.OperationError(fmt.Sprintf("status long poll failed: %s", status.CurrentActivity))
	}
	return nil
}

func statusIsInProgress(status Status) bool {
	s := status.Summary.Status
	return s == statusInProgress || s == statusScheduled
}
