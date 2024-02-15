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

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/app-net-interface/catalyst-sdwan-app-client/vmanage"
)

func main() {
	url := ""
	username := ""
	password := ""
	ctx := context.TODO()
	client, err := vmanage.NewDefaultClient(url)
	if err := client.Login(ctx, username, password); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	templates, err := client.Device().List(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(templates) //nolint

	vpcs, err := client.VPC().List(ctx, "aws")
	if err != nil {
		panic(err)
	}
	fmt.Println(vpcs) //nolint

	//////////////////////////////////////////////////
	// self-signed certificate

	localCertFile := ""
	rootCAs := addCertToRootCAs(localCertFile)

	transport := http.DefaultTransport.(*http.Transport).Clone() //nolint:forcetypeassert
	transport.TLSClientConfig = &tls.Config{
		MinVersion:         tls.VersionTLS12,
		InsecureSkipVerify: false,
		RootCAs:            rootCAs,
		ServerName:         "", // Set to Subject Alternative Name (SAN)
	}
	c := &http.Client{
		Transport: transport,
	}
	client, err = vmanage.NewClient(url, c, logrus.New(), 1, 1)
	if err != nil {
		panic(err)
	}
	if err := client.Login(ctx, username, password); err != nil {
		panic(err)
	}
	if err != nil {
		panic(err)
	}

	vpcs, err = client.VPC().List(ctx, "aws")
	if err != nil {
		panic(err)
	}
	fmt.Println(vpcs) //nolint
}

func addCertToRootCAs(localCertFile string) *x509.CertPool {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	certs, err := ioutil.ReadFile(localCertFile)
	if err != nil {
		panic(err)
	}

	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		panic(err)
	}

	return rootCAs
}
