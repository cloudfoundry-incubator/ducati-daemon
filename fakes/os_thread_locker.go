// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/ossupport"
)

type OSThreadLocker struct {
	LockOSThreadStub          func()
	lockOSThreadMutex         sync.RWMutex
	lockOSThreadArgsForCall   []struct{}
	UnlockOSThreadStub        func()
	unlockOSThreadMutex       sync.RWMutex
	unlockOSThreadArgsForCall []struct{}
}

func (fake *OSThreadLocker) LockOSThread() {
	fake.lockOSThreadMutex.Lock()
	fake.lockOSThreadArgsForCall = append(fake.lockOSThreadArgsForCall, struct{}{})
	fake.lockOSThreadMutex.Unlock()
	if fake.LockOSThreadStub != nil {
		fake.LockOSThreadStub()
	}
}

func (fake *OSThreadLocker) LockOSThreadCallCount() int {
	fake.lockOSThreadMutex.RLock()
	defer fake.lockOSThreadMutex.RUnlock()
	return len(fake.lockOSThreadArgsForCall)
}

func (fake *OSThreadLocker) UnlockOSThread() {
	fake.unlockOSThreadMutex.Lock()
	fake.unlockOSThreadArgsForCall = append(fake.unlockOSThreadArgsForCall, struct{}{})
	fake.unlockOSThreadMutex.Unlock()
	if fake.UnlockOSThreadStub != nil {
		fake.UnlockOSThreadStub()
	}
}

func (fake *OSThreadLocker) UnlockOSThreadCallCount() int {
	fake.unlockOSThreadMutex.RLock()
	defer fake.unlockOSThreadMutex.RUnlock()
	return len(fake.unlockOSThreadArgsForCall)
}

var _ ossupport.OSThreadLocker = new(OSThreadLocker)
