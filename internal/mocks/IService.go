// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// Hello provides a mock function with given fields:
func (_m *IService) Hello() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewIService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIService(t mockConstructorTestingTNewIService) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
