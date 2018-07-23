// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/ligato/cn-infra/examples/crd-plugin/pkg/client/clientset/versioned/typed/crdexample.io/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCrdexampleV1 struct {
	*testing.Fake
}

func (c *FakeCrdexampleV1) CrdExamples(namespace string) v1.CrdExampleInterface {
	return &FakeCrdExamples{c, namespace}
}

func (c *FakeCrdexampleV1) CrdExampleEmbeds(namespace string) v1.CrdExampleEmbedInterface {
	return &FakeCrdExampleEmbeds{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCrdexampleV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
