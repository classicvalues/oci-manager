/*
Copyright 2018 Oracle and/or its affiliates. All rights reserved.

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
package v1alpha1

import (
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocilb.oracle.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ListenerLister helps list Listeners.
type ListenerLister interface {
	// List lists all Listeners in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Listener, err error)
	// Listeners returns an object that can list and get Listeners.
	Listeners(namespace string) ListenerNamespaceLister
	ListenerListerExpansion
}

// listenerLister implements the ListenerLister interface.
type listenerLister struct {
	indexer cache.Indexer
}

// NewListenerLister returns a new ListenerLister.
func NewListenerLister(indexer cache.Indexer) ListenerLister {
	return &listenerLister{indexer: indexer}
}

// List lists all Listeners in the indexer.
func (s *listenerLister) List(selector labels.Selector) (ret []*v1alpha1.Listener, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Listener))
	})
	return ret, err
}

// Listeners returns an object that can list and get Listeners.
func (s *listenerLister) Listeners(namespace string) ListenerNamespaceLister {
	return listenerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ListenerNamespaceLister helps list and get Listeners.
type ListenerNamespaceLister interface {
	// List lists all Listeners in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Listener, err error)
	// Get retrieves the Listener from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Listener, error)
	ListenerNamespaceListerExpansion
}

// listenerNamespaceLister implements the ListenerNamespaceLister
// interface.
type listenerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Listeners in the indexer for a given namespace.
func (s listenerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Listener, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Listener))
	})
	return ret, err
}

// Get retrieves the Listener from the indexer for a given namespace and name.
func (s listenerNamespaceLister) Get(name string) (*v1alpha1.Listener, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("listener"), name)
	}
	return obj.(*v1alpha1.Listener), nil
}
