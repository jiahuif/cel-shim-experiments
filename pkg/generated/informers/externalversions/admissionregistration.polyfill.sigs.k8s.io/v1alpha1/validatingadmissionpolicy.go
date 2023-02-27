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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	admissionregistrationpolyfillsigsk8siov1alpha1 "github.com/jiahuif/cel-shim-experiments/pkg/apis/admissionregistration.polyfill.sigs.k8s.io/v1alpha1"
	versioned "github.com/jiahuif/cel-shim-experiments/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/jiahuif/cel-shim-experiments/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/jiahuif/cel-shim-experiments/pkg/generated/listers/admissionregistration.polyfill.sigs.k8s.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ValidatingAdmissionPolicyInformer provides access to a shared informer and lister for
// ValidatingAdmissionPolicies.
type ValidatingAdmissionPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ValidatingAdmissionPolicyLister
}

type validatingAdmissionPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewValidatingAdmissionPolicyInformer constructs a new informer for ValidatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewValidatingAdmissionPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredValidatingAdmissionPolicyInformer constructs a new informer for ValidatingAdmissionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredValidatingAdmissionPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1alpha1().ValidatingAdmissionPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AdmissionregistrationV1alpha1().ValidatingAdmissionPolicies().Watch(context.TODO(), options)
			},
		},
		&admissionregistrationpolyfillsigsk8siov1alpha1.ValidatingAdmissionPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *validatingAdmissionPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredValidatingAdmissionPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *validatingAdmissionPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&admissionregistrationpolyfillsigsk8siov1alpha1.ValidatingAdmissionPolicy{}, f.defaultInformer)
}

func (f *validatingAdmissionPolicyInformer) Lister() v1alpha1.ValidatingAdmissionPolicyLister {
	return v1alpha1.NewValidatingAdmissionPolicyLister(f.Informer().GetIndexer())
}
