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

package controllers

import (
	"context"
	"fmt"
	"github.com/VictoriaMetrics/operator/controllers/factory"
	"github.com/VictoriaMetrics/operator/controllers/factory/alertmanager"
	"github.com/VictoriaMetrics/operator/controllers/factory/limiter"
	"github.com/VictoriaMetrics/operator/internal/config"
	"github.com/go-logr/logr"
	"gopkg.in/yaml.v2"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	operatorv1beta1 "github.com/VictoriaMetrics/operator/api/v1beta1"
)

var (
	vmaConfigRateLimiter = limiter.NewRateLimiter("vmalertmanager", 5)
)

// VMAlertmanagerConfigReconciler reconciles a VMAlertmanagerConfig object
type VMAlertmanagerConfigReconciler struct {
	client.Client
	Log          logr.Logger
	OriginScheme *runtime.Scheme
	BaseConf     *config.BaseOperatorConf
}

// Scheme implements interface.
func (r *VMAlertmanagerConfigReconciler) Scheme() *runtime.Scheme {
	return r.OriginScheme
}

// Reconcile implements interface
// +kubebuilder:rbac:groups=operator.victoriametrics.com,resources=vmalertmanagerconfigs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.victoriametrics.com,resources=vmalertmanagerconfigs/status,verbs=get;update;patch
func (r *VMAlertmanagerConfigReconciler) Reconcile(ctx context.Context, req ctrl.Request) (result ctrl.Result, err error) {
	l := r.Log.WithValues("vmalertmanagerconfig", req.NamespacedName, "name", req.Name)

	var instance operatorv1beta1.VMAlertmanagerConfig
	if err := r.Client.Get(ctx, req.NamespacedName, &instance); err != nil {
		return handleGetError(req, "vmalertmanagerconfig", err)
	}
	original := instance.DeepCopy()
	RegisterObjectStat(&instance, "vmalertmanagerconfig")

	if vmaConfigRateLimiter.MustThrottleReconcile() {
		return
	}

	alertmanagerLock.Lock()
	defer alertmanagerLock.Unlock()

	// select alertmanagers
	var vmams operatorv1beta1.VMAlertmanagerList
	if err := r.Client.List(ctx, &vmams, config.MustGetNamespaceListOptions()); err != nil {
		l.Error(err, "cannot list vmalertmanagers")
		return ctrl.Result{}, err
	}
	for _, item := range vmams.Items {
		am := &item
		if !am.DeletionTimestamp.IsZero() || am.Spec.ParsingError != "" {
			continue
		}
		l := l.WithValues("alertmanager", am.Name)
		ismatch, err := isSelectorsMatches(&instance, am, am.Spec.ConfigSelector)
		if err != nil {
			l.Error(err, "cannot match alertmanager against selector, probably bug")
			continue
		}
		if !ismatch {
			// selector do not match fast path
			continue
		}

		if err := r.ValidateGlobal(instance, am); err != nil {
			l.Error(err, "cannot  reconcile alertmanager")
			instance.Status.ErrorReason += fmt.Sprintf("cannot  reconcile alertmanager: %s; ", err.Error())
			continue
		}

		if err := factory.CreateOrUpdateAlertManager(ctx, am, r.Client, r.BaseConf); err != nil {
			l.Error(err, "cannot  reconcile alertmanager")
			continue
		}
	}

	defer func() {
		if instance.DeletionTimestamp == nil {
			if statusError := r.Client.Status().Patch(ctx, &instance, client.MergeFrom(original)); statusError != nil {
				l.Error(err, "failed resource patch status VmAlertManagerConfig")
				if err == nil {
					err = statusError
				}
			}
		}
	}()

	return
}

func (ac *VMAlertmanagerConfigReconciler) ValidateGlobal(c operatorv1beta1.VMAlertmanagerConfig, m *operatorv1beta1.VMAlertmanager) error {
	globalConfig := alertmanager.GlobalConfig{}
	if len(m.Spec.ConfigRawYaml) == 0 {
		globalConfig = alertmanager.DefaultGlobalConfig()
	}
	if err := yaml.Unmarshal([]byte(m.Spec.ConfigRawYaml), &globalConfig); err != nil {
		return err
	}
	global := globalConfig.Global
	for _, rcv := range c.Spec.Receivers {
		for _, ec := range rcv.EmailConfigs {
			if ec.Smarthost == "" {
				if global.SMTPSmartHostPort == "" {
					return fmt.Errorf("no global SMTP smarthost set")
				}
			}
			if ec.From == "" {
				if global.SMTPFrom == "" {
					return fmt.Errorf("no global SMTP from set")
				}
			}
		}
		for _, sc := range rcv.SlackConfigs {

			if sc.APIURL == nil {
				if _, err := alertmanager.ToURL(global.SlackAPIURL); err != nil && len(global.SlackAPIURLFile) == 0 {
					return fmt.Errorf("no global Slack API URL set either inline or in a file")
				}
			}
		}

		for _, pdc := range rcv.PagerDutyConfigs {
			if pdc.URL == "" {
				if _, err := alertmanager.ToURL(global.PagerdutyURL); err != nil {
					return fmt.Errorf("no global PagerDuty URL set")
				}
			}
		}
		for _, ogc := range rcv.OpsGenieConfigs {
			if ogc.APIURL == "" {
				if _, err := alertmanager.ToURL(global.OpsGenieAPIURL); err != nil {
					return fmt.Errorf("no global OpsGenie URL set")
				}
			}
			if ogc.APIKey == nil {
				if global.OpsGenieAPIKey == "" && len(global.OpsGenieAPIKeyFile) == 0 {
					return fmt.Errorf("no global OpsGenie API Key set either inline or in a file")
				}
			}
		}
		for _, wcc := range rcv.WeChatConfigs {
			if wcc.APIURL == "" {
				if _, err := alertmanager.ToURL(global.WeChatAPIURL); err != nil {
					return fmt.Errorf("no global Wechat URL set")
				}
			}

			if wcc.APISecret == nil {
				if global.WeChatAPISecret == "" {
					return fmt.Errorf("no global Wechat ApiSecret set")
				}
			}

			if wcc.CorpID == "" {
				if global.WeChatAPICorpID == "" {
					return fmt.Errorf("no global Wechat CorpID set")
				}
			}
		}
		for _, voc := range rcv.VictorOpsConfigs {
			if voc.APIURL == "" {
				if _, err := alertmanager.ToURL(global.VictorOpsAPIURL); err != nil {
					return fmt.Errorf("no global VictorOps URL set")
				}
			}
			if voc.APIKey == nil {
				if global.VictorOpsAPIKey == "" && len(global.VictorOpsAPIKeyFile) == 0 {
					return fmt.Errorf("no global VictorOps API Key set")
				}
			}
		}
		for _, discord := range rcv.DiscordConfigs {
			if discord.URL == nil {
				return fmt.Errorf("no discord webhook URL provided")
			}
		}
		for _, webex := range rcv.WebexConfigs {
			if webex.URL == nil {
				if _, err := alertmanager.ToURL(global.WebexAPIURL); err != nil {
					return fmt.Errorf("no global Webex URL set")
				}
			}
		}
		for _, msteams := range rcv.MSTeamsConfigs {
			if msteams.URL == nil {
				return fmt.Errorf("no msteams webhook URL provided")
			}
		}
	}
	return nil
}

// SetupWithManager configures reconcile
func (r *VMAlertmanagerConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1beta1.VMAlertmanagerConfig{}).
		WithOptions(getDefaultOptions()).
		Complete(r)
}
