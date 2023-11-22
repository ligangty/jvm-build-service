/*
Copyright 2021-2022 Red Hat, Inc.

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
	"context"

	v1alpha1 "github.com/redhat-appstudio/jvm-build-service/pkg/apis/jvmbuildservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJvmImageScans implements JvmImageScanInterface
type FakeJvmImageScans struct {
	Fake *FakeJvmbuildserviceV1alpha1
	ns   string
}

var jvmimagescansResource = schema.GroupVersionResource{Group: "jvmbuildservice.io", Version: "v1alpha1", Resource: "jvmimagescans"}

var jvmimagescansKind = schema.GroupVersionKind{Group: "jvmbuildservice.io", Version: "v1alpha1", Kind: "JvmImageScan"}

// Get takes name of the jvmImageScan, and returns the corresponding jvmImageScan object, and an error if there is any.
func (c *FakeJvmImageScans) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.JvmImageScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jvmimagescansResource, c.ns, name), &v1alpha1.JvmImageScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JvmImageScan), err
}

// List takes label and field selectors, and returns the list of JvmImageScans that match those selectors.
func (c *FakeJvmImageScans) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.JvmImageScanList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jvmimagescansResource, jvmimagescansKind, c.ns, opts), &v1alpha1.JvmImageScanList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.JvmImageScanList{ListMeta: obj.(*v1alpha1.JvmImageScanList).ListMeta}
	for _, item := range obj.(*v1alpha1.JvmImageScanList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jvmImageScans.
func (c *FakeJvmImageScans) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jvmimagescansResource, c.ns, opts))

}

// Create takes the representation of a jvmImageScan and creates it.  Returns the server's representation of the jvmImageScan, and an error, if there is any.
func (c *FakeJvmImageScans) Create(ctx context.Context, jvmImageScan *v1alpha1.JvmImageScan, opts v1.CreateOptions) (result *v1alpha1.JvmImageScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jvmimagescansResource, c.ns, jvmImageScan), &v1alpha1.JvmImageScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JvmImageScan), err
}

// Update takes the representation of a jvmImageScan and updates it. Returns the server's representation of the jvmImageScan, and an error, if there is any.
func (c *FakeJvmImageScans) Update(ctx context.Context, jvmImageScan *v1alpha1.JvmImageScan, opts v1.UpdateOptions) (result *v1alpha1.JvmImageScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jvmimagescansResource, c.ns, jvmImageScan), &v1alpha1.JvmImageScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JvmImageScan), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJvmImageScans) UpdateStatus(ctx context.Context, jvmImageScan *v1alpha1.JvmImageScan, opts v1.UpdateOptions) (*v1alpha1.JvmImageScan, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jvmimagescansResource, "status", c.ns, jvmImageScan), &v1alpha1.JvmImageScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JvmImageScan), err
}

// Delete takes name of the jvmImageScan and deletes it. Returns an error if one occurs.
func (c *FakeJvmImageScans) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(jvmimagescansResource, c.ns, name, opts), &v1alpha1.JvmImageScan{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJvmImageScans) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jvmimagescansResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.JvmImageScanList{})
	return err
}

// Patch applies the patch and returns the patched jvmImageScan.
func (c *FakeJvmImageScans) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.JvmImageScan, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jvmimagescansResource, c.ns, name, pt, data, subresources...), &v1alpha1.JvmImageScan{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.JvmImageScan), err
}