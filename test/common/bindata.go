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

// Code generated for package common by go-bindata DO NOT EDIT. (@generated)
// sources:
// test/common/fixtures/clusterroles.rbac.authorization.k8s.io.yaml
// test/common/fixtures/configmaps.yaml
// test/common/fixtures/deployments.apps.yaml
// test/common/fixtures/ingresses.extensions.yaml
// test/common/fixtures/jobs.batch.yaml
// test/common/fixtures/namespaces.yaml
// test/common/fixtures/replicasets.apps.yaml
// test/common/fixtures/secrets.yaml
// test/common/fixtures/serviceaccounts.yaml
// test/common/fixtures/services.yaml
// config/kubefedconfig.yaml
// config/enabletypedirectives/clusterroles.rbac.authorization.k8s.io.yaml
// config/enabletypedirectives/configmaps.yaml
// config/enabletypedirectives/deployments.apps.yaml
// config/enabletypedirectives/ingresses.extensions.yaml
// config/enabletypedirectives/jobs.batch.yaml
// config/enabletypedirectives/namespaces.yaml
// config/enabletypedirectives/replicasets.apps.yaml
// config/enabletypedirectives/secrets.yaml
// config/enabletypedirectives/serviceaccounts.yaml
// config/enabletypedirectives/services.yaml
package common

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _testCommonFixturesClusterrolesRbacAuthorizationK8sIoYaml = []byte(`kind: fixture
template:
  rules:
  - apiGroups:
      - '*'
    resources:
      - '*'
    verbs:
      - '*'
`)

func testCommonFixturesClusterrolesRbacAuthorizationK8sIoYamlBytes() ([]byte, error) {
	return _testCommonFixturesClusterrolesRbacAuthorizationK8sIoYaml, nil
}

