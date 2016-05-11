// This file was generated by counterfeiter
package fakes

import "sync"

type Writer struct {
	WriteStub        func(p []byte) (n int, err error)
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		p []byte
	}
	writeReturns struct {
		result1 int
		result2 error
	}
}

func (fake *Writer) Write(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.writeMutex.Lock()
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(p)
	} else {
		return fake.writeReturns.result1, fake.writeReturns.result2
	}
}

func (fake *Writer) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *Writer) WriteArgsForCall(i int) []byte {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return fake.writeArgsForCall[i].p
}

func (fake *Writer) WriteReturns(result1 int, result2 error) {
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}
