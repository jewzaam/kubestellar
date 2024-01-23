/*
Copyright The KubeStellar Authors.

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

	v1alpha1 "github.com/kubestellar/kubestellar/api/edge/v1alpha1"
)

// PlacementDecisionLister helps list PlacementDecisions.
// All objects returned here must be treated as read-only.
type PlacementDecisionLister interface {
	// List lists all PlacementDecisions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PlacementDecision, err error)
	// Get retrieves the PlacementDecision from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PlacementDecision, error)
	PlacementDecisionListerExpansion
}

// placementDecisionLister implements the PlacementDecisionLister interface.
type placementDecisionLister struct {
	indexer cache.Indexer
}

// NewPlacementDecisionLister returns a new PlacementDecisionLister.
func NewPlacementDecisionLister(indexer cache.Indexer) PlacementDecisionLister {
	return &placementDecisionLister{indexer: indexer}
}

// List lists all PlacementDecisions in the indexer.
func (s *placementDecisionLister) List(selector labels.Selector) (ret []*v1alpha1.PlacementDecision, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PlacementDecision))
	})
	return ret, err
}

// Get retrieves the PlacementDecision from the index for a given name.
func (s *placementDecisionLister) Get(name string) (*v1alpha1.PlacementDecision, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("placementdecision"), name)
	}
	return obj.(*v1alpha1.PlacementDecision), nil
}
