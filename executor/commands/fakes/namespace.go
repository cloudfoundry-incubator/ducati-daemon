// This file was generated by counterfeiter
package fakes

import (
	"os"
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/executor/commands"
)

type Namespace struct {
	ExecuteStub        func(func(*os.File) error) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 func(*os.File) error
	}
	executeReturns struct {
		result1 error
	}
	PathStub        func() string
	pathMutex       sync.RWMutex
	pathArgsForCall []struct{}
	pathReturns     struct {
		result1 string
	}
}

func (fake *Namespace) Execute(arg1 func(*os.File) error) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 func(*os.File) error
	}{arg1})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *Namespace) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *Namespace) ExecuteArgsForCall(i int) func(*os.File) error {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].arg1
}

func (fake *Namespace) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

func (fake *Namespace) Path() string {
	fake.pathMutex.Lock()
	fake.pathArgsForCall = append(fake.pathArgsForCall, struct{}{})
	fake.pathMutex.Unlock()
	if fake.PathStub != nil {
		return fake.PathStub()
	} else {
		return fake.pathReturns.result1
	}
}

func (fake *Namespace) PathCallCount() int {
	fake.pathMutex.RLock()
	defer fake.pathMutex.RUnlock()
	return len(fake.pathArgsForCall)
}

func (fake *Namespace) PathReturns(result1 string) {
	fake.PathStub = nil
	fake.pathReturns = struct {
		result1 string
	}{result1}
}

var _ commands.Namespace = new(Namespace)