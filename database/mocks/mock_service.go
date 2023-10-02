// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	database "github.com/quocbang/api-server-basic/database"
)

// Services is an autogenerated mock type for the Services type
type Services struct {
	*mock.Mock
}

// Tasks provides a mock function with given fields:
func (_m *Services) Tasks() database.Tasks {
	ret := _m.Called()

	var r0 database.Tasks
	if rf, ok := ret.Get(0).(func() database.Tasks); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Tasks)
		}
	}

	return r0
}

// Users provides a mock function with given fields:
func (_m *Services) Users() database.Users {
	ret := _m.Called()

	var r0 database.Users
	if rf, ok := ret.Get(0).(func() database.Users); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Users)
		}
	}

	return r0
}

// NewServices creates a new instance of Services. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewServices(t interface {
	mock.TestingT
	Cleanup(func())
}) *Services {
	mock := &Services{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
