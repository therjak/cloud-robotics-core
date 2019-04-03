// Copyright 2019 The Cloud Robotics Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/googlecloudrobotics/core/src/go/pkg/apis/apps/v1alpha1"
	scheme "github.com/googlecloudrobotics/core/src/go/pkg/client/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceSetsGetter has a method to return a ResourceSetInterface.
// A group's client should implement this interface.
type ResourceSetsGetter interface {
	ResourceSets() ResourceSetInterface
}

// ResourceSetInterface has methods to work with ResourceSet resources.
type ResourceSetInterface interface {
	Create(*v1alpha1.ResourceSet) (*v1alpha1.ResourceSet, error)
	Update(*v1alpha1.ResourceSet) (*v1alpha1.ResourceSet, error)
	UpdateStatus(*v1alpha1.ResourceSet) (*v1alpha1.ResourceSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ResourceSet, error)
	List(opts v1.ListOptions) (*v1alpha1.ResourceSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceSet, err error)
	ResourceSetExpansion
}

// resourceSets implements ResourceSetInterface
type resourceSets struct {
	client rest.Interface
}

// newResourceSets returns a ResourceSets
func newResourceSets(c *AppsV1alpha1Client) *resourceSets {
	return &resourceSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the resourceSet, and returns the corresponding resourceSet object, and an error if there is any.
func (c *resourceSets) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceSet, err error) {
	result = &v1alpha1.ResourceSet{}
	err = c.client.Get().
		Resource("resourcesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceSets that match those selectors.
func (c *resourceSets) List(opts v1.ListOptions) (result *v1alpha1.ResourceSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceSetList{}
	err = c.client.Get().
		Resource("resourcesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceSets.
func (c *resourceSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("resourcesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a resourceSet and creates it.  Returns the server's representation of the resourceSet, and an error, if there is any.
func (c *resourceSets) Create(resourceSet *v1alpha1.ResourceSet) (result *v1alpha1.ResourceSet, err error) {
	result = &v1alpha1.ResourceSet{}
	err = c.client.Post().
		Resource("resourcesets").
		Body(resourceSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a resourceSet and updates it. Returns the server's representation of the resourceSet, and an error, if there is any.
func (c *resourceSets) Update(resourceSet *v1alpha1.ResourceSet) (result *v1alpha1.ResourceSet, err error) {
	result = &v1alpha1.ResourceSet{}
	err = c.client.Put().
		Resource("resourcesets").
		Name(resourceSet.Name).
		Body(resourceSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *resourceSets) UpdateStatus(resourceSet *v1alpha1.ResourceSet) (result *v1alpha1.ResourceSet, err error) {
	result = &v1alpha1.ResourceSet{}
	err = c.client.Put().
		Resource("resourcesets").
		Name(resourceSet.Name).
		SubResource("status").
		Body(resourceSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the resourceSet and deletes it. Returns an error if one occurs.
func (c *resourceSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("resourcesets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("resourcesets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched resourceSet.
func (c *resourceSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceSet, err error) {
	result = &v1alpha1.ResourceSet{}
	err = c.client.Patch(pt).
		Resource("resourcesets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
