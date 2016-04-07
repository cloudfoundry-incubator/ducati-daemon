// This file was generated by counterfeiter
package fakes

import (
	"net/http"
	"sync"
)

type HTTPClient struct {
	DoStub        func(req *http.Request) (resp *http.Response, err error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		req *http.Request
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
}

func (fake *HTTPClient) Do(req *http.Request) (resp *http.Response, err error) {
	fake.doMutex.Lock()
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		req *http.Request
	}{req})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(req)
	} else {
		return fake.doReturns.result1, fake.doReturns.result2
	}
}

func (fake *HTTPClient) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *HTTPClient) DoArgsForCall(i int) *http.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.doArgsForCall[i].req
}

func (fake *HTTPClient) DoReturns(result1 *http.Response, result2 error) {
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}
