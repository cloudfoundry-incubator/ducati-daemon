// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/models"
	"github.com/cloudfoundry-incubator/ducati-daemon/network"
)

type NetworkMapper struct {
	GetVNIStub        func(networkID string) (int, error)
	getVNIMutex       sync.RWMutex
	getVNIArgsForCall []struct {
		networkID string
	}
	getVNIReturns struct {
		result1 int
		result2 error
	}
	GetNetworkIDStub        func(netPayload models.NetworkPayload) (string, error)
	getNetworkIDMutex       sync.RWMutex
	getNetworkIDArgsForCall []struct {
		netPayload models.NetworkPayload
	}
	getNetworkIDReturns struct {
		result1 string
		result2 error
	}
}

func (fake *NetworkMapper) GetVNI(networkID string) (int, error) {
	fake.getVNIMutex.Lock()
	fake.getVNIArgsForCall = append(fake.getVNIArgsForCall, struct {
		networkID string
	}{networkID})
	fake.getVNIMutex.Unlock()
	if fake.GetVNIStub != nil {
		return fake.GetVNIStub(networkID)
	} else {
		return fake.getVNIReturns.result1, fake.getVNIReturns.result2
	}
}

func (fake *NetworkMapper) GetVNICallCount() int {
	fake.getVNIMutex.RLock()
	defer fake.getVNIMutex.RUnlock()
	return len(fake.getVNIArgsForCall)
}

func (fake *NetworkMapper) GetVNIArgsForCall(i int) string {
	fake.getVNIMutex.RLock()
	defer fake.getVNIMutex.RUnlock()
	return fake.getVNIArgsForCall[i].networkID
}

func (fake *NetworkMapper) GetVNIReturns(result1 int, result2 error) {
	fake.GetVNIStub = nil
	fake.getVNIReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *NetworkMapper) GetNetworkID(netPayload models.NetworkPayload) (string, error) {
	fake.getNetworkIDMutex.Lock()
	fake.getNetworkIDArgsForCall = append(fake.getNetworkIDArgsForCall, struct {
		netPayload models.NetworkPayload
	}{netPayload})
	fake.getNetworkIDMutex.Unlock()
	if fake.GetNetworkIDStub != nil {
		return fake.GetNetworkIDStub(netPayload)
	} else {
		return fake.getNetworkIDReturns.result1, fake.getNetworkIDReturns.result2
	}
}

func (fake *NetworkMapper) GetNetworkIDCallCount() int {
	fake.getNetworkIDMutex.RLock()
	defer fake.getNetworkIDMutex.RUnlock()
	return len(fake.getNetworkIDArgsForCall)
}

func (fake *NetworkMapper) GetNetworkIDArgsForCall(i int) models.NetworkPayload {
	fake.getNetworkIDMutex.RLock()
	defer fake.getNetworkIDMutex.RUnlock()
	return fake.getNetworkIDArgsForCall[i].netPayload
}

func (fake *NetworkMapper) GetNetworkIDReturns(result1 string, result2 error) {
	fake.GetNetworkIDStub = nil
	fake.getNetworkIDReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

var _ network.NetworkMapper = new(NetworkMapper)
