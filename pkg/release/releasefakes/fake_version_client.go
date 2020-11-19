// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	"k8s.io/release/pkg/release"
)

type FakeVersionClient struct {
	GetURLResponseStub        func(string) (string, error)
	getURLResponseMutex       sync.RWMutex
	getURLResponseArgsForCall []struct {
		arg1 string
	}
	getURLResponseReturns struct {
		result1 string
		result2 error
	}
	getURLResponseReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVersionClient) GetURLResponse(arg1 string) (string, error) {
	fake.getURLResponseMutex.Lock()
	ret, specificReturn := fake.getURLResponseReturnsOnCall[len(fake.getURLResponseArgsForCall)]
	fake.getURLResponseArgsForCall = append(fake.getURLResponseArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetURLResponseStub
	fakeReturns := fake.getURLResponseReturns
	fake.recordInvocation("GetURLResponse", []interface{}{arg1})
	fake.getURLResponseMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVersionClient) GetURLResponseCallCount() int {
	fake.getURLResponseMutex.RLock()
	defer fake.getURLResponseMutex.RUnlock()
	return len(fake.getURLResponseArgsForCall)
}

func (fake *FakeVersionClient) GetURLResponseCalls(stub func(string) (string, error)) {
	fake.getURLResponseMutex.Lock()
	defer fake.getURLResponseMutex.Unlock()
	fake.GetURLResponseStub = stub
}

func (fake *FakeVersionClient) GetURLResponseArgsForCall(i int) string {
	fake.getURLResponseMutex.RLock()
	defer fake.getURLResponseMutex.RUnlock()
	argsForCall := fake.getURLResponseArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVersionClient) GetURLResponseReturns(result1 string, result2 error) {
	fake.getURLResponseMutex.Lock()
	defer fake.getURLResponseMutex.Unlock()
	fake.GetURLResponseStub = nil
	fake.getURLResponseReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionClient) GetURLResponseReturnsOnCall(i int, result1 string, result2 error) {
	fake.getURLResponseMutex.Lock()
	defer fake.getURLResponseMutex.Unlock()
	fake.GetURLResponseStub = nil
	if fake.getURLResponseReturnsOnCall == nil {
		fake.getURLResponseReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getURLResponseReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeVersionClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getURLResponseMutex.RLock()
	defer fake.getURLResponseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVersionClient) recordInvocation(key string, args []interface{}) {
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

var _ release.VersionClient = new(FakeVersionClient)
