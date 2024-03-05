package db

import (
	"sync"
)

var (
	db *InMemoryDB
)

type InMemoryDB struct {
	packSizes []int
	lock      sync.RWMutex
}

type PacksInMemoryDB interface {
	UpdatePacks(newPacks []int) error
	GetPackSizes() []int
}

func New() (m *InMemoryDB) {
	if db != nil {
		return db
	}
	db = &InMemoryDB{packSizes: []int{250, 500, 1000, 2000, 5000}}
	db.lock = sync.RWMutex{}
	return db
}

func (d *InMemoryDB) UpdatePacks(newPacks []int) error {
	d.set(newPacks)
	return nil
}

func (d *InMemoryDB) GetPackSizes() []int {
	return d.get()
}

func (d *InMemoryDB) get() []int {
	defer db.lock.RUnlock()
	db.lock.RLock()
	return d.packSizes
}

func (d *InMemoryDB) set(pz []int) {
	defer db.lock.Unlock()
	db.lock.Lock()
	d.packSizes = pz
}
