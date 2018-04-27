// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	garden "github.com/gardener/gardener/pkg/apis/garden"
	scheme "github.com/gardener/gardener/pkg/client/garden/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ShootsGetter has a method to return a ShootInterface.
// A group's client should implement this interface.
type ShootsGetter interface {
	Shoots(namespace string) ShootInterface
}

// ShootInterface has methods to work with Shoot resources.
type ShootInterface interface {
	Create(*garden.Shoot) (*garden.Shoot, error)
	Update(*garden.Shoot) (*garden.Shoot, error)
	UpdateStatus(*garden.Shoot) (*garden.Shoot, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*garden.Shoot, error)
	List(opts v1.ListOptions) (*garden.ShootList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.Shoot, err error)
	ShootExpansion
}

// shoots implements ShootInterface
type shoots struct {
	client rest.Interface
	ns     string
}

// newShoots returns a Shoots
func newShoots(c *GardenClient, namespace string) *shoots {
	return &shoots{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the shoot, and returns the corresponding shoot object, and an error if there is any.
func (c *shoots) Get(name string, options v1.GetOptions) (result *garden.Shoot, err error) {
	result = &garden.Shoot{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shoots").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Shoots that match those selectors.
func (c *shoots) List(opts v1.ListOptions) (result *garden.ShootList, err error) {
	result = &garden.ShootList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shoots.
func (c *shoots) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a shoot and creates it.  Returns the server's representation of the shoot, and an error, if there is any.
func (c *shoots) Create(shoot *garden.Shoot) (result *garden.Shoot, err error) {
	result = &garden.Shoot{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("shoots").
		Body(shoot).
		Do().
		Into(result)
	return
}

// Update takes the representation of a shoot and updates it. Returns the server's representation of the shoot, and an error, if there is any.
func (c *shoots) Update(shoot *garden.Shoot) (result *garden.Shoot, err error) {
	result = &garden.Shoot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shoots").
		Name(shoot.Name).
		Body(shoot).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *shoots) UpdateStatus(shoot *garden.Shoot) (result *garden.Shoot, err error) {
	result = &garden.Shoot{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shoots").
		Name(shoot.Name).
		SubResource("status").
		Body(shoot).
		Do().
		Into(result)
	return
}

// Delete takes name of the shoot and deletes it. Returns an error if one occurs.
func (c *shoots) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shoots").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shoots) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shoots").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched shoot.
func (c *shoots) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *garden.Shoot, err error) {
	result = &garden.Shoot{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("shoots").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}