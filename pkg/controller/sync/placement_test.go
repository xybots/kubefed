/*
Copyright 2018 The Kubernetes Authors.

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

package sync

import (
	"reflect"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/sets"

	fedv1b1 "sigs.k8s.io/kubefed/pkg/apis/core/v1beta1"
	"sigs.k8s.io/kubefed/pkg/controller/util"
)

func TestSelectedClusterNames(t *testing.T) {
	clusters := []*fedv1b1.KubeFedCluster{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "cluster1",
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{
				Name: "cluster2",
				Labels: map[string]string{
					"foo": "bar",
				},
			},
		},
	}

	testCases := map[string]struct {
		clusterNames    []string
		clusterSelector map[string]string
		expectedNames   sets.String
	}{
		"ignore cluster selector when cluster names present": {
			clusterNames:    []string{"cluster1"},
			clusterSelector: map[string]string{},
			expectedNames:   sets.NewString("cluster1"),
		},
		"no clusters when cluster names and selector absent": {
			expectedNames: sets.NewString(),
		},
		"no clusters when cluster names empty and selector not empty": {
			clusterNames: []string{},
			clusterSelector: map[string]string{
				"foo": "bar",
			},
			expectedNames: sets.NewString(),
		},
		"all clusters when cluster names absent and selector empty": {
			clusterSelector: map[string]string{},
			expectedNames:   sets.NewString("cluster1", "cluster2"),
		},
		"selected clusters when cluster names absent and selector not empty": {
			clusterSelector: map[string]string{
				"foo": "bar",
			},
			expectedNames: sets.NewString("cluster2"),
		},
	}

	for testName, testCase := range testCases {
		t.Run(testName, func(t *testing.T) {
			obj := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"spec": make(map[string]interface{}),
				},
			}
			if err := util.SetClusterNames(obj, testCase.clusterNames); err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if testCase.clusterSelector != nil {
				if err := unstructured.SetNestedStringMap(obj.Object, testCase.clusterSelector, util.SpecField, util.PlacementField, util.ClusterSelectorField, util.MatchLabelsField); err != nil {
					t.Fatalf("Unexpected error: %v", err)
				}
			}

			selectedNames, err := selectedClusterNames(obj, clusters)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if !reflect.DeepEqual(selectedNames, testCase.expectedNames) {
				t.Fatalf("Expected names %v, got %v", testCase.expectedNames, selectedNames)
			}
		})
	}
}
