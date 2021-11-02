package repo

import (
	"sort"
	"sync"

	"github.com/ozonmp/rcn-service-api/internal/model"
)

type EventRepo interface {
	Lock(n uint64) ([]model.ServiceEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.ServiceEvent) error
	Remove(eventIDs []uint64) error
}

type DbEventRepo struct {
	counter uint64
	storage map[uint64]*model.ServiceEvent
	m       sync.Mutex
}

func NewDbEventRepo(initMap map[uint64]*model.ServiceEvent) EventRepo {
	if initMap != nil {
		return &DbEventRepo{storage: initMap}
	} else {
		return &DbEventRepo{storage: make(map[uint64]*model.ServiceEvent)}
	}
}

func (db *DbEventRepo) Lock(n uint64) ([]model.ServiceEvent, error) {
	defer db.m.Unlock()
	var data = make([]model.ServiceEvent, 0, n)
	db.m.Lock()
	keys := make([]uint64, 0)
	for k, _ := range db.storage {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	for _, key := range keys {
		e := db.storage[key]
		if e.Status == model.Deferred {
			e.Status = model.Processed
			data = append(data, *e)
			n--
			if n == 0 {
				break
			}
		}
	}
	return data, nil
}

func (db *DbEventRepo) Unlock(eventIDs []uint64) error {
	defer db.m.Unlock()
	for _, eventID := range eventIDs {
		if event, ok := db.storage[eventID]; ok {
			event.Status = model.Deferred
		}
	}
	db.m.Lock()
	return nil
}

func (db *DbEventRepo) Add(event []model.ServiceEvent) error {
	defer db.m.Unlock()
	db.m.Lock()
	for _, e := range event {
		db.storage[e.ID] = &e
	}
	return nil
}
func (db *DbEventRepo) Remove(eventIDs []uint64) error {
	defer db.m.Unlock()
	db.m.Lock()
	for _, id := range eventIDs {
		delete(db.storage, id)
	}
	return nil

}
