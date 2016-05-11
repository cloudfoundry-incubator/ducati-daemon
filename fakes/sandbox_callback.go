// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"
	"github.com/cloudfoundry-incubator/ducati-daemon/sandbox"
)

type SandboxCallback struct {
	CallbackStub        func(ns namespace.Namespace) error
	callbackMutex       sync.RWMutex
	callbackArgsForCall []struct {
		ns namespace.Namespace
	}
	callbackReturns struct {
		result1 error
	}
}

func (fake *SandboxCallback) Callback(ns namespace.Namespace) error {
	fake.callbackMutex.Lock()
	fake.callbackArgsForCall = append(fake.callbackArgsForCall, struct {
		ns namespace.Namespace
	}{ns})
	fake.callbackMutex.Unlock()
	if fake.CallbackStub != nil {
		return fake.CallbackStub(ns)
	} else {
		return fake.callbackReturns.result1
	}
}

func (fake *SandboxCallback) CallbackCallCount() int {
	fake.callbackMutex.RLock()
	defer fake.callbackMutex.RUnlock()
	return len(fake.callbackArgsForCall)
}

func (fake *SandboxCallback) CallbackArgsForCall(i int) namespace.Namespace {
	fake.callbackMutex.RLock()
	defer fake.callbackMutex.RUnlock()
	return fake.callbackArgsForCall[i].ns
}

func (fake *SandboxCallback) CallbackReturns(result1 error) {
	fake.CallbackStub = nil
	fake.callbackReturns = struct {
		result1 error
	}{result1}
}

var _ sandbox.SandboxCallback = new(SandboxCallback)
