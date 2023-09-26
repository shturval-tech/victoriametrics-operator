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

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/shturval-tech/victoriametrics-operator/api/client/versioned/scheme"
	v1beta1 "github.com/shturval-tech/victoriametrics-operator/api/victoriametrics/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VMAuthsGetter has a method to return a VMAuthInterface.
// A group's client should implement this interface.
type VMAuthsGetter interface {
	VMAuths(namespace string) VMAuthInterface
}

// VMAuthInterface has methods to work with VMAuth resources.
type VMAuthInterface interface {
	Create(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.CreateOptions) (*v1beta1.VMAuth, error)
	Update(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.UpdateOptions) (*v1beta1.VMAuth, error)
	UpdateStatus(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.UpdateOptions) (*v1beta1.VMAuth, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VMAuth, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VMAuthList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAuth, err error)
	VMAuthExpansion
}

// vMAuths implements VMAuthInterface
type vMAuths struct {
	client rest.Interface
	ns     string
}

// newVMAuths returns a VMAuths
func newVMAuths(c *VictoriametricsV1beta1Client, namespace string) *vMAuths {
	return &vMAuths{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vMAuth, and returns the corresponding vMAuth object, and an error if there is any.
func (c *vMAuths) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMAuth, err error) {
	result = &v1beta1.VMAuth{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmauths").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VMAuths that match those selectors.
func (c *vMAuths) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMAuthList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.VMAuthList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vmauths").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vMAuths.
func (c *vMAuths) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vmauths").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vMAuth and creates it.  Returns the server's representation of the vMAuth, and an error, if there is any.
func (c *vMAuths) Create(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.CreateOptions) (result *v1beta1.VMAuth, err error) {
	result = &v1beta1.VMAuth{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vmauths").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMAuth).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vMAuth and updates it. Returns the server's representation of the vMAuth, and an error, if there is any.
func (c *vMAuths) Update(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.UpdateOptions) (result *v1beta1.VMAuth, err error) {
	result = &v1beta1.VMAuth{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmauths").
		Name(vMAuth.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMAuth).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vMAuths) UpdateStatus(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.UpdateOptions) (result *v1beta1.VMAuth, err error) {
	result = &v1beta1.VMAuth{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vmauths").
		Name(vMAuth.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vMAuth).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vMAuth and deletes it. Returns an error if one occurs.
func (c *vMAuths) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmauths").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vMAuths) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vmauths").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vMAuth.
func (c *vMAuths) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAuth, err error) {
	result = &v1beta1.VMAuth{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vmauths").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
