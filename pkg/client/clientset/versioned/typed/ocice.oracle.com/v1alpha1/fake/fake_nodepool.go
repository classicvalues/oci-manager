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
package fake

import (
	v1alpha1 "github.com/oracle/oci-manager/pkg/apis/ocice.oracle.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodePools implements NodePoolInterface
type FakeNodePools struct {
	Fake *FakeOciceV1alpha1
	ns   string
}

var nodepoolsResource = schema.GroupVersionResource{Group: "ocice.oracle.com", Version: "v1alpha1", Resource: "nodepools"}

var nodepoolsKind = schema.GroupVersionKind{Group: "ocice.oracle.com", Version: "v1alpha1", Kind: "NodePool"}

// Get takes name of the nodePool, and returns the corresponding nodePool object, and an error if there is any.
func (c *FakeNodePools) Get(name string, options v1.GetOptions) (result *v1alpha1.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nodepoolsResource, c.ns, name), &v1alpha1.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePool), err
}

// List takes label and field selectors, and returns the list of NodePools that match those selectors.
func (c *FakeNodePools) List(opts v1.ListOptions) (result *v1alpha1.NodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nodepoolsResource, nodepoolsKind, c.ns, opts), &v1alpha1.NodePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodePoolList{ListMeta: obj.(*v1alpha1.NodePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodePools.
func (c *FakeNodePools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nodepoolsResource, c.ns, opts))

}

// Create takes the representation of a nodePool and creates it.  Returns the server's representation of the nodePool, and an error, if there is any.
func (c *FakeNodePools) Create(nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nodepoolsResource, c.ns, nodePool), &v1alpha1.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePool), err
}

// Update takes the representation of a nodePool and updates it. Returns the server's representation of the nodePool, and an error, if there is any.
func (c *FakeNodePools) Update(nodePool *v1alpha1.NodePool) (result *v1alpha1.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nodepoolsResource, c.ns, nodePool), &v1alpha1.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePool), err
}

// Delete takes name of the nodePool and deletes it. Returns an error if one occurs.
func (c *FakeNodePools) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nodepoolsResource, c.ns, name), &v1alpha1.NodePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodePools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nodepoolsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodePoolList{})
	return err
}

// Patch applies the patch and returns the patched nodePool.
func (c *FakeNodePools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nodepoolsResource, c.ns, name, data, subresources...), &v1alpha1.NodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodePool), err
}
