package repo

import (
	"context"
	"sync"

	"github.com/ozonmp/rcn-service-api/internal/model"
)

type NotFound struct {
}

func (e NotFound) Error() string {
	return "Service not found in repository"
}

// Repo is DAO for Service
type Repo interface {
	CreateService(ctx context.Context, service model.Service) (uint64, error)
	DescribeService(ctx context.Context, serviceID uint64) (*model.Service, error)
	ListService(ctx context.Context) ([]*model.Service, error)
	RemoveService(ctx context.Context, serviceID uint64) (bool, error)
}

type repo struct {
	i       uint64
	m       sync.Mutex
	storage map[uint64]model.Service
}

// NewRepo returns Repo interface
func NewRepo() Repo {
	return &repo{
		i: 3,
		storage: map[uint64]model.Service{
			1: {Id: 1, Title: "Test1", Description: "Description1", Rating: 1},
			2: {Id: 2, Title: "Test2", Description: "Description2", Rating: 2},
			3: {Id: 3, Title: "Test", Description: "Description3", Rating: 3},
		},
	}
}

func (r *repo) CreateService(ctx context.Context, service model.Service) (uint64, error) {
	defer r.m.Unlock()
	r.m.Lock()
	r.i++
	r.storage[r.i] = model.Service{Id: r.i, Title: service.Title, Description: service.Description, Rating: service.Rating}
	return r.i, nil
}

func (r *repo) DescribeService(ctx context.Context, serviceID uint64) (*model.Service, error) {
	defer r.m.Unlock()
	r.m.Lock()
	if val, ok := r.storage[serviceID]; ok {
		return &val, nil
	}
	return nil, NotFound{}
}

func (r *repo) ListService(ctx context.Context) ([]*model.Service, error) {
	defer r.m.Unlock()
	r.m.Lock()
	var result []*model.Service
	for _, v := range r.storage {
		copy := v
		result = append(result, &copy)
	}
	return result, nil
}

func (r *repo) RemoveService(ctx context.Context, serviceID uint64) (bool, error) {
	defer r.m.Unlock()
	r.m.Lock()
	if _, ok := r.storage[serviceID]; ok {
		delete(r.storage, serviceID)
		return true, nil
	}
	return false, NotFound{}
}
