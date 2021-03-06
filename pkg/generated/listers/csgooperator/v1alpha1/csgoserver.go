// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/ShootingRange/csgo-operator/pkg/apis/csgooperator/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CSGOServerLister helps list CSGOServers.
// All objects returned here must be treated as read-only.
type CSGOServerLister interface {
	// List lists all CSGOServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CSGOServer, err error)
	// CSGOServers returns an object that can list and get CSGOServers.
	CSGOServers(namespace string) CSGOServerNamespaceLister
	CSGOServerListerExpansion
}

// cSGOServerLister implements the CSGOServerLister interface.
type cSGOServerLister struct {
	indexer cache.Indexer
}

// NewCSGOServerLister returns a new CSGOServerLister.
func NewCSGOServerLister(indexer cache.Indexer) CSGOServerLister {
	return &cSGOServerLister{indexer: indexer}
}

// List lists all CSGOServers in the indexer.
func (s *cSGOServerLister) List(selector labels.Selector) (ret []*v1alpha1.CSGOServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CSGOServer))
	})
	return ret, err
}

// CSGOServers returns an object that can list and get CSGOServers.
func (s *cSGOServerLister) CSGOServers(namespace string) CSGOServerNamespaceLister {
	return cSGOServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CSGOServerNamespaceLister helps list and get CSGOServers.
// All objects returned here must be treated as read-only.
type CSGOServerNamespaceLister interface {
	// List lists all CSGOServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CSGOServer, err error)
	// Get retrieves the CSGOServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CSGOServer, error)
	CSGOServerNamespaceListerExpansion
}

// cSGOServerNamespaceLister implements the CSGOServerNamespaceLister
// interface.
type cSGOServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CSGOServers in the indexer for a given namespace.
func (s cSGOServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CSGOServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CSGOServer))
	})
	return ret, err
}

// Get retrieves the CSGOServer from the indexer for a given namespace and name.
func (s cSGOServerNamespaceLister) Get(name string) (*v1alpha1.CSGOServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("csgoserver"), name)
	}
	return obj.(*v1alpha1.CSGOServer), nil
}
