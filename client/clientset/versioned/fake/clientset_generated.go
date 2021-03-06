/*
Copyright AppsCode Inc. and Contributors

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
	clientset "kubeform.dev/provider-upcloud-api/client/clientset/versioned"
	firewallv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/firewall/v1alpha1"
	fakefirewallv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/firewall/v1alpha1/fake"
	floatingv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/floating/v1alpha1"
	fakefloatingv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/floating/v1alpha1/fake"
	managedv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/managed/v1alpha1"
	fakemanagedv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/managed/v1alpha1/fake"
	networkv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/network/v1alpha1"
	fakenetworkv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/network/v1alpha1/fake"
	objectv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/object/v1alpha1"
	fakeobjectv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/object/v1alpha1/fake"
	routerv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/router/v1alpha1"
	fakerouterv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/router/v1alpha1/fake"
	serverv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/server/v1alpha1"
	fakeserverv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/server/v1alpha1/fake"
	storagev1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/storage/v1alpha1"
	fakestoragev1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/storage/v1alpha1/fake"
	tagv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/tag/v1alpha1"
	faketagv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/tag/v1alpha1/fake"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return &fakefirewallv1alpha1.FakeFirewallV1alpha1{Fake: &c.Fake}
}

// FloatingV1alpha1 retrieves the FloatingV1alpha1Client
func (c *Clientset) FloatingV1alpha1() floatingv1alpha1.FloatingV1alpha1Interface {
	return &fakefloatingv1alpha1.FakeFloatingV1alpha1{Fake: &c.Fake}
}

// ManagedV1alpha1 retrieves the ManagedV1alpha1Client
func (c *Clientset) ManagedV1alpha1() managedv1alpha1.ManagedV1alpha1Interface {
	return &fakemanagedv1alpha1.FakeManagedV1alpha1{Fake: &c.Fake}
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return &fakenetworkv1alpha1.FakeNetworkV1alpha1{Fake: &c.Fake}
}

// ObjectV1alpha1 retrieves the ObjectV1alpha1Client
func (c *Clientset) ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface {
	return &fakeobjectv1alpha1.FakeObjectV1alpha1{Fake: &c.Fake}
}

// RouterV1alpha1 retrieves the RouterV1alpha1Client
func (c *Clientset) RouterV1alpha1() routerv1alpha1.RouterV1alpha1Interface {
	return &fakerouterv1alpha1.FakeRouterV1alpha1{Fake: &c.Fake}
}

// ServerV1alpha1 retrieves the ServerV1alpha1Client
func (c *Clientset) ServerV1alpha1() serverv1alpha1.ServerV1alpha1Interface {
	return &fakeserverv1alpha1.FakeServerV1alpha1{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}

// TagV1alpha1 retrieves the TagV1alpha1Client
func (c *Clientset) TagV1alpha1() tagv1alpha1.TagV1alpha1Interface {
	return &faketagv1alpha1.FakeTagV1alpha1{Fake: &c.Fake}
}
