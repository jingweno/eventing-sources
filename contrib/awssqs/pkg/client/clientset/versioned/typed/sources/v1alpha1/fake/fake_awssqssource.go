/*
Copyright 2018 The Knative Authors

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/knative/eventing-sources/contrib/awssqs/pkg/apis/sources/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAwsSqsSources implements AwsSqsSourceInterface
type FakeAwsSqsSources struct {
	Fake *FakeSourcesV1alpha1
	ns   string
}

var awssqssourcesResource = schema.GroupVersionResource{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Resource: "awssqssources"}

var awssqssourcesKind = schema.GroupVersionKind{Group: "sources.eventing.knative.dev", Version: "v1alpha1", Kind: "AwsSqsSource"}

// Get takes name of the awsSqsSource, and returns the corresponding awsSqsSource object, and an error if there is any.
func (c *FakeAwsSqsSources) Get(name string, options v1.GetOptions) (result *v1alpha1.AwsSqsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awssqssourcesResource, c.ns, name), &v1alpha1.AwsSqsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsSource), err
}

// List takes label and field selectors, and returns the list of AwsSqsSources that match those selectors.
func (c *FakeAwsSqsSources) List(opts v1.ListOptions) (result *v1alpha1.AwsSqsSourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awssqssourcesResource, awssqssourcesKind, c.ns, opts), &v1alpha1.AwsSqsSourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AwsSqsSourceList{ListMeta: obj.(*v1alpha1.AwsSqsSourceList).ListMeta}
	for _, item := range obj.(*v1alpha1.AwsSqsSourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested awsSqsSources.
func (c *FakeAwsSqsSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awssqssourcesResource, c.ns, opts))

}

// Create takes the representation of a awsSqsSource and creates it.  Returns the server's representation of the awsSqsSource, and an error, if there is any.
func (c *FakeAwsSqsSources) Create(awsSqsSource *v1alpha1.AwsSqsSource) (result *v1alpha1.AwsSqsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awssqssourcesResource, c.ns, awsSqsSource), &v1alpha1.AwsSqsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsSource), err
}

// Update takes the representation of a awsSqsSource and updates it. Returns the server's representation of the awsSqsSource, and an error, if there is any.
func (c *FakeAwsSqsSources) Update(awsSqsSource *v1alpha1.AwsSqsSource) (result *v1alpha1.AwsSqsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awssqssourcesResource, c.ns, awsSqsSource), &v1alpha1.AwsSqsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsSource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAwsSqsSources) UpdateStatus(awsSqsSource *v1alpha1.AwsSqsSource) (*v1alpha1.AwsSqsSource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awssqssourcesResource, "status", c.ns, awsSqsSource), &v1alpha1.AwsSqsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsSource), err
}

// Delete takes name of the awsSqsSource and deletes it. Returns an error if one occurs.
func (c *FakeAwsSqsSources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awssqssourcesResource, c.ns, name), &v1alpha1.AwsSqsSource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAwsSqsSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awssqssourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.AwsSqsSourceList{})
	return err
}

// Patch applies the patch and returns the patched awsSqsSource.
func (c *FakeAwsSqsSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AwsSqsSource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awssqssourcesResource, c.ns, name, pt, data, subresources...), &v1alpha1.AwsSqsSource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AwsSqsSource), err
}
