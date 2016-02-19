// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/commands"
	"github.com/cloudfoundry-incubator/ducati-daemon/executor"
)

type Executor struct {
	ExecuteStub        func(commands.Command) error
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		arg1 commands.Command
	}
	executeReturns struct {
		result1 error
	}
}

func (fake *Executor) Execute(arg1 commands.Command) error {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		arg1 commands.Command
	}{arg1})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(arg1)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *Executor) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *Executor) ExecuteArgsForCall(i int) commands.Command {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].arg1
}

func (fake *Executor) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}

var _ executor.Executor = new(Executor)
