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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	managedv1alpha1 "kubeform.dev/provider-upcloud-api/apis/managed/v1alpha1"
	versioned "kubeform.dev/provider-upcloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-upcloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-upcloud-api/client/listers/managed/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DatabasePostgresqlInformer provides access to a shared informer and lister for
// DatabasePostgresqls.
type DatabasePostgresqlInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.DatabasePostgresqlLister
}

type databasePostgresqlInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDatabasePostgresqlInformer constructs a new informer for DatabasePostgresql type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDatabasePostgresqlInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDatabasePostgresqlInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDatabasePostgresqlInformer constructs a new informer for DatabasePostgresql type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDatabasePostgresqlInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagedV1alpha1().DatabasePostgresqls(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ManagedV1alpha1().DatabasePostgresqls(namespace).Watch(context.TODO(), options)
			},
		},
		&managedv1alpha1.DatabasePostgresql{},
		resyncPeriod,
		indexers,
	)
}

func (f *databasePostgresqlInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDatabasePostgresqlInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *databasePostgresqlInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&managedv1alpha1.DatabasePostgresql{}, f.defaultInformer)
}

func (f *databasePostgresqlInformer) Lister() v1alpha1.DatabasePostgresqlLister {
	return v1alpha1.NewDatabasePostgresqlLister(f.Informer().GetIndexer())
}
