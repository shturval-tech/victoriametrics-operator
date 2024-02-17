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

package v1beta1

import (
	"net/http"

	"github.com/shturval-tech/victoriametrics-operator/api/client/versioned/scheme"
	v1beta1 "github.com/shturval-tech/victoriametrics-operator/api/victoriametrics/v1beta1"
	rest "k8s.io/client-go/rest"
)

type VictoriametricsV1beta1Interface interface {
	RESTClient() rest.Interface
	VMAgentsGetter
	VMAlertsGetter
	VMAlertmanagersGetter
	VMAlertmanagerConfigsGetter
	VMAuthsGetter
	VMClustersGetter
	VMNodeScrapesGetter
	VMPodScrapesGetter
	VMProbesGetter
	VMRulesGetter
	VMServiceScrapesGetter
	VMSinglesGetter
	VMStaticScrapesGetter
	VMUsersGetter
}

// VictoriametricsV1beta1Client is used to interact with features provided by the victoriametrics group.
type VictoriametricsV1beta1Client struct {
	restClient rest.Interface
}

func (c *VictoriametricsV1beta1Client) VMAgents(namespace string) VMAgentInterface {
	return newVMAgents(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMAlerts(namespace string) VMAlertInterface {
	return newVMAlerts(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMAlertmanagers(namespace string) VMAlertmanagerInterface {
	return newVMAlertmanagers(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMAlertmanagerConfigs(namespace string) VMAlertmanagerConfigInterface {
	return newVMAlertmanagerConfigs(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMAuths(namespace string) VMAuthInterface {
	return newVMAuths(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMClusters(namespace string) VMClusterInterface {
	return newVMClusters(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMNodeScrapes(namespace string) VMNodeScrapeInterface {
	return newVMNodeScrapes(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMPodScrapes(namespace string) VMPodScrapeInterface {
	return newVMPodScrapes(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMProbes(namespace string) VMProbeInterface {
	return newVMProbes(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMRules(namespace string) VMRuleInterface {
	return newVMRules(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMServiceScrapes(namespace string) VMServiceScrapeInterface {
	return newVMServiceScrapes(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMSingles(namespace string) VMSingleInterface {
	return newVMSingles(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMStaticScrapes(namespace string) VMStaticScrapeInterface {
	return newVMStaticScrapes(c, namespace)
}

func (c *VictoriametricsV1beta1Client) VMUsers(namespace string) VMUserInterface {
	return newVMUsers(c, namespace)
}

// NewForConfig creates a new VictoriametricsV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*VictoriametricsV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new VictoriametricsV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*VictoriametricsV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &VictoriametricsV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new VictoriametricsV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *VictoriametricsV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new VictoriametricsV1beta1Client for the given RESTClient.
func New(c rest.Interface) *VictoriametricsV1beta1Client {
	return &VictoriametricsV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *VictoriametricsV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
