/*
Copyright 2020 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/container-object-storage-interface-api/apis/objectstorage.k8s.io/v1alpha1"
)

// BucketAccessLister helps list BucketAccesses.
type BucketAccessLister interface {
	// List lists all BucketAccesses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.BucketAccess, err error)
	// Get retrieves the BucketAccess from the index for a given name.
	Get(name string) (*v1alpha1.BucketAccess, error)
	BucketAccessListerExpansion
}

// bucketAccessLister implements the BucketAccessLister interface.
type bucketAccessLister struct {
	indexer cache.Indexer
}

// NewBucketAccessLister returns a new BucketAccessLister.
func NewBucketAccessLister(indexer cache.Indexer) BucketAccessLister {
	return &bucketAccessLister{indexer: indexer}
}

// List lists all BucketAccesses in the indexer.
func (s *bucketAccessLister) List(selector labels.Selector) (ret []*v1alpha1.BucketAccess, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BucketAccess))
	})
	return ret, err
}

// Get retrieves the BucketAccess from the index for a given name.
func (s *bucketAccessLister) Get(name string) (*v1alpha1.BucketAccess, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bucketaccess"), name)
	}
	return obj.(*v1alpha1.BucketAccess), nil
}
