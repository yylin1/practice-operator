/*
Copyright © 2018 Jack Lin (jacklin@lsalab.cs.nthu.edu.tw)

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

package v1

import (
	v1 "practice-operator/pkg/apis/student/v1"
	scheme "practice-operator/pkg/client/clientset/versioned/scheme"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// StudentsGetter has a method to return a StudentInterface.
// A group's client should implement this interface.
type StudentsGetter interface {
	Students(namespace string) StudentInterface
}

// StudentInterface has methods to work with Student resources.
type StudentInterface interface {
	Create(*v1.Student) (*v1.Student, error)
	Update(*v1.Student) (*v1.Student, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Student, error)
	List(opts meta_v1.ListOptions) (*v1.StudentList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Student, err error)
	StudentExpansion
}

// students implements StudentInterface
type students struct {
	client rest.Interface
	ns     string
}

// newStudents returns a Students
func newStudents(c *MyprojectV1Client, namespace string) *students {
	return &students{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the student, and returns the corresponding student object, and an error if there is any.
func (c *students) Get(name string, options meta_v1.GetOptions) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("students").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Students that match those selectors.
func (c *students) List(opts meta_v1.ListOptions) (result *v1.StudentList, err error) {
	result = &v1.StudentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested students.
func (c *students) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a student and creates it.  Returns the server's representation of the student, and an error, if there is any.
func (c *students) Create(student *v1.Student) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("students").
		Body(student).
		Do().
		Into(result)
	return
}

// Update takes the representation of a student and updates it. Returns the server's representation of the student, and an error, if there is any.
func (c *students) Update(student *v1.Student) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("students").
		Name(student.Name).
		Body(student).
		Do().
		Into(result)
	return
}

// Delete takes name of the student and deletes it. Returns an error if one occurs.
func (c *students) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("students").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *students) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("students").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched student.
func (c *students) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Student, err error) {
	result = &v1.Student{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("students").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
