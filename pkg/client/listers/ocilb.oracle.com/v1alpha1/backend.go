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

// BackendLister helps list Backends.
type BackendLister interface {
	// List lists all Backends in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Backend, err error)
	// Backends returns an object that can list and get Backends.
	Backends(namespace string) BackendNamespaceLister
	BackendListerExpansion
}

// backendLister implements the BackendLister interface.
type backendLister struct {
	indexer cache.Indexer
}

// NewBackendLister returns a new BackendLister.
func NewBackendLister(indexer cache.Indexer) BackendLister {
	return &backendLister{indexer: indexer}
}

// List lists all Backends in the indexer.
func (s *backendLister) List(selector labels.Selector) (ret []*v1alpha1.Backend, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Backend))
	})
	return ret, err
}

// Backends returns an object that can list and get Backends.
func (s *backendLister) Backends(namespace string) BackendNamespaceLister {
	return backendNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackendNamespaceLister helps list and get Backends.
type BackendNamespaceLister interface {
	// List lists all Backends in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Backend, err error)
	// Get retrieves the Backend from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Backend, error)
	BackendNamespaceListerExpansion
}

// backendNamespaceLister implements the BackendNamespaceLister
// interface.
type backendNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Backends in the indexer for a given namespace.
func (s backendNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Backend, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Backend))
	})
	return ret, err
}

// Get retrieves the Backend from the indexer for a given namespace and name.
func (s backendNamespaceLister) Get(name string) (*v1alpha1.Backend, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backend"), name)
	}
	return obj.(*v1alpha1.Backend), nil
}
