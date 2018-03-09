// This file was generated by counterfeiter
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeServiceBrokerRepository struct {
	ListServiceBrokersStub        func(callback func(models.ServiceBroker) bool) error
	listServiceBrokersMutex       sync.RWMutex
	listServiceBrokersArgsForCall []struct {
		callback func(models.ServiceBroker) bool
	}
	listServiceBrokersReturns struct {
		result1 error
	}
	FindByNameStub        func(name string) (serviceBroker models.ServiceBroker, apiErr error)
	findByNameMutex       sync.RWMutex
	findByNameArgsForCall []struct {
		name string
	}
	findByNameReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	FindByGUIDStub        func(guid string) (serviceBroker models.ServiceBroker, apiErr error)
	findByGUIDMutex       sync.RWMutex
	findByGUIDArgsForCall []struct {
		guid string
	}
	findByGUIDReturns struct {
		result1 models.ServiceBroker
		result2 error
	}
	CreateStub        func(name, url, username, password, spaceGUID string) (apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		name      string
		url       string
		username  string
		password  string
		spaceGUID string
	}
	createReturns struct {
		result1 error
	}
	UpdateStub        func(serviceBroker models.ServiceBroker) (apiErr error)
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		serviceBroker models.ServiceBroker
	}
	updateReturns struct {
		result1 error
	}
	RenameStub        func(guid, name string) (apiErr error)
	renameMutex       sync.RWMutex
	renameArgsForCall []struct {
		guid string
		name string
	}
	renameReturns struct {
		result1 error
	}
	DeleteStub        func(guid string) (apiErr error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		guid string
	}
	deleteReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokers(callback func(models.ServiceBroker) bool) error {
	fake.listServiceBrokersMutex.Lock()
	fake.listServiceBrokersArgsForCall = append(fake.listServiceBrokersArgsForCall, struct {
		callback func(models.ServiceBroker) bool
	}{callback})
	fake.recordInvocation("ListServiceBrokers", []interface{}{callback})
	fake.listServiceBrokersMutex.Unlock()
	if fake.ListServiceBrokersStub != nil {
		return fake.ListServiceBrokersStub(callback)
	} else {
		return fake.listServiceBrokersReturns.result1
	}
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersCallCount() int {
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	return len(fake.listServiceBrokersArgsForCall)
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersArgsForCall(i int) func(models.ServiceBroker) bool {
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	return fake.listServiceBrokersArgsForCall[i].callback
}

func (fake *FakeServiceBrokerRepository) ListServiceBrokersReturns(result1 error) {
	fake.ListServiceBrokersStub = nil
	fake.listServiceBrokersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) FindByName(name string) (serviceBroker models.ServiceBroker, apiErr error) {
	fake.findByNameMutex.Lock()
	fake.findByNameArgsForCall = append(fake.findByNameArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("FindByName", []interface{}{name})
	fake.findByNameMutex.Unlock()
	if fake.FindByNameStub != nil {
		return fake.FindByNameStub(name)
	} else {
		return fake.findByNameReturns.result1, fake.findByNameReturns.result2
	}
}

func (fake *FakeServiceBrokerRepository) FindByNameCallCount() int {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return len(fake.findByNameArgsForCall)
}

func (fake *FakeServiceBrokerRepository) FindByNameArgsForCall(i int) string {
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	return fake.findByNameArgsForCall[i].name
}

func (fake *FakeServiceBrokerRepository) FindByNameReturns(result1 models.ServiceBroker, result2 error) {
	fake.FindByNameStub = nil
	fake.findByNameReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerRepository) FindByGUID(guid string) (serviceBroker models.ServiceBroker, apiErr error) {
	fake.findByGUIDMutex.Lock()
	fake.findByGUIDArgsForCall = append(fake.findByGUIDArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("FindByGUID", []interface{}{guid})
	fake.findByGUIDMutex.Unlock()
	if fake.FindByGUIDStub != nil {
		return fake.FindByGUIDStub(guid)
	} else {
		return fake.findByGUIDReturns.result1, fake.findByGUIDReturns.result2
	}
}

func (fake *FakeServiceBrokerRepository) FindByGUIDCallCount() int {
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	return len(fake.findByGUIDArgsForCall)
}

func (fake *FakeServiceBrokerRepository) FindByGUIDArgsForCall(i int) string {
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	return fake.findByGUIDArgsForCall[i].guid
}

func (fake *FakeServiceBrokerRepository) FindByGUIDReturns(result1 models.ServiceBroker, result2 error) {
	fake.FindByGUIDStub = nil
	fake.findByGUIDReturns = struct {
		result1 models.ServiceBroker
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceBrokerRepository) Create(name string, url string, username string, password string, spaceGUID string) (apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		name      string
		url       string
		username  string
		password  string
		spaceGUID string
	}{name, url, username, password, spaceGUID})
	fake.recordInvocation("Create", []interface{}{name, url, username, password, spaceGUID})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(name, url, username, password, spaceGUID)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeServiceBrokerRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeServiceBrokerRepository) CreateArgsForCall(i int) (string, string, string, string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].name, fake.createArgsForCall[i].url, fake.createArgsForCall[i].username, fake.createArgsForCall[i].password, fake.createArgsForCall[i].spaceGUID
}

func (fake *FakeServiceBrokerRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Update(serviceBroker models.ServiceBroker) (apiErr error) {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		serviceBroker models.ServiceBroker
	}{serviceBroker})
	fake.recordInvocation("Update", []interface{}{serviceBroker})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(serviceBroker)
	} else {
		return fake.updateReturns.result1
	}
}

func (fake *FakeServiceBrokerRepository) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeServiceBrokerRepository) UpdateArgsForCall(i int) models.ServiceBroker {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].serviceBroker
}

func (fake *FakeServiceBrokerRepository) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Rename(guid string, name string) (apiErr error) {
	fake.renameMutex.Lock()
	fake.renameArgsForCall = append(fake.renameArgsForCall, struct {
		guid string
		name string
	}{guid, name})
	fake.recordInvocation("Rename", []interface{}{guid, name})
	fake.renameMutex.Unlock()
	if fake.RenameStub != nil {
		return fake.RenameStub(guid, name)
	} else {
		return fake.renameReturns.result1
	}
}

func (fake *FakeServiceBrokerRepository) RenameCallCount() int {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return len(fake.renameArgsForCall)
}

func (fake *FakeServiceBrokerRepository) RenameArgsForCall(i int) (string, string) {
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	return fake.renameArgsForCall[i].guid, fake.renameArgsForCall[i].name
}

func (fake *FakeServiceBrokerRepository) RenameReturns(result1 error) {
	fake.RenameStub = nil
	fake.renameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Delete(guid string) (apiErr error) {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		guid string
	}{guid})
	fake.recordInvocation("Delete", []interface{}{guid})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(guid)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeServiceBrokerRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServiceBrokerRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].guid
}

func (fake *FakeServiceBrokerRepository) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServiceBrokerRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listServiceBrokersMutex.RLock()
	defer fake.listServiceBrokersMutex.RUnlock()
	fake.findByNameMutex.RLock()
	defer fake.findByNameMutex.RUnlock()
	fake.findByGUIDMutex.RLock()
	defer fake.findByGUIDMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	fake.renameMutex.RLock()
	defer fake.renameMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeServiceBrokerRepository) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ api.ServiceBrokerRepository = new(FakeServiceBrokerRepository)
