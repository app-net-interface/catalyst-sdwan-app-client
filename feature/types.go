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

package feature

type TemplateInput struct {
	TemplateID           string   `json:"templateId"`
	TemplateName         string   `json:"templateName"`
	TemplateDescription  string   `json:"templateDescription"`
	TemplateType         string   `json:"templateType"`
	DeviceType           []string `json:"deviceType"`
	LastUpdatedBy        string   `json:"lastUpdatedBy"`
	LastUpdatedOn        int64    `json:"lastUpdatedOn"`
	FactoryDefault       bool     `json:"factoryDefault"`
	DevicesAttached      int      `json:"devicesAttached"`
	AttachedMastersCount int      `json:"attachedMastersCount"`
	TemplateMinVersion   string   `json:"templateMinVersion"`
	ConfigType           string   `json:"configType"`
	CreatedBy            string   `json:"createdBy"`
	CreatedOn            int64    `json:"createdOn"`
	ResourceGroup        string   `json:"resourceGroup"`
	TemplateDefinition   string   `json:"templateDefinition"`
}

type VipValue[T any] struct {
	VipObjectType   string `json:"vipObjectType"`
	VipType         string `json:"vipType"`
	VipValue        T      `json:"vipValue,omitempty"`
	VipVariableName string `json:"vipVariableName,omitempty"`
}

type VipPrimary[T any] struct {
	VipType       string   `json:"vipType"`
	VipValue      []T      `json:"vipValue"`
	VipObjectType string   `json:"vipObjectType"`
	VipPrimaryKey []string `json:"vipPrimaryKey"`
}

type AccessListValue struct {
	ACLName       VipValue[string] `json:"acl-name"`
	Direction     VipValue[string] `json:"direction"`
	PriorityOrder []string         `json:"priority-order"`
}

type TDArp struct {
	IP VipPrimary[interface{}] `json:"ip"`
}

type TDIP struct {
	DHCPClient       VipValue[string]        `json:"dhcp-client"`
	DHCPDistance     VipValue[int]           `json:"dhcp-distance"`
	SecondaryAddress VipPrimary[interface{}] `json:"secondary-address"`
}

type VipValueNonEmpty[T any] struct {
	VipObjectType   string `json:"vipObjectType"`
	VipType         string `json:"vipType"`
	VipValue        T      `json:"vipValue"`
	VipVariableName string `json:"vipVariableName,omitempty"`
}

type TDIPv6 struct {
	AccessList       VipPrimary[interface{}]  `json:"access-list"`
	Address          VipValueNonEmpty[string] `json:"address"`
	DhcpHelperV6     VipPrimary[interface{}]  `json:"dhcp-helper-v6"`
	SecondaryAddress VipPrimary[interface{}]  `json:"secondary-address"`
}

type TDRewriteRule struct {
	RuleName VipValue[interface{}] `json:"rule-name"`
}

type TDTLOCExtensionGreFrom struct {
	SourceIP VipValue[interface{}] `json:"src-ip"`
	XConnect VipValue[interface{}] `json:"xconnect"`
}

type TDTrustSecEnforcement struct {
	Enable  VipValue[interface{}] `json:"enable"`
	Segment VipValue[interface{}] `json:"sgt"`
}

type TDTrustSec struct {
	Enforcement TDTrustSecEnforcement `json:"enforcement"`
}

type TemplateDefinition struct {
	IfName               VipValue[string]            `json:"if-name"`
	Description          VipValue[interface{}]       `json:"description"`
	IP                   TDIP                        `json:"ip"`
	DHCPHelper           VipValue[interface{}]       `json:"dhcp-helper"`
	FlowControl          interface{}                 `json:"flow-control"`
	ClearDontFragment    interface{}                 `json:"clear-dont-fragment"`
	PMTU                 interface{}                 `json:"pmtu"`
	MTU                  VipValue[int]               `json:"mtu"`
	StaticIngressQOS     interface{}                 `json:"static-ingress-qos"`
	TCPMssAdjust         VipValue[interface{}]       `json:"tcp-mss-adjust"`
	MACAddress           VipValue[interface{}]       `json:"mac-address"`
	Speed                VipValue[string]            `json:"speed"`
	Duplex               VipValue[string]            `json:"duplex"`
	Shutdown             VipValue[string]            `json:"shutdown"`
	ARPTimeout           VipValue[int]               `json:"arp-timeout"`
	AutoNegotiate        VipValue[interface{}]       `json:"autonegotiate"`
	ShapingRate          VipValue[interface{}]       `json:"shaping-rate"`
	QOSMap               VipValue[interface{}]       `json:"qos-map"`
	QOSMapVPN            VipValue[interface{}]       `json:"qos-map-vpn"`
	Tracker              VipValue[interface{}]       `json:"tracker"`
	BandwidthUpstream    VipValue[interface{}]       `json:"bandwidth-upstream"`
	BandwidthDownstream  VipValue[interface{}]       `json:"bandwidth-downstream"`
	BlockNonSourceIP     VipValue[string]            `json:"block-non-source-ip"`
	RewriteRule          TDRewriteRule               `json:"rewrite-rule"`
	TLOCExtension        VipValue[interface{}]       `json:"tloc-extension"`
	LoadInterval         VipValue[int]               `json:"load-interval"`
	IcmpRedirectDisable  VipValue[string]            `json:"icmp-redirect-disable"`
	TLOCExtensionGREFrom TDTLOCExtensionGreFrom      `json:"tloc-extension-gre-from"`
	AccessList           VipPrimary[AccessListValue] `json:"access-list"`
	AutoBandwidthDetect  VipValue[string]            `json:"auto-bandwidth-detect"`
	IperfServer          VipValue[interface{}]       `json:"iperf-server"`
	MediaType            VipValue[string]            `json:"media-type"`
	InterfaceMTU         VipValue[int]               `json:"intrf-mtu"`
	IPDirectedBroadcast  VipValue[string]            `json:"ip-directed-broadcast"`
	TrustSec             TDTrustSec                  `json:"trustsec"`
	IPv6                 TDIPv6                      `json:"ipv6"`
	Arp                  TDArp                       `json:"arp"`
	Vrrp                 VipPrimary[interface{}]     `json:"vrrp"`
	IPv6Vrrp             VipPrimary[interface{}]     `json:"ipv6-vrrp"`
}

type Template struct {
	TemplateName        string             `json:"templateName"`
	TemplateDescription string             `json:"templateDescription"`
	TemplateType        string             `json:"templateType"`
	DeviceType          []string           `json:"deviceType"`
	TemplateMinVersion  string             `json:"templateMinVersion"`
	TemplateDefinition  TemplateDefinition `json:"templateDefinition"`
	FactoryDefault      bool               `json:"factoryDefault"`
}
