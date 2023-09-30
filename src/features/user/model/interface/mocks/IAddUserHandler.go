// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// IAddUserHandler is an autogenerated mock type for the IAddUserHandler type
type IAddUserHandler struct {
	mock.Mock
}

// Add provides a mock function with given fields: c
func (_m *IAddUserHandler) Add(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIAddUserHandler creates a new instance of IAddUserHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAddUserHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAddUserHandler {
	mock := &IAddUserHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
