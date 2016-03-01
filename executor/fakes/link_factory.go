// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/cloudfoundry-incubator/ducati-daemon/executor"
)

type LinkFactory struct {
	CreateBridgeStub        func(name string) error
	createBridgeMutex       sync.RWMutex
	createBridgeArgsForCall []struct {
		name string
	}
	createBridgeReturns struct {
		result1 error
	}
	HardwareAddressStub        func(linkName string) (net.HardwareAddr, error)
	hardwareAddressMutex       sync.RWMutex
	hardwareAddressArgsForCall []struct {
		linkName string
	}
	hardwareAddressReturns struct {
		result1 net.HardwareAddr
		result2 error
	}
	SetMasterStub        func(slave, master string) error
	setMasterMutex       sync.RWMutex
	setMasterArgsForCall []struct {
		slave  string
		master string
	}
	setMasterReturns struct {
		result1 error
	}
	SetNamespaceStub        func(intefaceName, namespace string) error
	setNamespaceMutex       sync.RWMutex
	setNamespaceArgsForCall []struct {
		intefaceName string
		namespace    string
	}
	setNamespaceReturns struct {
		result1 error
	}
	SetUpStub        func(name string) error
	setUpMutex       sync.RWMutex
	setUpArgsForCall []struct {
		name string
	}
	setUpReturns struct {
		result1 error
	}
	CreateVethStub        func(name, peerName string, mtu int) error
	createVethMutex       sync.RWMutex
	createVethArgsForCall []struct {
		name     string
		peerName string
		mtu      int
	}
	createVethReturns struct {
		result1 error
	}
	CreateVxlanStub        func(name string, vni int) error
	createVxlanMutex       sync.RWMutex
	createVxlanArgsForCall []struct {
		name string
		vni  int
	}
	createVxlanReturns struct {
		result1 error
	}
	DeleteLinkByNameStub        func(name string) error
	deleteLinkByNameMutex       sync.RWMutex
	deleteLinkByNameArgsForCall []struct {
		name string
	}
	deleteLinkByNameReturns struct {
		result1 error
	}
	VethDeviceCountStub        func() (int, error)
	vethDeviceCountMutex       sync.RWMutex
	vethDeviceCountArgsForCall []struct{}
	vethDeviceCountReturns     struct {
		result1 int
		result2 error
	}
}

func (fake *LinkFactory) CreateBridge(name string) error {
	fake.createBridgeMutex.Lock()
	fake.createBridgeArgsForCall = append(fake.createBridgeArgsForCall, struct {
		name string
	}{name})
	fake.createBridgeMutex.Unlock()
	if fake.CreateBridgeStub != nil {
		return fake.CreateBridgeStub(name)
	} else {
		return fake.createBridgeReturns.result1
	}
}

func (fake *LinkFactory) CreateBridgeCallCount() int {
	fake.createBridgeMutex.RLock()
	defer fake.createBridgeMutex.RUnlock()
	return len(fake.createBridgeArgsForCall)
}

func (fake *LinkFactory) CreateBridgeArgsForCall(i int) string {
	fake.createBridgeMutex.RLock()
	defer fake.createBridgeMutex.RUnlock()
	return fake.createBridgeArgsForCall[i].name
}

func (fake *LinkFactory) CreateBridgeReturns(result1 error) {
	fake.CreateBridgeStub = nil
	fake.createBridgeReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) HardwareAddress(linkName string) (net.HardwareAddr, error) {
	fake.hardwareAddressMutex.Lock()
	fake.hardwareAddressArgsForCall = append(fake.hardwareAddressArgsForCall, struct {
		linkName string
	}{linkName})
	fake.hardwareAddressMutex.Unlock()
	if fake.HardwareAddressStub != nil {
		return fake.HardwareAddressStub(linkName)
	} else {
		return fake.hardwareAddressReturns.result1, fake.hardwareAddressReturns.result2
	}
}

