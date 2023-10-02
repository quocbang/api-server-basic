// Code generated by mockery v2.32.0. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo"
	mock "github.com/stretchr/testify/mock"

	requests "github.com/quocbang/api-server-basic/impl/requests"
)

// Tasks is an autogenerated mock type for the Tasks type
type Tasks struct {
	*mock.Mock
}

// CreateTasks provides a mock function with given fields: _a0, _a1
func (_m *Tasks) CreateTasks(_a0 echo.Context, _a1 requests.CreateTaskRequest) (requests.CreateTaskReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.CreateTaskReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.CreateTaskRequest) (requests.CreateTaskReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.CreateTaskRequest) requests.CreateTaskReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.CreateTaskReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.CreateTaskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTasks provides a mock function with given fields: _a0, _a1
func (_m *Tasks) DeleteTasks(_a0 echo.Context, _a1 requests.DeleteTaskRequest) (requests.DeleteTaskReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.DeleteTaskReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.DeleteTaskRequest) (requests.DeleteTaskReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.DeleteTaskRequest) requests.DeleteTaskReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.DeleteTaskReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.DeleteTaskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTasks provides a mock function with given fields: _a0, _a1
func (_m *Tasks) GetTasks(_a0 echo.Context, _a1 requests.GetTaskRequest) (requests.GetTaskReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.GetTaskReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.GetTaskRequest) (requests.GetTaskReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.GetTaskRequest) requests.GetTaskReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.GetTaskReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.GetTaskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: _a0, _a1
func (_m *Tasks) UpdateTask(_a0 echo.Context, _a1 requests.UpdateTaskRequest) (requests.UpdateTaskReply, error) {
	ret := _m.Called(_a0, _a1)

	var r0 requests.UpdateTaskReply
	var r1 error
	if rf, ok := ret.Get(0).(func(echo.Context, requests.UpdateTaskRequest) (requests.UpdateTaskReply, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(echo.Context, requests.UpdateTaskRequest) requests.UpdateTaskReply); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(requests.UpdateTaskReply)
	}

	if rf, ok := ret.Get(1).(func(echo.Context, requests.UpdateTaskRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewTasks creates a new instance of Tasks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTasks(t interface {
	mock.TestingT
	Cleanup(func())
}) *Tasks {
	mock := &Tasks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
