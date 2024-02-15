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

package example

import (
	"github.com/app-net-interface/catalyst-sdwan-app-client/connection"
	"github.com/app-net-interface/catalyst-sdwan-app-client/mocks"
	"github.com/app-net-interface/catalyst-sdwan-app-client/site"
	"github.com/app-net-interface/catalyst-sdwan-app-client/vpc"
	"github.com/app-net-interface/catalyst-sdwan-app-client/vpn"
	"github.com/stretchr/testify/mock"
)

func NewFakeClient() *mocks.Client {
	client := &mocks.Client{}
	client.On("Login").Return(nil)
	client.On("GetToken").Return("token")
	client.On("SetToken").Return()
	client.On("ACL").Return(&mocks.ACL{})
	client.On("Connection").Return(Connection())
	client.On("Device").Return(&mocks.Device{})
	client.On("Feature").Return(&mocks.Feature{})
	client.On("VPN").Return(VPN())
	client.On("VPC").Return(VPC())
	client.On("Status").Return(&mocks.Status{})
	client.On("Site").Return(Site())
	client.On("Policy").Return(&mocks.Policy{})
	return client
}

func Connection() *mocks.Connection {
	conn := &mocks.Connection{}
	conn.On("GetStatus", mock.Anything, mock.Anything).Return([]*connection.Status{}, nil)
	return conn
}

func VPN() *mocks.VPN {
	vpnExample := &mocks.VPN{}
	vpnExample.On("List", mock.Anything, mock.Anything).Return([]*vpn.Data{
		{
			SegmentID:   "10",
			SegmentName: "vpn10",
			ID:          "10",
		},
	}, nil)
	return vpnExample
}

func VPC() *mocks.VPC {
	vpcExample := &mocks.VPC{}
	vpcExample.On("ListWithParameters", mock.Anything, mock.Anything).Return([]*vpc.VPC{
		{
			CommonData: vpc.CommonData{
				Region:      "us-east-1",
				AccountName: "awi",
				CloudType:   "AWS",
			},
			HostVPCID:       "vpc-043444b8201ca823d",
			HostVPCName:     "development",
			Tag:             "development",
			InterconnectTag: "development",
		},
		{
			CommonData: vpc.CommonData{
				Region:      "us-east-1",
				AccountName: "awi",
				CloudType:   "AWS",
			},
			HostVPCID:       "vpc-08532095b01548260",
			HostVPCName:     "staging",
			Tag:             "staging",
			InterconnectTag: "staging",
		},
	}, nil)
	return vpcExample
}

func Site() *mocks.Site {
	siteExample := &mocks.Site{}
	siteExample.On("List", mock.Anything).Return([]*site.Site{
		{
			SiteID:        "201",
			ChasisNumber:  "201",
			DeviceIP:      "10.0.100.1",
			NcsDeviceName: "Site 201",
		},
	}, nil)
	return siteExample
}
