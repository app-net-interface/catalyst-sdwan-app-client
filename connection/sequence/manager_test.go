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

package sequence

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/app-net-interface/catalyst-sdwan-app-client/acl"
	"github.com/app-net-interface/catalyst-sdwan-app-client/types"
)

func TestManagerCreateSequenceForCIDR(t *testing.T) {
	cidr := "10.10.10.0/24"
	manager := &Manager{}
	expected := []acl.Sequence{
		createSequence(0, manager.createEntriesForSource("1", "icmp", cidr)),
		createSequence(1, manager.createEntriesForDestination("1", "icmp", cidr)),
		createSequence(2, manager.createEntriesForSource("2", "icmp", cidr)),
		createSequence(3, manager.createEntriesForDestination("2", "icmp", cidr)),
		createSequence(4, manager.createEntriesForSource("3", "tcp", cidr)),
		createSequence(5, manager.createEntriesForDestination("3", "tcp", cidr)),
		createSequence(6, manager.createEntriesForSource("4", "tcp", cidr)),
		createSequence(7, manager.createEntriesForDestination("4", "tcp", cidr)),
	}
	protocols := types.ProtocolPorts{
		"icmp": []string{"1", "2"},
		"tcp":  []string{"3", "4"},
	}

	manager.AddSequencesForCIDR(cidr, protocols)

	compareSequences(t, expected, manager.sequences)
}

func TestManagerCreateSequenceForCIDRWithoutPort(t *testing.T) {
	cidr := "10.10.10.0/24"
	protocols := types.ProtocolPorts{
		"icmp": []string{},
		"tcp":  []string{},
	}
	manager := &Manager{}
	expected := []acl.Sequence{
		createSequence(0, manager.createEntriesForSource("", "icmp", cidr)),
		createSequence(1, manager.createEntriesForDestination("", "icmp", cidr)),
		createSequence(2, manager.createEntriesForSource("", "tcp", cidr)),
		createSequence(3, manager.createEntriesForDestination("", "tcp", cidr)),
	}

	manager.AddSequencesForCIDR(cidr, protocols)

	compareSequences(t, expected, manager.sequences)
}

func TestManagerCreateSequenceForCIDRWithoutPortAndProtocol(t *testing.T) {
	cidr := "10.10.10.0/24"
	protocols := types.ProtocolPorts{}
	manager := &Manager{}
	expected := []acl.Sequence{
		createSequence(0, manager.createEntriesForSource("", "", cidr)),
		createSequence(1, manager.createEntriesForDestination("", "", cidr)),
	}

	manager.AddSequencesForCIDR(cidr, protocols)

	compareSequences(t, expected, manager.sequences)
}

// order does not matter
func compareSequences(t *testing.T, expected, actual []acl.Sequence) {
	require.Equal(t, len(expected), len(actual))

	for i := 1; i < len(actual); i++ {
		assert.Less(t, actual[i-1].SequenceID, actual[i].SequenceID)
	}

	for _, expSeq := range expected {
		found := false
		for _, actSeq := range actual {
			if reflect.DeepEqual(expSeq.Match, actSeq.Match) {
				found = true
				break
			}
		}
		require.True(t, found, "Expected sequence %v not found", expSeq)
	}
}
