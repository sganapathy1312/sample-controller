/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/healthmonitor/v1alpha1"
	scheme "k8s.io/sample-controller/pkg/generated/clientset/versioned/scheme"
)

// DeploymentHealthMonitorsGetter has a method to return a DeploymentHealthMonitorInterface.
// A group's client should implement this interface.
type DeploymentHealthMonitorsGetter interface {
	DeploymentHealthMonitors(namespace string) DeploymentHealthMonitorInterface
}

// DeploymentHealthMonitorInterface has methods to work with DeploymentHealthMonitor resources.
type DeploymentHealthMonitorInterface interface {
	Create(ctx context.Context, deploymentHealthMonitor *v1alpha1.DeploymentHealthMonitor, opts v1.CreateOptions) (*v1alpha1.DeploymentHealthMonitor, error)
	Update(ctx context.Context, deploymentHealthMonitor *v1alpha1.DeploymentHealthMonitor, opts v1.UpdateOptions) (*v1alpha1.DeploymentHealthMonitor, error)
	UpdateStatus(ctx context.Context, deploymentHealthMonitor *v1alpha1.DeploymentHealthMonitor, opts v1.UpdateOptions) (*v1alpha1.DeploymentHealthMonitor, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DeploymentHealthMonitor, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DeploymentHealthMonitorList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentHealthMonitor, err error)
	DeploymentHealthMonitorExpansion
}

// deploymentHealthMonitors implements DeploymentHealthMonitorInterface
type deploymentHealthMonitors struct {
	client rest.Interface
	ns     string
}

// newDeploymentHealthMonitors returns a DeploymentHealthMonitors
func newDeploymentHealthMonitors(c *HealthmonitorV1alpha1Client, namespace string) *deploymentHealthMonitors {
	return &deploymentHealthMonitors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deploymentHealthMonitor, and returns the corresponding deploymentHealthMonitor object, and an error if there is any.
func (c *deploymentHealthMonitors) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DeploymentHealthMonitor, err error) {
	result = &v1alpha1.DeploymentHealthMonitor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeploymentHealthMonitors that match those selectors.
func (c *deploymentHealthMonitors) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DeploymentHealthMonitorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeploymentHealthMonitorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deploymentHealthMonitors.
func (c *deploymentHealthMonitors) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deploymentHealthMonitor and creates it.  Returns the server's representation of the deploymentHealthMonitor, and an error, if there is any.
func (c *deploymentHealthMonitors) Create(ctx context.Context, deploymentHealthMonitor *v1alpha1.DeploymentHealthMonitor, opts v1.CreateOptions) (result *v1alpha1.DeploymentHealthMonitor, err error) {
	result = &v1alpha1.DeploymentHealthMonitor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentHealthMonitor).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deploymentHealthMonitor and updates it. Returns the server's representation of the deploymentHealthMonitor, and an error, if there is any.
func (c *deploymentHealthMonitors) Update(ctx context.Context, deploymentHealthMonitor *v1alpha1.DeploymentHealthMonitor, opts v1.UpdateOptions) (result *v1alpha1.DeploymentHealthMonitor, err error) {
	result = &v1alpha1.DeploymentHealthMonitor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		Name(deploymentHealthMonitor.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentHealthMonitor).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deploymentHealthMonitors) UpdateStatus(ctx context.Context, deploymentHealthMonitor *v1alpha1.DeploymentHealthMonitor, opts v1.UpdateOptions) (result *v1alpha1.DeploymentHealthMonitor, err error) {
	result = &v1alpha1.DeploymentHealthMonitor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		Name(deploymentHealthMonitor.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(deploymentHealthMonitor).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deploymentHealthMonitor and deletes it. Returns an error if one occurs.
func (c *deploymentHealthMonitors) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deploymentHealthMonitors) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deploymentHealthMonitor.
func (c *deploymentHealthMonitors) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DeploymentHealthMonitor, err error) {
	result = &v1alpha1.DeploymentHealthMonitor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deploymenthealthmonitors").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
