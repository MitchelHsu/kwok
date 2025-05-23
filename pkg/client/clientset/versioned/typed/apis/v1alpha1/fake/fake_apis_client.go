/*
Copyright The Kubernetes Authors.

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
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kwok/pkg/client/clientset/versioned/typed/apis/v1alpha1"
)

type FakeKwokV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKwokV1alpha1) Attaches(namespace string) v1alpha1.AttachInterface {
	return newFakeAttaches(c, namespace)
}

func (c *FakeKwokV1alpha1) ClusterAttaches() v1alpha1.ClusterAttachInterface {
	return newFakeClusterAttaches(c)
}

func (c *FakeKwokV1alpha1) ClusterExecs() v1alpha1.ClusterExecInterface {
	return newFakeClusterExecs(c)
}

func (c *FakeKwokV1alpha1) ClusterLogs() v1alpha1.ClusterLogsInterface {
	return newFakeClusterLogs(c)
}

func (c *FakeKwokV1alpha1) ClusterPortForwards() v1alpha1.ClusterPortForwardInterface {
	return newFakeClusterPortForwards(c)
}

func (c *FakeKwokV1alpha1) ClusterResourceUsages() v1alpha1.ClusterResourceUsageInterface {
	return newFakeClusterResourceUsages(c)
}

func (c *FakeKwokV1alpha1) Execs(namespace string) v1alpha1.ExecInterface {
	return newFakeExecs(c, namespace)
}

func (c *FakeKwokV1alpha1) Logs(namespace string) v1alpha1.LogsInterface {
	return newFakeLogs(c, namespace)
}

func (c *FakeKwokV1alpha1) Metrics() v1alpha1.MetricInterface {
	return newFakeMetrics(c)
}

func (c *FakeKwokV1alpha1) PortForwards(namespace string) v1alpha1.PortForwardInterface {
	return newFakePortForwards(c, namespace)
}

func (c *FakeKwokV1alpha1) ResourceUsages(namespace string) v1alpha1.ResourceUsageInterface {
	return newFakeResourceUsages(c, namespace)
}

func (c *FakeKwokV1alpha1) Stages() v1alpha1.StageInterface {
	return newFakeStages(c)
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKwokV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
