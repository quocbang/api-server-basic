// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo"
	mock "github.com/stretchr/testify/mock"

	requests "github.com/quocbang/api-server-basic/impl/requests"

	time "time"
)

// Users is an autogenerated mock type for the Users type
type Users struct {
	*mock.Mock
}

// CreateAccount provides a mock function with given fields: _a0, _a1
func (_m *Users) CreateAccount(_a0 echo.Context, _a1 requests.CreateAccountRequest) (requests.CreateAccountReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.CreateAccountReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.CreateAccountRequest) (requests.CreateAccountReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.CreateAccountRequest) requests.CreateAccountReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.CreateAccountReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.CreateAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAccounts provides a mock function with given fields: _a0, _a1
func (_m *Users) DeleteAccounts(_a0 echo.Context, _a1 requests.DeleteAccountRequest) (requests.DeleteAccountReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.DeleteAccountReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.DeleteAccountRequest) (requests.DeleteAccountReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.DeleteAccountRequest) requests.DeleteAccountReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.DeleteAccountReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.DeleteAccountRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: _a0, _a1
func (_m *Users) Login(_a0 echo.Context, _a1 requests.LoginRequest) (requests.LoginReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.LoginReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.LoginRequest) (requests.LoginReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.LoginRequest) requests.LoginReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.LoginReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.LoginRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignOut provides a mock function with given fields: _a0, _a1, _a2
func (_m *Users) SignOut(_a0 echo.Context, _a1 string, _a2 time.Duration) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context, string, time.Duration) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUsers creates a new instance of Users. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsers(t interface {
	mock.TestingT
	Cleanup(func())
}) *Users {
	mock := &Users{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