func (fake *LinkFactory) HardwareAddressCallCount() int {
	fake.hardwareAddressMutex.RLock()
	defer fake.hardwareAddressMutex.RUnlock()
	return len(fake.hardwareAddressArgsForCall)
}

func (fake *LinkFactory) HardwareAddressArgsForCall(i int) string {
	fake.hardwareAddressMutex.RLock()
	defer fake.hardwareAddressMutex.RUnlock()
	return fake.hardwareAddressArgsForCall[i].linkName
}

func (fake *LinkFactory) HardwareAddressReturns(result1 net.HardwareAddr, result2 error) {
	fake.HardwareAddressStub = nil
	fake.hardwareAddressReturns = struct {
		result1 net.HardwareAddr
		result2 error
	}{result1, result2}
}

func (fake *LinkFactory) SetMaster(slave string, master string) error {
	fake.setMasterMutex.Lock()
	fake.setMasterArgsForCall = append(fake.setMasterArgsForCall, struct {
		slave  string
		master string
	}{slave, master})
	fake.setMasterMutex.Unlock()
	if fake.SetMasterStub != nil {
		return fake.SetMasterStub(slave, master)
	} else {
		return fake.setMasterReturns.result1
	}
}

func (fake *LinkFactory) SetMasterCallCount() int {
	fake.setMasterMutex.RLock()
	defer fake.setMasterMutex.RUnlock()
	return len(fake.setMasterArgsForCall)
}

func (fake *LinkFactory) SetMasterArgsForCall(i int) (string, string) {
	fake.setMasterMutex.RLock()
	defer fake.setMasterMutex.RUnlock()
	return fake.setMasterArgsForCall[i].slave, fake.setMasterArgsForCall[i].master
}

