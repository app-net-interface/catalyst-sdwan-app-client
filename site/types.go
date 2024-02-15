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

package site

type Site struct {
	DeviceType                 string   `json:"deviceType"`
	SerialNumber               string   `json:"serialNumber"`
	NcsDeviceName              string   `json:"ncsDeviceName"`
	ConfigStatusMessage        string   `json:"configStatusMessage"`
	TemplateApplyLog           []string `json:"templateApplyLog"`
	UUID                       string   `json:"uuid"`
	ManagementSystemIP         string   `json:"managementSystemIP"`
	TemplateStatus             string   `json:"templateStatus"`
	ChasisNumber               string   `json:"chasisNumber"`
	ConfigStatusMessageDetails string   `json:"configStatusMessageDetails"`
	ConfigOperationMode        string   `json:"configOperationMode"`
	DeviceModel                string   `json:"deviceModel"`
	DeviceState                string   `json:"deviceState"`
	Validity                   string   `json:"validity"`
	PlatformFamily             string   `json:"platformFamily"`
	VedgeCertificateState      string   `json:"vedgeCertificateState"`
	VedgeCSR                   string   `json:"vedgeCSR"`
	VedgeCSRCommonName         string   `json:"vedgeCSRCommonName"`
	RootCertHash               string   `json:"rootCertHash"`
	State                      string   `json:"state"`
	GlobalState                string   `json:"globalState"`
	ExpirationDate             string   `json:"expirationDate"`
	ExpirationDateLong         int64    `json:"expirationDateLong"`
	DeviceIP                   string   `json:"deviceIP"`
	Activity                   []string `json:"activity"`
	CertInstallStatus          string   `json:"certInstallStatus"`
	Personality                string   `json:"personality"`
	ExpirationStatus           string   `json:"expirationStatus"`
	UploadSource               string   `json:"uploadSource"`
	TimeRemainingForExpiration int64    `json:"timeRemainingForExpiration"`
	SubjectSerialNumber        string   `json:"subjectSerialNumber"`
	LocalSystemIP              string   `json:"local-system-ip"`
	SystemIP                   string   `json:"system-ip"`
	ModelSku                   string   `json:"model_sku"`
	SiteID                     string   `json:"site-id"`
	HostName                   string   `json:"host-name"`
	Version                    string   `json:"version"`
	Vbond                      string   `json:"vbond"`
	VmanageConnectionState     string   `json:"vmanageConnectionState"`
	Lastupdated                int64    `json:"lastupdated"`
	Reachability               string   `json:"reachability"`
	UptimeDate                 int64    `json:"uptime-date"`
	DefaultVersion             string   `json:"defaultVersion"`
	AvailableVersions          []string `json:"availableVersions"`
	LifeCycleRequired          bool     `json:"lifeCycleRequired"`
	HardwareCertSerialNumber   string   `json:"hardwareCertSerialNumber"`
	DraftMode                  string   `json:"draftMode"`
}
