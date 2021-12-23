// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/ShootingRange/csgo-operator/pkg/apis/csgooperator/v1alpha1"
	scheme "github.com/ShootingRange/csgo-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CSGOServersGetter has a method to return a CSGOServerInterface.
// A group's client should implement this interface.
type CSGOServersGetter interface {
	CSGOServers(namespace string) CSGOServerInterface
}

// CSGOServerInterface has methods to work with CSGOServer resources.
type CSGOServerInterface interface {
	Create(ctx context.Context, cSGOServer *v1alpha1.CSGOServer, opts v1.CreateOptions) (*v1alpha1.CSGOServer, error)
	Update(ctx context.Context, cSGOServer *v1alpha1.CSGOServer, opts v1.UpdateOptions) (*v1alpha1.CSGOServer, error)
	UpdateStatus(ctx context.Context, cSGOServer *v1alpha1.CSGOServer, opts v1.UpdateOptions) (*v1alpha1.CSGOServer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CSGOServer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CSGOServerList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CSGOServer, err error)
	CSGOServerExpansion
}

// cSGOServers implements CSGOServerInterface
type cSGOServers struct {
	client rest.Interface
	ns     string
}

// newCSGOServers returns a CSGOServers
func newCSGOServers(c *CsgooperatorV1alpha1Client, namespace string) *cSGOServers {
	return &cSGOServers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cSGOServer, and returns the corresponding cSGOServer object, and an error if there is any.
func (c *cSGOServers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CSGOServer, err error) {
	result = &v1alpha1.CSGOServer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("csgoservers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSGOServers that match those selectors.
func (c *cSGOServers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CSGOServerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CSGOServerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("csgoservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSGOServers.
func (c *cSGOServers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("csgoservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cSGOServer and creates it.  Returns the server's representation of the cSGOServer, and an error, if there is any.
func (c *cSGOServers) Create(ctx context.Context, cSGOServer *v1alpha1.CSGOServer, opts v1.CreateOptions) (result *v1alpha1.CSGOServer, err error) {
	result = &v1alpha1.CSGOServer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("csgoservers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSGOServer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cSGOServer and updates it. Returns the server's representation of the cSGOServer, and an error, if there is any.
func (c *cSGOServers) Update(ctx context.Context, cSGOServer *v1alpha1.CSGOServer, opts v1.UpdateOptions) (result *v1alpha1.CSGOServer, err error) {
	result = &v1alpha1.CSGOServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("csgoservers").
		Name(cSGOServer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSGOServer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cSGOServers) UpdateStatus(ctx context.Context, cSGOServer *v1alpha1.CSGOServer, opts v1.UpdateOptions) (result *v1alpha1.CSGOServer, err error) {
	result = &v1alpha1.CSGOServer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("csgoservers").
		Name(cSGOServer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cSGOServer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cSGOServer and deletes it. Returns an error if one occurs.
func (c *cSGOServers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("csgoservers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSGOServers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("csgoservers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cSGOServer.
func (c *cSGOServers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CSGOServer, err error) {
	result = &v1alpha1.CSGOServer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("csgoservers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
