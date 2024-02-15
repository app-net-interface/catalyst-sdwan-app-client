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

package vmanage

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/app-net-interface/catalyst-sdwan-app-client/acl"
	c "github.com/app-net-interface/catalyst-sdwan-app-client/client"
	"github.com/app-net-interface/catalyst-sdwan-app-client/common"
	"github.com/app-net-interface/catalyst-sdwan-app-client/connection"
	"github.com/app-net-interface/catalyst-sdwan-app-client/device"
	"github.com/app-net-interface/catalyst-sdwan-app-client/feature"
	"github.com/app-net-interface/catalyst-sdwan-app-client/policy"
	"github.com/app-net-interface/catalyst-sdwan-app-client/site"
	"github.com/app-net-interface/catalyst-sdwan-app-client/status"
	ual "github.com/app-net-interface/catalyst-sdwan-app-client/urlallowlist"
	udl "github.com/app-net-interface/catalyst-sdwan-app-client/urldenylist"
	"github.com/app-net-interface/catalyst-sdwan-app-client/urlfiltering"
	"github.com/app-net-interface/catalyst-sdwan-app-client/vpc"
	"github.com/app-net-interface/catalyst-sdwan-app-client/vpn"
)

type Client interface {
	Login(ctx context.Context, username, password string) error
	GetToken() string
	SetToken(token string)
	ACL() ACL
	Connection() Connection
	Device() Device
	Feature() Feature
	VPN() VPN
	VPC() VPC
	Status() Status
	Site() Site
	Policy() Policy
	URLFiltering() URLFiltering
	URLDenylist() URLDenylist
	URLAllowlist() URLAllowlist
}

type ACL interface {
	List(ctx context.Context) ([]*acl.ACL, error)
	Delete(ctx context.Context, aclID string) error
	Create(ctx context.Context, data *acl.Input) (string, error)
	Get(ctx context.Context, aclID string) (*acl.ACL, error)
	GetByName(ctx context.Context, aclName string) (*acl.ACL, error)
	Update(ctx context.Context, aclID string, data *acl.Input) (*common.UpdateResponse, error)
}

type Connection interface {
	Create(ctx context.Context, parameters *connection.Parameters) (string, error)
	GetStatus(ctx context.Context, cloudType string) ([]*connection.Status, error)
}

type Device interface {
	Get(ctx context.Context, templateID string) (*device.Template, error)
	GetFromSite(ctx context.Context, siteID string) (*device.Template, error)
	GetByName(ctx context.Context, templateName string) (*device.Template, error)
	List(ctx context.Context) ([]*device.Template, error)
	Update(ctx context.Context, template *device.Template) (*device.UpdateResponse, error)
	GetAttachedDevices(ctx context.Context, templateID string) ([]*device.AttachedDevice, error)
	PushConfiguration(ctx context.Context, attachedDevices []*device.AttachedDevice, templateID string) (string, error)
}

type Feature interface {
	Create(ctx context.Context, template *feature.TemplateInput) (string, error)
	ListByType(ctx context.Context, templateType string) ([]*feature.TemplateInput, error)
	List(ctx context.Context) ([]*feature.TemplateInput, error)
	Get(ctx context.Context, templateID string) (*feature.Template, error)
	Update(ctx context.Context, templateID string, template *feature.Template) (*common.UpdateResponse, error)
}

type Policy interface {
	Get(ctx context.Context, policyID string) (*policy.Input, error)
	Delete(ctx context.Context, policyID string) error
	List(ctx context.Context) ([]*policy.Policy, error)
	Create(ctx context.Context, data *policy.Input) (string, error)
}

type Site interface {
	Get(ctx context.Context, siteID string) (*site.Site, error)
	List(ctx context.Context) ([]*site.Site, error)
}

type Status interface {
	ActionStatusLongPoll(ctx context.Context, id string) error
}

type VPC interface {
	Get(ctx context.Context, cloudType, vpcID string) (*vpc.VPC, error)
	GetByName(ctx context.Context, cloudType, vpcName string) (*vpc.VPC, error)
	ListAll(ctx context.Context) ([]*vpc.VPC, error)
	List(ctx context.Context, cloudType string) ([]*vpc.VPC, error)
	ListWithParameters(ctx context.Context, parameters *vpc.ListVPCParameters) ([]*vpc.VPC, error)
	ListWithTag(ctx context.Context, cloudType, tag string) ([]*vpc.VPC, error)
	ListWithTagWithParameters(ctx context.Context, parameters *vpc.ListVPCTagParameters) ([]*vpc.VPC, error)
	CreateVPCTag(ctx context.Context, tagName string, vpc *vpc.VPC) (string, error)
	DeleteVPCTag(ctx context.Context, cloudType, tag string) (string, error)
}

type VPN interface {
	List(ctx context.Context, cloudType string) ([]*vpn.Data, error)
}

