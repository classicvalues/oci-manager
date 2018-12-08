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
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/cloud.k8s.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageLister helps list Storages.
type StorageLister interface {
	// List lists all Storages in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Storage, err error)
	// Storages returns an object that can list and get Storages.
	Storages(namespace string) StorageNamespaceLister
	StorageListerExpansion
}

// storageLister implements the StorageLister interface.
type storageLister struct {
	indexer cache.Indexer
}

// NewStorageLister returns a new StorageLister.
func NewStorageLister(indexer cache.Indexer) StorageLister {
	return &storageLister{indexer: indexer}
}

// List lists all Storages in the indexer.
func (s *storageLister) List(selector labels.Selector) (ret []*v1alpha1.Storage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Storage))
	})
	return ret, err
}

// Storages returns an object that can list and get Storages.
func (s *storageLister) Storages(namespace string) StorageNamespaceLister {
	return storageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageNamespaceLister helps list and get Storages.
type StorageNamespaceLister interface {
	// List lists all Storages in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Storage, err error)
	// Get retrieves the Storage from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Storage, error)
	StorageNamespaceListerExpansion
}

// storageNamespaceLister implements the StorageNamespaceLister
// interface.
type storageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Storages in the indexer for a given namespace.
func (s storageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Storage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Storage))
	})
	return ret, err
}

// Get retrieves the Storage from the indexer for a given namespace and name.
func (s storageNamespaceLister) Get(name string) (*v1alpha1.Storage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storage"), name)
	}
	return obj.(*v1alpha1.Storage), nil
}
