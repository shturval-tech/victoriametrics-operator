/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/shturval-tech/victoriametrics-operator/api/client/versioned/typed/victoriametrics/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeVictoriametricsV1beta1 struct {
	*testing.Fake
}

func (c *FakeVictoriametricsV1beta1) VMAgents(namespace string) v1beta1.VMAgentInterface {
	return &FakeVMAgents{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMAlerts(namespace string) v1beta1.VMAlertInterface {
	return &FakeVMAlerts{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMAlertmanagers(namespace string) v1beta1.VMAlertmanagerInterface {
	return &FakeVMAlertmanagers{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMAlertmanagerConfigs(namespace string) v1beta1.VMAlertmanagerConfigInterface {
	return &FakeVMAlertmanagerConfigs{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMAuths(namespace string) v1beta1.VMAuthInterface {
	return &FakeVMAuths{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMClusters(namespace string) v1beta1.VMClusterInterface {
	return &FakeVMClusters{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMNodeScrapes(namespace string) v1beta1.VMNodeScrapeInterface {
	return &FakeVMNodeScrapes{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMPodScrapes(namespace string) v1beta1.VMPodScrapeInterface {
	return &FakeVMPodScrapes{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMProbes(namespace string) v1beta1.VMProbeInterface {
	return &FakeVMProbes{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMRules(namespace string) v1beta1.VMRuleInterface {
	return &FakeVMRules{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMServiceScrapes(namespace string) v1beta1.VMServiceScrapeInterface {
	return &FakeVMServiceScrapes{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMSingles(namespace string) v1beta1.VMSingleInterface {
	return &FakeVMSingles{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMStaticScrapes(namespace string) v1beta1.VMStaticScrapeInterface {
	return &FakeVMStaticScrapes{c, namespace}
}

func (c *FakeVictoriametricsV1beta1) VMUsers(namespace string) v1beta1.VMUserInterface {
	return &FakeVMUsers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeVictoriametricsV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
