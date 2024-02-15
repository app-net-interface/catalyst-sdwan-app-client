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
	"github.com/app-net-interface/catalyst-sdwan-app-client/acl"
)

const (
	sourceIPField        = "sourceIp"
	destinationIPField   = "destinationIp"
	protocolField        = "protocol"
	sourcePortField      = "sourcePort"
	destinationPortField = "destinationPort"
)

type Ports []string
type ProtocolsAndPorts map[string]Ports

type Manager struct {
	sequenceID int
	sequences  []acl.Sequence
}

func NewManager() *Manager {
	return &Manager{
		sequenceID: 1,
	}
}

func (m *Manager) Build() []acl.Sequence {
	return m.sequences
}

func (m *Manager) AddSequencesForIP(name, ip string, protocols ProtocolsAndPorts) {
	m.AddSequencesForCIDR(name, ip+"/32", protocols)
}
func (m *Manager) AddSequencesForCIDR(name, cidr string, protocols ProtocolsAndPorts) {
	if len(protocols) == 0 {
		protocols = ProtocolsAndPorts{"": Ports{""}}
	}
	for protocol, ports := range protocols {
		if len(ports) == 0 {
			ports = Ports{""}
		}
		for _, port := range ports {
			m.addSequence(name, m.createEntriesForSource(port, protocol, cidr))
			m.addSequence(name, m.createEntriesForDestination(port, protocol, cidr))
		}
	}
}

func (m *Manager) createEntriesForSource(port, protocol, cidr string) []acl.Entry {
	return m.createEntriesForField(port, protocol, cidr, sourcePortField, sourceIPField)
}

func (m *Manager) createEntriesForDestination(port, protocol, cidr string) []acl.Entry {
	return m.createEntriesForField(port, protocol, cidr, destinationPortField, destinationIPField)
}

func (m *Manager) createEntriesForField(port, protocol, cidr, portField, cidrField string) []acl.Entry {
	entries := []acl.Entry{
		{
			Field: cidrField,
			Value: cidr,
		},
	}
	if protocol != "" {
		entries = append(entries, acl.Entry{Field: protocolField, Value: protocolNameToNumber(protocol)})
	}
	if port != "" {
		entries = append(entries, acl.Entry{Field: portField, Value: port})
	}
	return entries
}

func (m *Manager) addSequence(name string, entries []acl.Entry) {
	m.sequences = append(m.sequences, createSequence(m.sequenceID, name, entries))
	m.sequenceID++
}

func createSequence(sequenceID int, name string, entries []acl.Entry) acl.Sequence {
	return acl.Sequence{
		SequenceName:   name,
		BaseAction:     "accept",
		SequenceType:   "acl",
		SequenceIPType: "ipv4",
		SequenceID:     sequenceID,
		Match: acl.Match{
			Entries: entries,
		},
		Actions: []string{},
	}
}
