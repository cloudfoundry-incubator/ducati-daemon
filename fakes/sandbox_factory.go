// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/lib/namespace"
	"github.com/cloudfoundry-incubator/ducati-daemon/sandbox"
	"github.com/cloudfoundry-incubator/ducati-daemon/watcher"
	"github.com/pivotal-golang/lager"
)

type SandboxFactory struct {
	NewStub        func(lager.Logger, namespace.Namespace, sandbox.Invoker, sandbox.LinkFactory, watcher.MissWatcher) sandbox.Sandbox
	newMutex       sync.RWMutex
	newArgsForCall []struct {
		arg1 lager.Logger
		arg2 namespace.Namespace
		arg3 sandbox.Invoker
		arg4 sandbox.LinkFactory
		arg5 watcher.MissWatcher
	}
	newReturns struct {
		result1 sandbox.Sandbox
	}
}

func (fake *SandboxFactory) New(arg1 lager.Logger, arg2 namespace.Namespace, arg3 sandbox.Invoker, arg4 sandbox.LinkFactory, arg5 watcher.MissWatcher) sandbox.Sandbox {
	fake.newMutex.Lock()
	fake.newArgsForCall = append(fake.newArgsForCall, struct {
		arg1 lager.Logger
		arg2 namespace.Namespace
		arg3 sandbox.Invoker
		arg4 sandbox.LinkFactory
		arg5 watcher.MissWatcher
	}{arg1, arg2, arg3, arg4, arg5})
	fake.newMutex.Unlock()
	if fake.NewStub != nil {
		return fake.NewStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.newReturns.result1
	}
}

func (fake *SandboxFactory) NewCallCount() int {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return len(fake.newArgsForCall)
}

func (fake *SandboxFactory) NewArgsForCall(i int) (lager.Logger, namespace.Namespace, sandbox.Invoker, sandbox.LinkFactory, watcher.MissWatcher) {
	fake.newMutex.RLock()
	defer fake.newMutex.RUnlock()
	return fake.newArgsForCall[i].arg1, fake.newArgsForCall[i].arg2, fake.newArgsForCall[i].arg3, fake.newArgsForCall[i].arg4, fake.newArgsForCall[i].arg5
}

func (fake *SandboxFactory) NewReturns(result1 sandbox.Sandbox) {
	fake.NewStub = nil
	fake.newReturns = struct {
		result1 sandbox.Sandbox
	}{result1}
}