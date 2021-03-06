// This file was generated by counterfeiter
package commandsfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands"
	"github.com/pivotal-cf/pivnet-cli/rc"
)

type FakeRCHandler struct {
	SaveProfileStub        func(profileName string, apiToken string, host string) error
	saveProfileMutex       sync.RWMutex
	saveProfileArgsForCall []struct {
		profileName string
		apiToken    string
		host        string
	}
	saveProfileReturns struct {
		result1 error
	}
	ProfileForNameStub        func(profileName string) (*rc.PivnetProfile, error)
	profileForNameMutex       sync.RWMutex
	profileForNameArgsForCall []struct {
		profileName string
	}
	profileForNameReturns struct {
		result1 *rc.PivnetProfile
		result2 error
	}
	RemoveProfileWithNameStub        func(profileName string) error
	removeProfileWithNameMutex       sync.RWMutex
	removeProfileWithNameArgsForCall []struct {
		profileName string
	}
	removeProfileWithNameReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRCHandler) SaveProfile(profileName string, apiToken string, host string) error {
	fake.saveProfileMutex.Lock()
	fake.saveProfileArgsForCall = append(fake.saveProfileArgsForCall, struct {
		profileName string
		apiToken    string
		host        string
	}{profileName, apiToken, host})
	fake.recordInvocation("SaveProfile", []interface{}{profileName, apiToken, host})
	fake.saveProfileMutex.Unlock()
	if fake.SaveProfileStub != nil {
		return fake.SaveProfileStub(profileName, apiToken, host)
	} else {
		return fake.saveProfileReturns.result1
	}
}

func (fake *FakeRCHandler) SaveProfileCallCount() int {
	fake.saveProfileMutex.RLock()
	defer fake.saveProfileMutex.RUnlock()
	return len(fake.saveProfileArgsForCall)
}

func (fake *FakeRCHandler) SaveProfileArgsForCall(i int) (string, string, string) {
	fake.saveProfileMutex.RLock()
	defer fake.saveProfileMutex.RUnlock()
	return fake.saveProfileArgsForCall[i].profileName, fake.saveProfileArgsForCall[i].apiToken, fake.saveProfileArgsForCall[i].host
}

func (fake *FakeRCHandler) SaveProfileReturns(result1 error) {
	fake.SaveProfileStub = nil
	fake.saveProfileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRCHandler) ProfileForName(profileName string) (*rc.PivnetProfile, error) {
	fake.profileForNameMutex.Lock()
	fake.profileForNameArgsForCall = append(fake.profileForNameArgsForCall, struct {
		profileName string
	}{profileName})
	fake.recordInvocation("ProfileForName", []interface{}{profileName})
	fake.profileForNameMutex.Unlock()
	if fake.ProfileForNameStub != nil {
		return fake.ProfileForNameStub(profileName)
	} else {
		return fake.profileForNameReturns.result1, fake.profileForNameReturns.result2
	}
}

func (fake *FakeRCHandler) ProfileForNameCallCount() int {
	fake.profileForNameMutex.RLock()
	defer fake.profileForNameMutex.RUnlock()
	return len(fake.profileForNameArgsForCall)
}

func (fake *FakeRCHandler) ProfileForNameArgsForCall(i int) string {
	fake.profileForNameMutex.RLock()
	defer fake.profileForNameMutex.RUnlock()
	return fake.profileForNameArgsForCall[i].profileName
}

func (fake *FakeRCHandler) ProfileForNameReturns(result1 *rc.PivnetProfile, result2 error) {
	fake.ProfileForNameStub = nil
	fake.profileForNameReturns = struct {
		result1 *rc.PivnetProfile
		result2 error
	}{result1, result2}
}

func (fake *FakeRCHandler) RemoveProfileWithName(profileName string) error {
	fake.removeProfileWithNameMutex.Lock()
	fake.removeProfileWithNameArgsForCall = append(fake.removeProfileWithNameArgsForCall, struct {
		profileName string
	}{profileName})
	fake.recordInvocation("RemoveProfileWithName", []interface{}{profileName})
	fake.removeProfileWithNameMutex.Unlock()
	if fake.RemoveProfileWithNameStub != nil {
		return fake.RemoveProfileWithNameStub(profileName)
	} else {
		return fake.removeProfileWithNameReturns.result1
	}
}

func (fake *FakeRCHandler) RemoveProfileWithNameCallCount() int {
	fake.removeProfileWithNameMutex.RLock()
	defer fake.removeProfileWithNameMutex.RUnlock()
	return len(fake.removeProfileWithNameArgsForCall)
}

func (fake *FakeRCHandler) RemoveProfileWithNameArgsForCall(i int) string {
	fake.removeProfileWithNameMutex.RLock()
	defer fake.removeProfileWithNameMutex.RUnlock()
	return fake.removeProfileWithNameArgsForCall[i].profileName
}

func (fake *FakeRCHandler) RemoveProfileWithNameReturns(result1 error) {
	fake.RemoveProfileWithNameStub = nil
	fake.removeProfileWithNameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRCHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.saveProfileMutex.RLock()
	defer fake.saveProfileMutex.RUnlock()
	fake.profileForNameMutex.RLock()
	defer fake.profileForNameMutex.RUnlock()
	fake.removeProfileWithNameMutex.RLock()
	defer fake.removeProfileWithNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRCHandler) recordInvocation(key string, args []interface{}) {
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

var _ commands.RCHandler = new(FakeRCHandler)
