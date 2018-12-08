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
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocicore.oracle.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VolumeBackupLister helps list VolumeBackups.
type VolumeBackupLister interface {
	// List lists all VolumeBackups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeBackup, err error)
	// VolumeBackups returns an object that can list and get VolumeBackups.
	VolumeBackups(namespace string) VolumeBackupNamespaceLister
	VolumeBackupListerExpansion
}

// volumeBackupLister implements the VolumeBackupLister interface.
type volumeBackupLister struct {
	indexer cache.Indexer
}

// NewVolumeBackupLister returns a new VolumeBackupLister.
func NewVolumeBackupLister(indexer cache.Indexer) VolumeBackupLister {
	return &volumeBackupLister{indexer: indexer}
}

// List lists all VolumeBackups in the indexer.
func (s *volumeBackupLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeBackup))
	})
	return ret, err
}

// VolumeBackups returns an object that can list and get VolumeBackups.
func (s *volumeBackupLister) VolumeBackups(namespace string) VolumeBackupNamespaceLister {
	return volumeBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeBackupNamespaceLister helps list and get VolumeBackups.
type VolumeBackupNamespaceLister interface {
	// List lists all VolumeBackups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VolumeBackup, err error)
	// Get retrieves the VolumeBackup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VolumeBackup, error)
	VolumeBackupNamespaceListerExpansion
}

// volumeBackupNamespaceLister implements the VolumeBackupNamespaceLister
// interface.
type volumeBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeBackups in the indexer for a given namespace.
func (s volumeBackupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VolumeBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VolumeBackup))
	})
	return ret, err
}

// Get retrieves the VolumeBackup from the indexer for a given namespace and name.
func (s volumeBackupNamespaceLister) Get(name string) (*v1alpha1.VolumeBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("volumebackup"), name)
	}
	return obj.(*v1alpha1.VolumeBackup), nil
}
