/*


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

	v1beta1 "github.com/shturval-tech/victoriametrics-operator/api/victoriametrics/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVMUsers implements VMUserInterface
type FakeVMUsers struct {
	Fake *FakeVictoriametricsV1beta1
	ns   string
}

var vmusersResource = schema.GroupVersionResource{Group: "victoriametrics", Version: "v1beta1", Resource: "vmusers"}

var vmusersKind = schema.GroupVersionKind{Group: "victoriametrics", Version: "v1beta1", Kind: "VMUser"}

// Get takes name of the vMUser, and returns the corresponding vMUser object, and an error if there is any.
func (c *FakeVMUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vmusersResource, c.ns, name), &v1beta1.VMUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMUser), err
}

// List takes label and field selectors, and returns the list of VMUsers that match those selectors.
func (c *FakeVMUsers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMUserList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vmusersResource, vmusersKind, c.ns, opts), &v1beta1.VMUserList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VMUserList{ListMeta: obj.(*v1beta1.VMUserList).ListMeta}
	for _, item := range obj.(*v1beta1.VMUserList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vMUsers.
func (c *FakeVMUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vmusersResource, c.ns, opts))

}

// Create takes the representation of a vMUser and creates it.  Returns the server's representation of the vMUser, and an error, if there is any.
func (c *FakeVMUsers) Create(ctx context.Context, vMUser *v1beta1.VMUser, opts v1.CreateOptions) (result *v1beta1.VMUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vmusersResource, c.ns, vMUser), &v1beta1.VMUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMUser), err
}

// Update takes the representation of a vMUser and updates it. Returns the server's representation of the vMUser, and an error, if there is any.
func (c *FakeVMUsers) Update(ctx context.Context, vMUser *v1beta1.VMUser, opts v1.UpdateOptions) (result *v1beta1.VMUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vmusersResource, c.ns, vMUser), &v1beta1.VMUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMUser), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVMUsers) UpdateStatus(ctx context.Context, vMUser *v1beta1.VMUser, opts v1.UpdateOptions) (*v1beta1.VMUser, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vmusersResource, "status", c.ns, vMUser), &v1beta1.VMUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMUser), err
}

// Delete takes name of the vMUser and deletes it. Returns an error if one occurs.
func (c *FakeVMUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vmusersResource, c.ns, name, opts), &v1beta1.VMUser{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVMUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vmusersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VMUserList{})
	return err
}

// Patch applies the patch and returns the patched vMUser.
func (c *FakeVMUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMUser, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vmusersResource, c.ns, name, pt, data, subresources...), &v1beta1.VMUser{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMUser), err
}
