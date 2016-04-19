// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"
	"github.com/cloudfoundry-incubator/ducati-daemon/sandbox"
)

type Sandbox struct {
	LockStub             func()
	lockMutex            sync.RWMutex
	lockArgsForCall      []struct{}
	UnlockStub           func()
	unlockMutex          sync.RWMutex
	unlockArgsForCall    []struct{}
	NamespaceStub        func() namespace.Namespace
	namespaceMutex       sync.RWMutex
	namespaceArgsForCall []struct{}
	namespaceReturns     struct {
		result1 namespace.Namespace
	}
}

func (fake *Sandbox) Lock() {
	fake.lockMutex.Lock()
	fake.lockArgsForCall = append(fake.lockArgsForCall, struct{}{})
	fake.lockMutex.Unlock()
	if fake.LockStub != nil {
		fake.LockStub()
	}
}

func (fake *Sandbox) LockCallCount() int {
	fake.lockMutex.RLock()
	defer fake.lockMutex.RUnlock()
	return len(fake.lockArgsForCall)
}

func (fake *Sandbox) Unlock() {
	fake.unlockMutex.Lock()
	fake.unlockArgsForCall = append(fake.unlockArgsForCall, struct{}{})
	fake.unlockMutex.Unlock()
	if fake.UnlockStub != nil {
		fake.UnlockStub()
	}
}

func (fake *Sandbox) UnlockCallCount() int {
	fake.unlockMutex.RLock()
	defer fake.unlockMutex.RUnlock()
	return len(fake.unlockArgsForCall)
}

func (fake *Sandbox) Namespace() namespace.Namespace {
	fake.namespaceMutex.Lock()
	fake.namespaceArgsForCall = append(fake.namespaceArgsForCall, struct{}{})
	fake.namespaceMutex.Unlock()
	if fake.NamespaceStub != nil {
		return fake.NamespaceStub()
	} else {
		return fake.namespaceReturns.result1
	}
}

func (fake *Sandbox) NamespaceCallCount() int {
	fake.namespaceMutex.RLock()
	defer fake.namespaceMutex.RUnlock()
	return len(fake.namespaceArgsForCall)
}

func (fake *Sandbox) NamespaceReturns(result1 namespace.Namespace) {
	fake.NamespaceStub = nil
	fake.namespaceReturns = struct {
		result1 namespace.Namespace
	}{result1}
}

var _ sandbox.Sandbox = new(Sandbox)