func (fake *LinkFactory) SetMasterReturns(result1 error) {
	fake.SetMasterStub = nil
	fake.setMasterReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) SetNamespace(intefaceName string, namespace string) error {
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

func (fake *LinkFactory) SetNamespaceCallCount() int {
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	return len(fake.setNamespaceArgsForCall)
}

func (fake *LinkFactory) SetNamespaceArgsForCall(i int) (string, string) {
	fake.setNamespaceMutex.RLock()
	defer fake.setNamespaceMutex.RUnlock()
	return fake.setNamespaceArgsForCall[i].intefaceName, fake.setNamespaceArgsForCall[i].namespace
}

func (fake *LinkFactory) SetNamespaceReturns(result1 error) {
	fake.SetNamespaceStub = nil
	fake.setNamespaceReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) SetUp(name string) error {
	fake.setUpMutex.Lock()
	fake.setUpArgsForCall = append(fake.setUpArgsForCall, struct {
		name string
	}{name})
	fake.setUpMutex.Unlock()
	if fake.SetUpStub != nil {
		return fake.SetUpStub(name)
	} else {
		return fake.setUpReturns.result1
	}
}

func (fake *LinkFactory) SetUpCallCount() int {
	fake.setUpMutex.RLock()
	defer fake.setUpMutex.RUnlock()
	return len(fake.setUpArgsForCall)
}

func (fake *LinkFactory) SetUpArgsForCall(i int) string {
	fake.setUpMutex.RLock()
	defer fake.setUpMutex.RUnlock()
	return fake.setUpArgsForCall[i].name
}

func (fake *LinkFactory) SetUpReturns(result1 error) {
	fake.SetUpStub = nil
	fake.setUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) CreateVeth(name string, peerName string, mtu int) error {
	fake.createVethMutex.Lock()
	fake.createVethArgsForCall = append(fake.createVethArgsForCall, struct {
		name     string
		peerName string
		mtu      int
	}{name, peerName, mtu})
	fake.createVethMutex.Unlock()
	if fake.CreateVethStub != nil {
		return fake.CreateVethStub(name, peerName, mtu)
	} else {
		return fake.createVethReturns.result1
	}
}

func (fake *LinkFactory) CreateVethCallCount() int {
	fake.createVethMutex.RLock()
	defer fake.createVethMutex.RUnlock()
	return len(fake.createVethArgsForCall)
}

func (fake *LinkFactory) CreateVethArgsForCall(i int) (string, string, int) {
	fake.createVethMutex.RLock()
	defer fake.createVethMutex.RUnlock()
	return fake.createVethArgsForCall[i].name, fake.createVethArgsForCall[i].peerName, fake.createVethArgsForCall[i].mtu
}

func (fake *LinkFactory) CreateVethReturns(result1 error) {
	fake.CreateVethStub = nil
	fake.createVethReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) CreateVxlan(name string, vni int) error {
	fake.createVxlanMutex.Lock()
	fake.createVxlanArgsForCall = append(fake.createVxlanArgsForCall, struct {
		name string
		vni  int
	}{name, vni})
	fake.createVxlanMutex.Unlock()
	if fake.CreateVxlanStub != nil {
		return fake.CreateVxlanStub(name, vni)
	} else {
		return fake.createVxlanReturns.result1
	}
}

func (fake *LinkFactory) CreateVxlanCallCount() int {
	fake.createVxlanMutex.RLock()
	defer fake.createVxlanMutex.RUnlock()
	return len(fake.createVxlanArgsForCall)
}

func (fake *LinkFactory) CreateVxlanArgsForCall(i int) (string, int) {
	fake.createVxlanMutex.RLock()
	defer fake.createVxlanMutex.RUnlock()
	return fake.createVxlanArgsForCall[i].name, fake.createVxlanArgsForCall[i].vni
}

func (fake *LinkFactory) CreateVxlanReturns(result1 error) {
	fake.CreateVxlanStub = nil
	fake.createVxlanReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) DeleteLinkByName(name string) error {
	fake.deleteLinkByNameMutex.Lock()
	fake.deleteLinkByNameArgsForCall = append(fake.deleteLinkByNameArgsForCall, struct {
		name string
	}{name})
	fake.deleteLinkByNameMutex.Unlock()
	if fake.DeleteLinkByNameStub != nil {
		return fake.DeleteLinkByNameStub(name)
	} else {
		return fake.deleteLinkByNameReturns.result1
	}
}

func (fake *LinkFactory) DeleteLinkByNameCallCount() int {
	fake.deleteLinkByNameMutex.RLock()
	defer fake.deleteLinkByNameMutex.RUnlock()
	return len(fake.deleteLinkByNameArgsForCall)
}

func (fake *LinkFactory) DeleteLinkByNameArgsForCall(i int) string {
	fake.deleteLinkByNameMutex.RLock()
	defer fake.deleteLinkByNameMutex.RUnlock()
	return fake.deleteLinkByNameArgsForCall[i].name
}

func (fake *LinkFactory) DeleteLinkByNameReturns(result1 error) {
	fake.DeleteLinkByNameStub = nil
	fake.deleteLinkByNameReturns = struct {
		result1 error
	}{result1}
}

func (fake *LinkFactory) VethDeviceCount() (int, error) {
	fake.vethDeviceCountMutex.Lock()
	fake.vethDeviceCountArgsForCall = append(fake.vethDeviceCountArgsForCall, struct{}{})
	fake.vethDeviceCountMutex.Unlock()
	if fake.VethDeviceCountStub != nil {
		return fake.VethDeviceCountStub()
	} else {
		return fake.vethDeviceCountReturns.result1, fake.vethDeviceCountReturns.result2
	}
}

func (fake *LinkFactory) VethDeviceCountCallCount() int {
	fake.vethDeviceCountMutex.RLock()
	defer fake.vethDeviceCountMutex.RUnlock()
	return len(fake.vethDeviceCountArgsForCall)
}

func (fake *LinkFactory) VethDeviceCountReturns(result1 int, result2 error) {
	fake.VethDeviceCountStub = nil
	fake.vethDeviceCountReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

var _ executor.LinkFactory = new(LinkFactory)
