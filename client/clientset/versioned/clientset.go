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

package versioned

import (
	"fmt"

	firewallv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/firewall/v1alpha1"
	floatingv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/floating/v1alpha1"
	managedv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/managed/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/network/v1alpha1"
	objectv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/object/v1alpha1"
	routerv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/router/v1alpha1"
	serverv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/server/v1alpha1"
	storagev1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/storage/v1alpha1"
	tagv1alpha1 "kubeform.dev/provider-upcloud-api/client/clientset/versioned/typed/tag/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface
	FloatingV1alpha1() floatingv1alpha1.FloatingV1alpha1Interface
	ManagedV1alpha1() managedv1alpha1.ManagedV1alpha1Interface
	NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface
	ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface
	RouterV1alpha1() routerv1alpha1.RouterV1alpha1Interface
	ServerV1alpha1() serverv1alpha1.ServerV1alpha1Interface
	StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface
	TagV1alpha1() tagv1alpha1.TagV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	firewallV1alpha1 *firewallv1alpha1.FirewallV1alpha1Client
	floatingV1alpha1 *floatingv1alpha1.FloatingV1alpha1Client
	managedV1alpha1  *managedv1alpha1.ManagedV1alpha1Client
	networkV1alpha1  *networkv1alpha1.NetworkV1alpha1Client
	objectV1alpha1   *objectv1alpha1.ObjectV1alpha1Client
	routerV1alpha1   *routerv1alpha1.RouterV1alpha1Client
	serverV1alpha1   *serverv1alpha1.ServerV1alpha1Client
	storageV1alpha1  *storagev1alpha1.StorageV1alpha1Client
	tagV1alpha1      *tagv1alpha1.TagV1alpha1Client
}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return c.firewallV1alpha1
}

// FloatingV1alpha1 retrieves the FloatingV1alpha1Client
func (c *Clientset) FloatingV1alpha1() floatingv1alpha1.FloatingV1alpha1Interface {
	return c.floatingV1alpha1
}

// ManagedV1alpha1 retrieves the ManagedV1alpha1Client
func (c *Clientset) ManagedV1alpha1() managedv1alpha1.ManagedV1alpha1Interface {
	return c.managedV1alpha1
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return c.networkV1alpha1
}

// ObjectV1alpha1 retrieves the ObjectV1alpha1Client
func (c *Clientset) ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface {
	return c.objectV1alpha1
}

// RouterV1alpha1 retrieves the RouterV1alpha1Client
func (c *Clientset) RouterV1alpha1() routerv1alpha1.RouterV1alpha1Interface {
	return c.routerV1alpha1
}

// ServerV1alpha1 retrieves the ServerV1alpha1Client
func (c *Clientset) ServerV1alpha1() serverv1alpha1.ServerV1alpha1Interface {
	return c.serverV1alpha1
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return c.storageV1alpha1
}

// TagV1alpha1 retrieves the TagV1alpha1Client
func (c *Clientset) TagV1alpha1() tagv1alpha1.TagV1alpha1Interface {
	return c.tagV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.firewallV1alpha1, err = firewallv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.floatingV1alpha1, err = floatingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.managedV1alpha1, err = managedv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkV1alpha1, err = networkv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.objectV1alpha1, err = objectv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.routerV1alpha1, err = routerv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.serverV1alpha1, err = serverv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha1, err = storagev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tagV1alpha1, err = tagv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.firewallV1alpha1 = firewallv1alpha1.NewForConfigOrDie(c)
	cs.floatingV1alpha1 = floatingv1alpha1.NewForConfigOrDie(c)
	cs.managedV1alpha1 = managedv1alpha1.NewForConfigOrDie(c)
	cs.networkV1alpha1 = networkv1alpha1.NewForConfigOrDie(c)
	cs.objectV1alpha1 = objectv1alpha1.NewForConfigOrDie(c)
	cs.routerV1alpha1 = routerv1alpha1.NewForConfigOrDie(c)
	cs.serverV1alpha1 = serverv1alpha1.NewForConfigOrDie(c)
	cs.storageV1alpha1 = storagev1alpha1.NewForConfigOrDie(c)
	cs.tagV1alpha1 = tagv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.firewallV1alpha1 = firewallv1alpha1.New(c)
	cs.floatingV1alpha1 = floatingv1alpha1.New(c)
	cs.managedV1alpha1 = managedv1alpha1.New(c)
	cs.networkV1alpha1 = networkv1alpha1.New(c)
	cs.objectV1alpha1 = objectv1alpha1.New(c)
	cs.routerV1alpha1 = routerv1alpha1.New(c)
	cs.serverV1alpha1 = serverv1alpha1.New(c)
	cs.storageV1alpha1 = storagev1alpha1.New(c)
	cs.tagV1alpha1 = tagv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
