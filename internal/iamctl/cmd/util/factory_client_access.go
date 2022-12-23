// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// this file contains factories with no other dependencies

package util

import (
	"go_iam/pkg/cli/genericclioptions"
	"go_iam/pkg/sdk/marmotedu/service/iam"

	"go_iam/pkg/sdk/tools/clientcmd"

	restclient "go_iam/pkg/sdk/rest"
)

type factoryImpl struct {
	clientGetter genericclioptions.RESTClientGetter
}

func NewFactory(clientGetter genericclioptions.RESTClientGetter) Factory {
	if clientGetter == nil {
		panic("attempt to instantiate client_access_factory with nil clientGetter")
	}

	f := &factoryImpl{
		clientGetter: clientGetter,
	}

	return f
}

func (f *factoryImpl) ToRESTConfig() (*restclient.Config, error) {
	return f.clientGetter.ToRESTConfig()
}

func (f *factoryImpl) ToRawIAMConfigLoader() clientcmd.ClientConfig {
	return f.clientGetter.ToRawIAMConfigLoader()
}

func (f *factoryImpl) IAMClient() (*iam.IamClient, error) {
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	return iam.NewForConfig(clientConfig)
}

func (f *factoryImpl) RESTClient() (*restclient.RESTClient, error) {
	clientConfig, err := f.ToRESTConfig()
	if err != nil {
		return nil, err
	}
	setIAMDefaults(clientConfig)
	return restclient.RESTClientFor(clientConfig)
}
