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
	v1alpha1 "github.com/coreos-inc/alm/pkg/apis/clusterserviceversion/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterServiceVersions implements ClusterServiceVersionInterface
type FakeClusterServiceVersions struct {
	Fake *FakeClusterserviceversionV1alpha1
	ns   string
}

var clusterserviceversionsResource = schema.GroupVersionResource{Group: "clusterserviceversion-v1", Version: "v1alpha1", Resource: "clusterserviceversion-v1s"}

var clusterserviceversionsKind = schema.GroupVersionKind{Group: "clusterserviceversion-v1", Version: "v1alpha1", Kind: "ClusterServiceVersion"}

// Get takes name of the clusterServiceVersion, and returns the corresponding clusterServiceVersion object, and an error if there is any.
func (c *FakeClusterServiceVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterServiceVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterserviceversionsResource, c.ns, name), &v1alpha1.ClusterServiceVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServiceVersion), err
}

// List takes label and field selectors, and returns the list of ClusterServiceVersions that match those selectors.
func (c *FakeClusterServiceVersions) List(opts v1.ListOptions) (result *v1alpha1.ClusterServiceVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterserviceversionsResource, clusterserviceversionsKind, c.ns, opts), &v1alpha1.ClusterServiceVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterServiceVersionList{}
	for _, item := range obj.(*v1alpha1.ClusterServiceVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterServiceVersions.
func (c *FakeClusterServiceVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterserviceversionsResource, c.ns, opts))

}

// Create takes the representation of a clusterServiceVersion and creates it.  Returns the server's representation of the clusterServiceVersion, and an error, if there is any.
func (c *FakeClusterServiceVersions) Create(clusterServiceVersion *v1alpha1.ClusterServiceVersion) (result *v1alpha1.ClusterServiceVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterserviceversionsResource, c.ns, clusterServiceVersion), &v1alpha1.ClusterServiceVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServiceVersion), err
}

// Update takes the representation of a clusterServiceVersion and updates it. Returns the server's representation of the clusterServiceVersion, and an error, if there is any.
func (c *FakeClusterServiceVersions) Update(clusterServiceVersion *v1alpha1.ClusterServiceVersion) (result *v1alpha1.ClusterServiceVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterserviceversionsResource, c.ns, clusterServiceVersion), &v1alpha1.ClusterServiceVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServiceVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterServiceVersions) UpdateStatus(clusterServiceVersion *v1alpha1.ClusterServiceVersion) (*v1alpha1.ClusterServiceVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterserviceversionsResource, "status", c.ns, clusterServiceVersion), &v1alpha1.ClusterServiceVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServiceVersion), err
}

// Delete takes name of the clusterServiceVersion and deletes it. Returns an error if one occurs.
func (c *FakeClusterServiceVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterserviceversionsResource, c.ns, name), &v1alpha1.ClusterServiceVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterServiceVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterserviceversionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterServiceVersionList{})
	return err
}

// Patch applies the patch and returns the patched clusterServiceVersion.
func (c *FakeClusterServiceVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterServiceVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterserviceversionsResource, c.ns, name, data, subresources...), &v1alpha1.ClusterServiceVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterServiceVersion), err
}
