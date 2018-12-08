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
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocice.oracle.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NodePoolLister helps list NodePools.
type NodePoolLister interface {
	// List lists all NodePools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NodePool, err error)
	// NodePools returns an object that can list and get NodePools.
	NodePools(namespace string) NodePoolNamespaceLister
	NodePoolListerExpansion
}

// nodePoolLister implements the NodePoolLister interface.
type nodePoolLister struct {
	indexer cache.Indexer
}

// NewNodePoolLister returns a new NodePoolLister.
func NewNodePoolLister(indexer cache.Indexer) NodePoolLister {
	return &nodePoolLister{indexer: indexer}
}

// List lists all NodePools in the indexer.
func (s *nodePoolLister) List(selector labels.Selector) (ret []*v1alpha1.NodePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodePool))
	})
	return ret, err
}

// NodePools returns an object that can list and get NodePools.
func (s *nodePoolLister) NodePools(namespace string) NodePoolNamespaceLister {
	return nodePoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NodePoolNamespaceLister helps list and get NodePools.
type NodePoolNamespaceLister interface {
	// List lists all NodePools in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NodePool, err error)
	// Get retrieves the NodePool from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NodePool, error)
	NodePoolNamespaceListerExpansion
}

// nodePoolNamespaceLister implements the NodePoolNamespaceLister
// interface.
type nodePoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NodePools in the indexer for a given namespace.
func (s nodePoolNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NodePool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NodePool))
	})
	return ret, err
}

// Get retrieves the NodePool from the indexer for a given namespace and name.
func (s nodePoolNamespaceLister) Get(name string) (*v1alpha1.NodePool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nodepool"), name)
	}
	return obj.(*v1alpha1.NodePool), nil
}
