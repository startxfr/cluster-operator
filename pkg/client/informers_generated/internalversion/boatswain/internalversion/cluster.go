/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by informer-gen

package internalversion

import (
	boatswain "github.com/staebler/boatswain/pkg/apis/boatswain"
	internalclientset "github.com/staebler/boatswain/pkg/client/clientset_generated/internalclientset"
	internalinterfaces "github.com/staebler/boatswain/pkg/client/informers_generated/internalversion/internalinterfaces"
	internalversion "github.com/staebler/boatswain/pkg/client/listers_generated/boatswain/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterInformer provides access to a shared informer and lister for
// Clusters.
type ClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.ClusterLister
}

type clusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewClusterInformer constructs a new informer for Cluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterInformer(client internalclientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Boatswain().Clusters(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Boatswain().Clusters(namespace).Watch(options)
			},
		},
		&boatswain.Cluster{},
		resyncPeriod,
		indexers,
	)
}

func defaultClusterInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewClusterInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *clusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&boatswain.Cluster{}, defaultClusterInformer)
}

func (f *clusterInformer) Lister() internalversion.ClusterLister {
	return internalversion.NewClusterLister(f.Informer().GetIndexer())
}
