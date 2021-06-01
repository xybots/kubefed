/*
Copyright 2016 The Kubernetes Authors.

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

package kubefedcluster

import (
	"strings"
	"time"

	"github.com/pkg/errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	kubeclientset "k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/klog"

	fedcommon "sigs.k8s.io/kubefed/pkg/apis/core/common"
	fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	"sigs.k8s.io/kubefed/pkg/client/generic"
	"sigs.k8s.io/kubefed/pkg/controller/util"
	"sigs.k8s.io/kubefed/pkg/metrics"
)

const (
	UserAgentName = "Cluster-Controller"

	// Following labels come from k8s.io/kubernetes/pkg/kubelet/apis
	LabelZoneFailureDomain = "failure-domain.beta.kubernetes.io/zone"
	LabelZoneRegion        = "failure-domain.beta.kubernetes.io/region"

	// Common ClusterConditions for KubeFedClusterStatus
	ClusterReady              = "ClusterReady"
	HealthzOk                 = "/healthz responded with ok"
	ClusterNotReady           = "ClusterNotReady"
	HealthzNotOk              = "/healthz responded without ok"
	ClusterNotReachableReason = "ClusterNotReachable"
	ClusterNotReachableMsg    = "cluster is not reachable"
	ClusterReachableReason    = "ClusterReachable"
	ClusterReachableMsg       = "cluster is reachable"
)

// ClusterClient provides methods for determining the status and zones of a
// particular KubeFedCluster.
type ClusterClient struct {
	kubeClient  *kubeclientset.Clientset
	clusterName string
}

// NewClusterClientSet returns a ClusterClient for the given KubeFedCluster.
// The kubeClient is used to configure the ClusterClient's internal client
// with information from a kubeconfig stored in a kubernetes secret.
func NewClusterClientSet(c *fedv1b1.KubeFedCluster, client generic.Client, fedNamespace string, timeout time.Duration) (*ClusterClient, error) {
	clusterConfig, err := util.BuildClusterConfig(c, client, fedNamespace)
	if err != nil {
		return nil, err
	}
	clusterConfig.Timeout = timeout
	var clusterClientSet = ClusterClient{clusterName: c.Name}
	if clusterConfig != nil {
		clusterClientSet.kubeClient = kubeclientset.NewForConfigOrDie((restclient.AddUserAgent(clusterConfig, UserAgentName)))
		if clusterClientSet.kubeClient == nil {
			return nil, nil
		}
	}
	return &clusterClientSet, nil
}

// GetClusterHealthStatus gets the kubernetes cluster health status by requesting "/healthz"
func (self *ClusterClient) GetClusterHealthStatus() (*fedv1b1.KubeFedClusterStatus, error) {
	clusterStatus := fedv1b1.KubeFedClusterStatus{}
	currentTime := metav1.Now()
	clusterReady := ClusterReady
	healthzOk := HealthzOk
	newClusterReadyCondition := fedv1b1.ClusterCondition{
		Type:               fedcommon.ClusterReady,
		Status:             corev1.ConditionTrue,
		Reason:             &clusterReady,
		Message:            &healthzOk,
		LastProbeTime:      currentTime,
		LastTransitionTime: &currentTime,
	}
	clusterNotReady := ClusterNotReady
	healthzNotOk := HealthzNotOk
	newClusterNotReadyCondition := fedv1b1.ClusterCondition{
		Type:               fedcommon.ClusterReady,
		Status:             corev1.ConditionFalse,
		Reason:             &clusterNotReady,
		Message:            &healthzNotOk,
		LastProbeTime:      currentTime,
		LastTransitionTime: &currentTime,
	}
	clusterNotReachableReason := ClusterNotReachableReason
	clusterNotReachableMsg := ClusterNotReachableMsg
	newClusterOfflineCondition := fedv1b1.ClusterCondition{
		Type:               fedcommon.ClusterOffline,
		Status:             corev1.ConditionTrue,
		Reason:             &clusterNotReachableReason,
		Message:            &clusterNotReachableMsg,
		LastProbeTime:      currentTime,
		LastTransitionTime: &currentTime,
	}
	clusterReachableReason := ClusterReachableReason
	clusterReachableMsg := ClusterReachableMsg
	newClusterNotOfflineCondition := fedv1b1.ClusterCondition{
		Type:               fedcommon.ClusterOffline,
		Status:             corev1.ConditionFalse,
		Reason:             &clusterReachableReason,
		Message:            &clusterReachableMsg,
		LastProbeTime:      currentTime,
		LastTransitionTime: &currentTime,
	}
	body, err := self.kubeClient.DiscoveryClient.RESTClient().Get().AbsPath("/healthz").Do().Raw()
	if err != nil {
		runtime.HandleError(errors.Wrapf(err, "Failed to do cluster health check for cluster %q", self.clusterName))
		clusterStatus.Conditions = append(clusterStatus.Conditions, newClusterOfflineCondition)
		metrics.RegisterKubefedClusterTotal(metrics.ClusterOffline, self.clusterName)
	} else {
		if !strings.EqualFold(string(body), "ok") {
			metrics.RegisterKubefedClusterTotal(metrics.ClusterNotReady, self.clusterName)
			clusterStatus.Conditions = append(clusterStatus.Conditions, newClusterNotReadyCondition, newClusterNotOfflineCondition)
		} else {
			metrics.RegisterKubefedClusterTotal(metrics.ClusterReady, self.clusterName)
			clusterStatus.Conditions = append(clusterStatus.Conditions, newClusterReadyCondition)
		}
	}

	return &clusterStatus, err
}

// GetClusterZones gets the kubernetes cluster zones and region by inspecting labels on nodes in the cluster.
func (self *ClusterClient) GetClusterZones() ([]string, string, error) {
	nodes, err := self.kubeClient.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		klog.Errorf("Failed to list nodes while getting zone names: %v", err)
		return nil, "", err
	}

	zones := sets.NewString()
	region := ""
	for i, node := range nodes.Items {
		zone := getZoneNameForNode(node)
		// region is same for all nodes in the cluster, so just pick the region from first node.
		if i == 0 {
			region = getRegionNameForNode(node)
		}
		if zone != "" && !zones.Has(zone) {
			zones.Insert(zone)
		}
	}
	return zones.List(), region, nil
}

// Find the name of the zone in which a Node is running.
func getZoneNameForNode(node corev1.Node) string {
	for key, value := range node.Labels {
		if key == LabelZoneFailureDomain {
			return value
		}
	}
	return ""
}

// Find the name of the region in which a Node is running.
func getRegionNameForNode(node corev1.Node) string {
	for key, value := range node.Labels {
		if key == LabelZoneRegion {
			return value
		}
	}
	return ""
}
