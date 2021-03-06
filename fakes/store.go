// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/models"
	"github.com/cloudfoundry-incubator/ducati-daemon/store"
)

type Store struct {
	CreateStub        func(container models.Container) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		container models.Container
	}
	createReturns struct {
		result1 error
	}
	GetStub        func(id string) (models.Container, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		id string
	}
	getReturns struct {
		result1 models.Container
		result2 error
	}
	AllStub        func() ([]models.Container, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []models.Container
		result2 error
	}
	DeleteStub        func(id string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		id string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *Store) Create(container models.Container) error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		container models.Container
	}{container})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(container)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *Store) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *Store) CreateArgsForCall(i int) models.Container {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].container
}

func (fake *Store) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *Store) Get(id string) (models.Container, error) {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		id string
	}{id})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(id)
	} else {
		return fake.getReturns.result1, fake.getReturns.result2
	}
}

func (fake *Store) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *Store) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].id
}

func (fake *Store) GetReturns(result1 models.Container, result2 error) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 models.Container
		result2 error
	}{result1, result2}
}

func (fake *Store) All() ([]models.Container, error) {
	fake.allMutex.Lock()
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	} else {
		return fake.allReturns.result1, fake.allReturns.result2
	}
}

func (fake *Store) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *Store) AllReturns(result1 []models.Container, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []models.Container
		result2 error
	}{result1, result2}
}

func (fake *Store) Delete(id string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		id string
	}{id})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(id)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *Store) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *Store) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].id
}

func (fake *Store) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ store.Store = new(Store)
