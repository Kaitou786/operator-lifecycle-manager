/*
Copyright 2018 CoreOS, Inc.

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
	v1alpha1 "github.com/coreos-inc/alm/pkg/apis/catalogsource/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCatalogSources implements CatalogSourceInterface
type FakeCatalogSources struct {
	Fake *FakeCatalogsourceV1alpha1
	ns   string
}

var catalogsourcesResource = schema.GroupVersionResource{Group: "catalogsource-v1", Version: "v1alpha1", Resource: "catalogsource-v1s"}

var catalogsourcesKind = schema.GroupVersionKind{Group: "catalogsource-v1", Version: "v1alpha1", Kind: "CatalogSource"}

// Get takes name of the catalogSource, and returns the corresponding catalogSource object, and an error if there is any.
func (c *FakeCatalogSources) Get(name string, options v1.GetOptions) (result *v1alpha1.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogsourcesResource, c.ns, name), &v1alpha1.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogSource), err
}

// List takes label and field selectors, and returns the list of CatalogSources that match those selectors.
func (c *FakeCatalogSources) List(opts v1.ListOptions) (result *v1alpha1.CatalogSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogsourcesResource, catalogsourcesKind, c.ns, opts), &v1alpha1.CatalogSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CatalogSourceList{}
	for _, item := range obj.(*v1alpha1.CatalogSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogSources.
func (c *FakeCatalogSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogsourcesResource, c.ns, opts))

}

// Create takes the representation of a catalogSource and creates it.  Returns the server's representation of the catalogSource, and an error, if there is any.
func (c *FakeCatalogSources) Create(catalogSource *v1alpha1.CatalogSource) (result *v1alpha1.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogsourcesResource, c.ns, catalogSource), &v1alpha1.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogSource), err
}

// Update takes the representation of a catalogSource and updates it. Returns the server's representation of the catalogSource, and an error, if there is any.
func (c *FakeCatalogSources) Update(catalogSource *v1alpha1.CatalogSource) (result *v1alpha1.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogsourcesResource, c.ns, catalogSource), &v1alpha1.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCatalogSources) UpdateStatus(catalogSource *v1alpha1.CatalogSource) (*v1alpha1.CatalogSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(catalogsourcesResource, "status", c.ns, catalogSource), &v1alpha1.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogSource), err
}

// Delete takes name of the catalogSource and deletes it. Returns an error if one occurs.
func (c *FakeCatalogSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(catalogsourcesResource, c.ns, name), &v1alpha1.CatalogSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogsourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.CatalogSourceList{})
	return err
}

// Patch applies the patch and returns the patched catalogSource.
func (c *FakeCatalogSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.CatalogSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogsourcesResource, c.ns, name, data, subresources...), &v1alpha1.CatalogSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CatalogSource), err
}
