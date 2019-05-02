// Code generated by counterfeiter. DO NOT EDIT.
package serviceutilfakes

import (
	"sync"

	"github.com/pivotal-cf/spring-cloud-services-cli-plugin/serviceutil"
)

type FakeServiceInstanceUrlResolver struct {
	GetManagementUrlStub        func(string, string, bool) (string, error)
	getManagementUrlMutex       sync.RWMutex
	getManagementUrlArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 bool
	}
	getManagementUrlReturns struct {
		result1 string
		result2 error
	}
	getManagementUrlReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetServiceInstanceUrlStub        func(string, string) (string, error)
	getServiceInstanceUrlMutex       sync.RWMutex
	getServiceInstanceUrlArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getServiceInstanceUrlReturns struct {
		result1 string
		result2 error
	}
	getServiceInstanceUrlReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServiceInstanceUrlResolver) GetManagementUrl(arg1 string, arg2 string, arg3 bool) (string, error) {
	fake.getManagementUrlMutex.Lock()
	ret, specificReturn := fake.getManagementUrlReturnsOnCall[len(fake.getManagementUrlArgsForCall)]
	fake.getManagementUrlArgsForCall = append(fake.getManagementUrlArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 bool
	}{arg1, arg2, arg3})
	fake.recordInvocation("GetManagementUrl", []interface{}{arg1, arg2, arg3})
	fake.getManagementUrlMutex.Unlock()
	if fake.GetManagementUrlStub != nil {
		return fake.GetManagementUrlStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getManagementUrlReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInstanceUrlResolver) GetManagementUrlCallCount() int {
	fake.getManagementUrlMutex.RLock()
	defer fake.getManagementUrlMutex.RUnlock()
	return len(fake.getManagementUrlArgsForCall)
}

func (fake *FakeServiceInstanceUrlResolver) GetManagementUrlCalls(stub func(string, string, bool) (string, error)) {
	fake.getManagementUrlMutex.Lock()
	defer fake.getManagementUrlMutex.Unlock()
	fake.GetManagementUrlStub = stub
}

func (fake *FakeServiceInstanceUrlResolver) GetManagementUrlArgsForCall(i int) (string, string, bool) {
	fake.getManagementUrlMutex.RLock()
	defer fake.getManagementUrlMutex.RUnlock()
	argsForCall := fake.getManagementUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeServiceInstanceUrlResolver) GetManagementUrlReturns(result1 string, result2 error) {
	fake.getManagementUrlMutex.Lock()
	defer fake.getManagementUrlMutex.Unlock()
	fake.GetManagementUrlStub = nil
	fake.getManagementUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInstanceUrlResolver) GetManagementUrlReturnsOnCall(i int, result1 string, result2 error) {
	fake.getManagementUrlMutex.Lock()
	defer fake.getManagementUrlMutex.Unlock()
	fake.GetManagementUrlStub = nil
	if fake.getManagementUrlReturnsOnCall == nil {
		fake.getManagementUrlReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getManagementUrlReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInstanceUrlResolver) GetServiceInstanceUrl(arg1 string, arg2 string) (string, error) {
	fake.getServiceInstanceUrlMutex.Lock()
	ret, specificReturn := fake.getServiceInstanceUrlReturnsOnCall[len(fake.getServiceInstanceUrlArgsForCall)]
	fake.getServiceInstanceUrlArgsForCall = append(fake.getServiceInstanceUrlArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetServiceInstanceUrl", []interface{}{arg1, arg2})
	fake.getServiceInstanceUrlMutex.Unlock()
	if fake.GetServiceInstanceUrlStub != nil {
		return fake.GetServiceInstanceUrlStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getServiceInstanceUrlReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServiceInstanceUrlResolver) GetServiceInstanceUrlCallCount() int {
	fake.getServiceInstanceUrlMutex.RLock()
	defer fake.getServiceInstanceUrlMutex.RUnlock()
	return len(fake.getServiceInstanceUrlArgsForCall)
}

func (fake *FakeServiceInstanceUrlResolver) GetServiceInstanceUrlCalls(stub func(string, string) (string, error)) {
	fake.getServiceInstanceUrlMutex.Lock()
	defer fake.getServiceInstanceUrlMutex.Unlock()
	fake.GetServiceInstanceUrlStub = stub
}

func (fake *FakeServiceInstanceUrlResolver) GetServiceInstanceUrlArgsForCall(i int) (string, string) {
	fake.getServiceInstanceUrlMutex.RLock()
	defer fake.getServiceInstanceUrlMutex.RUnlock()
	argsForCall := fake.getServiceInstanceUrlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServiceInstanceUrlResolver) GetServiceInstanceUrlReturns(result1 string, result2 error) {
	fake.getServiceInstanceUrlMutex.Lock()
	defer fake.getServiceInstanceUrlMutex.Unlock()
	fake.GetServiceInstanceUrlStub = nil
	fake.getServiceInstanceUrlReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInstanceUrlResolver) GetServiceInstanceUrlReturnsOnCall(i int, result1 string, result2 error) {
	fake.getServiceInstanceUrlMutex.Lock()
	defer fake.getServiceInstanceUrlMutex.Unlock()
	fake.GetServiceInstanceUrlStub = nil
	if fake.getServiceInstanceUrlReturnsOnCall == nil {
		fake.getServiceInstanceUrlReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getServiceInstanceUrlReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeServiceInstanceUrlResolver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getManagementUrlMutex.RLock()
	defer fake.getManagementUrlMutex.RUnlock()
	fake.getServiceInstanceUrlMutex.RLock()
	defer fake.getServiceInstanceUrlMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServiceInstanceUrlResolver) recordInvocation(key string, args []interface{}) {
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

var _ serviceutil.ServiceInstanceUrlResolver = new(FakeServiceInstanceUrlResolver)