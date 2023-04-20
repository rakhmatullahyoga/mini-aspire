// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	auth "github.com/rakhmatullahyoga/mini-aspire/auth"
	mock "github.com/stretchr/testify/mock"
)

// AuthUsecase is an autogenerated mock type for the AuthUsecase type
type AuthUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: username, password
func (_m *AuthUsecase) Login(username string, password string) (auth.Token, error) {
	ret := _m.Called(username, password)

	var r0 auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (auth.Token, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) auth.Token); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(auth.Token)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: username, password
func (_m *AuthUsecase) Register(username string, password string) (auth.Token, error) {
	ret := _m.Called(username, password)

	var r0 auth.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (auth.Token, error)); ok {
		return rf(username, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) auth.Token); ok {
		r0 = rf(username, password)
	} else {
		r0 = ret.Get(0).(auth.Token)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(username, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthUsecase creates a new instance of AuthUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthUsecase(t mockConstructorTestingTNewAuthUsecase) *AuthUsecase {
	mock := &AuthUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}