func testCommonFixturesClusterrolesRbacAuthorizationK8sIoYaml() (*asset, error) {
	bytes, err := testCommonFixturesClusterrolesRbacAuthorizationK8sIoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/clusterroles.rbac.authorization.k8s.io.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesConfigmapsYaml = []byte(`kind: fixture
template:
  data:
    foo: bar
overrides:
  - clusterOverrides:
      - path: /data
        value:
          foo: baz
`)

func testCommonFixturesConfigmapsYamlBytes() ([]byte, error) {
	return _testCommonFixturesConfigmapsYaml, nil
}

func testCommonFixturesConfigmapsYaml() (*asset, error) {
	bytes, err := testCommonFixturesConfigmapsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/configmaps.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesDeploymentsAppsYaml = []byte(`kind: fixture
template:
  spec:
    replicas: 2
    selector:
      matchLabels:
        foo: bar
    template:
      metadata:
        labels:
          foo: bar
      spec:
        terminationGracePeriodSeconds: 0
        containers:
          - name: busybox
            image: busybox
            command: ["/bin/sh", "-c", "trap : TERM INT; (while true; do sleep 1000; done) & wait"]
overrides:
  - clusterOverrides:
    - path: /spec/replicas
      value: 1
    - path: "/spec/template/spec/containers/0/image"
      value: "busybox:1.31.0-uclibc"
    - path: "/metadata/annotations"
      op: "add"
      value:
        foo: bar
    - path: "/metadata/annotations/foo"
      op: "remove"
`)

func testCommonFixturesDeploymentsAppsYamlBytes() ([]byte, error) {
	return _testCommonFixturesDeploymentsAppsYaml, nil
}

func testCommonFixturesDeploymentsAppsYaml() (*asset, error) {
	bytes, err := testCommonFixturesDeploymentsAppsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/deployments.apps.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesIngressesExtensionsYaml = []byte(`kind: fixture
template:
  spec:
    backend:
      serviceName: testsvc
      servicePort: 80
`)

func testCommonFixturesIngressesExtensionsYamlBytes() ([]byte, error) {
	return _testCommonFixturesIngressesExtensionsYaml, nil
}

func testCommonFixturesIngressesExtensionsYaml() (*asset, error) {
	bytes, err := testCommonFixturesIngressesExtensionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/ingresses.extensions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesJobsBatchYaml = []byte(`kind: fixture
template:
  spec:
    parallelism: 1
    selector:
      matchLabels:
        foo: bar
    manualSelector: true
    template:
      metadata:
        labels:
          foo: bar
      spec:
        terminationGracePeriodSeconds: 0
        restartPolicy: Never
        containers:
          - name: busybox
            image: busybox
            command: ["/bin/sh", "-c", "trap : TERM INT; (while true; do sleep 1000; done) & wait"]
overrides:
  - clusterOverrides:
    - path: /spec/parallelism
      value: 2
`)

func testCommonFixturesJobsBatchYamlBytes() ([]byte, error) {
	return _testCommonFixturesJobsBatchYaml, nil
}

func testCommonFixturesJobsBatchYaml() (*asset, error) {
	bytes, err := testCommonFixturesJobsBatchYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/jobs.batch.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesNamespacesYaml = []byte(`kind: fixture
`)

func testCommonFixturesNamespacesYamlBytes() ([]byte, error) {
	return _testCommonFixturesNamespacesYaml, nil
}

func testCommonFixturesNamespacesYaml() (*asset, error) {
	bytes, err := testCommonFixturesNamespacesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/namespaces.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesReplicasetsAppsYaml = []byte(`kind: fixture
template:
  spec:
    replicas: 1
    selector:
      matchLabels:
        foo: bar
    template:
      metadata:
        labels:
          foo: bar
      spec:
        terminationGracePeriodSeconds: 0
        containers:
          - name: busybox
            image: busybox
            command: ["/bin/sh", "-c", "trap : TERM INT; (while true; do sleep 1000; done) & wait"]
overrides:
  - clusterOverrides:
    - path: /spec/replicas
      value: 2
`)

func testCommonFixturesReplicasetsAppsYamlBytes() ([]byte, error) {
	return _testCommonFixturesReplicasetsAppsYaml, nil
}

func testCommonFixturesReplicasetsAppsYaml() (*asset, error) {
	bytes, err := testCommonFixturesReplicasetsAppsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/replicasets.apps.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesSecretsYaml = []byte(`kind: fixture
template:
  data:
    foo: YmFy
  type: Opaque
overrides:
  - clusterOverrides:
      - path: /data
        value:
          foo: YmF6
`)

func testCommonFixturesSecretsYamlBytes() ([]byte, error) {
	return _testCommonFixturesSecretsYaml, nil
}

func testCommonFixturesSecretsYaml() (*asset, error) {
	bytes, err := testCommonFixturesSecretsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/secrets.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesServiceaccountsYaml = []byte(`kind: fixture
template:
  automountServiceAccountToken: true
`)

func testCommonFixturesServiceaccountsYamlBytes() ([]byte, error) {
	return _testCommonFixturesServiceaccountsYaml, nil
}

func testCommonFixturesServiceaccountsYaml() (*asset, error) {
	bytes, err := testCommonFixturesServiceaccountsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/serviceaccounts.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCommonFixturesServicesYaml = []byte(`kind: fixture
template:
  spec:
    type: ClusterIP
    ports:
      - name: http
        port: 80
`)

func testCommonFixturesServicesYamlBytes() ([]byte, error) {
	return _testCommonFixturesServicesYaml, nil
}

func testCommonFixturesServicesYaml() (*asset, error) {
	bytes, err := testCommonFixturesServicesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/common/fixtures/services.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configKubefedconfigYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: KubeFedConfig
metadata:
  name: kubefed
  namespace: kube-federation-system
spec:
  scope: Cluster
  controllerDuration:
    availableDelay: 20s
    unavailableDelay: 60s
  leaderElect:
    leaseDuration: 1500ms
    renewDeadline: 1000ms
    resourceLock: configmaps
    retryPeriod: 500ms
  featureGates:
  - name: PushReconciler
    configuration: "Enabled"
  - name: SchedulerPreferences
    configuration: "Enabled"
  - name: CrossClusterServiceDiscovery
    configuration: "Enabled"
  - name: FederatedIngress
    configuration: "Enabled"
  clusterHealthCheck:
    failureThreshold: 3
    period: 10s
    successThreshold: 1
    timeout: 3s
  syncController:
    adoptResources: Enabled
`)

func configKubefedconfigYamlBytes() ([]byte, error) {
	return _configKubefedconfigYaml, nil
}

func configKubefedconfigYaml() (*asset, error) {
	bytes, err := configKubefedconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/kubefedconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: clusterroles.rbac.authorization.k8s.io
`)

func configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYaml, nil
}

func configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/clusterroles.rbac.authorization.k8s.io.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesConfigmapsYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: configmaps
`)

func configEnabletypedirectivesConfigmapsYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesConfigmapsYaml, nil
}

func configEnabletypedirectivesConfigmapsYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesConfigmapsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/configmaps.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesDeploymentsAppsYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: deployments.apps
`)

func configEnabletypedirectivesDeploymentsAppsYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesDeploymentsAppsYaml, nil
}

func configEnabletypedirectivesDeploymentsAppsYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesDeploymentsAppsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/deployments.apps.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesIngressesExtensionsYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: ingresses.extensions
`)

func configEnabletypedirectivesIngressesExtensionsYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesIngressesExtensionsYaml, nil
}

func configEnabletypedirectivesIngressesExtensionsYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesIngressesExtensionsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/ingresses.extensions.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesJobsBatchYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: jobs.batch
`)

func configEnabletypedirectivesJobsBatchYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesJobsBatchYaml, nil
}

func configEnabletypedirectivesJobsBatchYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesJobsBatchYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/jobs.batch.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesNamespacesYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: namespaces
`)

func configEnabletypedirectivesNamespacesYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesNamespacesYaml, nil
}

func configEnabletypedirectivesNamespacesYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesNamespacesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/namespaces.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesReplicasetsAppsYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: replicasets.apps
`)

func configEnabletypedirectivesReplicasetsAppsYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesReplicasetsAppsYaml, nil
}

func configEnabletypedirectivesReplicasetsAppsYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesReplicasetsAppsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/replicasets.apps.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesSecretsYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: secrets
`)

func configEnabletypedirectivesSecretsYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesSecretsYaml, nil
}

func configEnabletypedirectivesSecretsYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesSecretsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/secrets.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesServiceaccountsYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: serviceaccounts
`)

func configEnabletypedirectivesServiceaccountsYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesServiceaccountsYaml, nil
}

func configEnabletypedirectivesServiceaccountsYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesServiceaccountsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/serviceaccounts.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configEnabletypedirectivesServicesYaml = []byte(`apiVersion: core.kubefed.io/v1beta1
kind: EnableTypeDirective
metadata:
  name: services
`)

func configEnabletypedirectivesServicesYamlBytes() ([]byte, error) {
	return _configEnabletypedirectivesServicesYaml, nil
}

func configEnabletypedirectivesServicesYaml() (*asset, error) {
	bytes, err := configEnabletypedirectivesServicesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/enabletypedirectives/services.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"test/common/fixtures/clusterroles.rbac.authorization.k8s.io.yaml":        testCommonFixturesClusterrolesRbacAuthorizationK8sIoYaml,
	"test/common/fixtures/configmaps.yaml":                                    testCommonFixturesConfigmapsYaml,
	"test/common/fixtures/deployments.apps.yaml":                              testCommonFixturesDeploymentsAppsYaml,
	"test/common/fixtures/ingresses.extensions.yaml":                          testCommonFixturesIngressesExtensionsYaml,
	"test/common/fixtures/jobs.batch.yaml":                                    testCommonFixturesJobsBatchYaml,
	"test/common/fixtures/namespaces.yaml":                                    testCommonFixturesNamespacesYaml,
	"test/common/fixtures/replicasets.apps.yaml":                              testCommonFixturesReplicasetsAppsYaml,
	"test/common/fixtures/secrets.yaml":                                       testCommonFixturesSecretsYaml,
	"test/common/fixtures/serviceaccounts.yaml":                               testCommonFixturesServiceaccountsYaml,
	"test/common/fixtures/services.yaml":                                      testCommonFixturesServicesYaml,
	"config/kubefedconfig.yaml":                                               configKubefedconfigYaml,
	"config/enabletypedirectives/clusterroles.rbac.authorization.k8s.io.yaml": configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYaml,
	"config/enabletypedirectives/configmaps.yaml":                             configEnabletypedirectivesConfigmapsYaml,
	"config/enabletypedirectives/deployments.apps.yaml":                       configEnabletypedirectivesDeploymentsAppsYaml,
	"config/enabletypedirectives/ingresses.extensions.yaml":                   configEnabletypedirectivesIngressesExtensionsYaml,
	"config/enabletypedirectives/jobs.batch.yaml":                             configEnabletypedirectivesJobsBatchYaml,
	"config/enabletypedirectives/namespaces.yaml":                             configEnabletypedirectivesNamespacesYaml,
	"config/enabletypedirectives/replicasets.apps.yaml":                       configEnabletypedirectivesReplicasetsAppsYaml,
	"config/enabletypedirectives/secrets.yaml":                                configEnabletypedirectivesSecretsYaml,
	"config/enabletypedirectives/serviceaccounts.yaml":                        configEnabletypedirectivesServiceaccountsYaml,
	"config/enabletypedirectives/services.yaml":                               configEnabletypedirectivesServicesYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"enabletypedirectives": &bintree{nil, map[string]*bintree{
			"clusterroles.rbac.authorization.k8s.io.yaml": &bintree{configEnabletypedirectivesClusterrolesRbacAuthorizationK8sIoYaml, map[string]*bintree{}},
			"configmaps.yaml":                             &bintree{configEnabletypedirectivesConfigmapsYaml, map[string]*bintree{}},
			"deployments.apps.yaml":                       &bintree{configEnabletypedirectivesDeploymentsAppsYaml, map[string]*bintree{}},
			"ingresses.extensions.yaml":                   &bintree{configEnabletypedirectivesIngressesExtensionsYaml, map[string]*bintree{}},
			"jobs.batch.yaml":                             &bintree{configEnabletypedirectivesJobsBatchYaml, map[string]*bintree{}},
			"namespaces.yaml":                             &bintree{configEnabletypedirectivesNamespacesYaml, map[string]*bintree{}},
			"replicasets.apps.yaml":                       &bintree{configEnabletypedirectivesReplicasetsAppsYaml, map[string]*bintree{}},
			"secrets.yaml":                                &bintree{configEnabletypedirectivesSecretsYaml, map[string]*bintree{}},
			"serviceaccounts.yaml":                        &bintree{configEnabletypedirectivesServiceaccountsYaml, map[string]*bintree{}},
			"services.yaml":                               &bintree{configEnabletypedirectivesServicesYaml, map[string]*bintree{}},
		}},
		"kubefedconfig.yaml": &bintree{configKubefedconfigYaml, map[string]*bintree{}},
	}},
	"test": &bintree{nil, map[string]*bintree{
		"common": &bintree{nil, map[string]*bintree{
			"fixtures": &bintree{nil, map[string]*bintree{
				"clusterroles.rbac.authorization.k8s.io.yaml": &bintree{testCommonFixturesClusterrolesRbacAuthorizationK8sIoYaml, map[string]*bintree{}},
				"configmaps.yaml":                             &bintree{testCommonFixturesConfigmapsYaml, map[string]*bintree{}},
				"deployments.apps.yaml":                       &bintree{testCommonFixturesDeploymentsAppsYaml, map[string]*bintree{}},
				"ingresses.extensions.yaml":                   &bintree{testCommonFixturesIngressesExtensionsYaml, map[string]*bintree{}},
				"jobs.batch.yaml":                             &bintree{testCommonFixturesJobsBatchYaml, map[string]*bintree{}},
				"namespaces.yaml":                             &bintree{testCommonFixturesNamespacesYaml, map[string]*bintree{}},
				"replicasets.apps.yaml":                       &bintree{testCommonFixturesReplicasetsAppsYaml, map[string]*bintree{}},
				"secrets.yaml":                                &bintree{testCommonFixturesSecretsYaml, map[string]*bintree{}},
				"serviceaccounts.yaml":                        &bintree{testCommonFixturesServiceaccountsYaml, map[string]*bintree{}},
				"services.yaml":                               &bintree{testCommonFixturesServicesYaml, map[string]*bintree{}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
