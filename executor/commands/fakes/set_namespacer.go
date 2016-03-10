// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/executor/commands"
)

type SetNamespacer struct {
	SetNamespaceStub        func(intefaceName, namespace string) error
	setNamespaceMutex       sync.RWMutex
	setNamespaceArgsForCall []struct {
		intefaceName string
		namespace    string
	}
	setNamespaceReturns struct {
		result1 error
	}
}

func (fake *SetNamespacer) SetNamespace(intefaceName string, namespace string) error {
	fake.setNamespaceMutex.Lock()
	fake.setNamespaceArgsForCall = append(fake.setNamespaceArgsForCall, struct {
		intefaceName string
		namespace    string
	}{intefaceName, namespace})
	fake.setNamespaceMutex.Unlock()
	if fake.SetNamespaceStub != nil {
		return fake.SetNamespaceStub(intefaceName, namespace)
	} else {
		return fake.setNamespaceReturns.result1
	}
}

func (fake *SetNamespacer) SetNamespaceCallCount() int {
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	return len(fake.setNamespaceArgsForCall)
}

func (fake *SetNamespacer) SetNamespaceArgsForCall(i int) (string, string) {
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	return fake.setNamespaceArgsForCall[i].intefaceName, fake.setNamespaceArgsForCall[i].namespace
}

func (fake *SetNamespacer) SetNamespaceReturns(result1 error) {
	fake.SetNamespaceStub = nil
	fake.setNamespaceReturns = struct {
		result1 error
	}{result1}
}

var _ commands.SetNamespacer = new(SetNamespacer)