type URLFiltering interface {
	List(ctx context.Context) ([]*urlfiltering.URLFilteringListRespStructure, error)
	Get(ctx context.Context, id string) (*urlfiltering.URLFilteringDefinition, error)
	GetByName(ctx context.Context, name string) (*urlfiltering.URLFilteringDefinition, error)
	Create(ctx context.Context, urlFiltering *urlfiltering.URLFilteringDefinition) (definitionId string, err error)
	Delete(ctx context.Context, urlFilteringId string) error
	DeleteByName(ctx context.Context, urlFilteringName string) error
}

type URLDenylist interface {
	List(ctx context.Context) ([]*udl.URLDenylistListResp, error)
	Get(ctx context.Context, id string) (*udl.URLDenylist, error)
	GetByName(ctx context.Context, name string) (*udl.URLDenylist, error)
	Create(ctx context.Context, urlDenylist *udl.URLDenylistInput) (id string, err error)
	Delete(ctx context.Context, urlDenylistId string) error
	DeleteByName(ctx context.Context, urlDenylistName string) error
}

type URLAllowlist interface {
	List(ctx context.Context) ([]*ual.URLAllowlistListResp, error)
	Get(ctx context.Context, id string) (*ual.URLAllowlist, error)
	GetByName(ctx context.Context, name string) (*ual.URLAllowlist, error)
	Create(ctx context.Context, urlDenylist *ual.URLAllowlistInput) (id string, err error)
	Delete(ctx context.Context, urlDenylistId string) error
	DeleteByName(ctx context.Context, urlDenylistName string) error
}

type ClientImpl struct {
	client             *c.Client
	aclClient          *acl.Client
	connectionClient   *connection.Client
	deviceClient       *device.Client
	featureClient      *feature.Client
	policyClient       *policy.Client
	siteClient         *site.Client
	statusClient       *status.Client
	vpcClient          *vpc.Client
	vpnClient          *vpn.Client
	urlFilteringClient *urlfiltering.Client
	urlDenylistClient  *udl.Client
	urlAllowlistClient *ual.Client
}

func (c *ClientImpl) Login(ctx context.Context, username, password string) error {
	return c.client.Login(ctx, username, password)
}

func (c *ClientImpl) GetToken() string {
	return c.client.Token
}

func (c *ClientImpl) SetToken(token string) {
	c.client.Token = token
}

func (c *ClientImpl) ACL() ACL {
	return c.aclClient
}

func (c *ClientImpl) Connection() Connection {
	return c.connectionClient
}

func (c *ClientImpl) Device() Device {
	return c.deviceClient
}

func (c *ClientImpl) Feature() Feature {
	return c.featureClient
}

func (c *ClientImpl) Policy() Policy {
	return c.policyClient
}

func (c *ClientImpl) Site() Site {
	return c.siteClient
}

func (c *ClientImpl) VPC() VPC {
	return c.vpcClient
}

func (c *ClientImpl) VPN() VPN {
	return c.vpnClient
}

func (c *ClientImpl) Status() Status {
	return c.statusClient
}

func (c *ClientImpl) URLFiltering() URLFiltering {
	return c.urlFilteringClient
}

func (c *ClientImpl) URLDenylist() URLDenylist {
	return c.urlDenylistClient
}

func (c *ClientImpl) URLAllowlist() URLAllowlist {
	return c.urlAllowlistClient
}

func NewDefaultClient(url string) (*ClientImpl, error) {
	defaultClient, err := c.NewDefaultClient(url)
	if err != nil {
		return nil, err
	}
	retries := 10
	return newClient(defaultClient, time.Second, retries), nil
}

func NewDefaultInsecureClientWithJar(url string) (*ClientImpl, error) {
	logger := logrus.New()
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("could not create cookie jar: %v", err)
	}
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
		MinVersion:         tls.VersionTLS12,
	}
	httpclient := &http.Client{
		Transport: transport,
		Jar:       jar,
	}
	vManageURL := url
	retries := 3
	retriesInterval := 10 * time.Second
	return NewClient(vManageURL, httpclient, logger, retriesInterval, retries)
}

func NewClient(
	url string,
	httpClient *http.Client,
	logger *logrus.Logger,
	longPollDuration time.Duration,
	longPollRetries int,
) (*ClientImpl, error) {
	defaultClient := c.NewClient(url, httpClient, logger)
	return newClient(defaultClient, longPollDuration, longPollRetries), nil
}

func newClient(defaultClient *c.Client, longPollDuration time.Duration, longPollRetries int) *ClientImpl {
	return &ClientImpl{
		client:             defaultClient,
		aclClient:          acl.NewClient(defaultClient),
		connectionClient:   connection.NewClient(defaultClient),
		deviceClient:       device.NewClient(defaultClient),
		siteClient:         site.NewClient(defaultClient),
		policyClient:       policy.NewClient(defaultClient),
		vpcClient:          vpc.NewClient(defaultClient),
		vpnClient:          vpn.NewClient(defaultClient),
		featureClient:      feature.NewClient(defaultClient),
		statusClient:       status.NewClient(defaultClient, longPollDuration, longPollRetries),
		urlFilteringClient: urlfiltering.NewClient(defaultClient),
		urlDenylistClient:  udl.NewClient(defaultClient),
		urlAllowlistClient: ual.NewClient(defaultClient),
	}
}
