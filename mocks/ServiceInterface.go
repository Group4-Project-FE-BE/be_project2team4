// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "be_project2team4/feature/posting/domain"

	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// ServiceInterface is an autogenerated mock type for the ServiceInterface type
type ServiceInterface struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ID
func (_m *ServiceInterface) Delete(ID string) (domain.Core, error) {
	ret := _m.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(string) domain.Core); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ID
func (_m *ServiceInterface) Get(ID string) (domain.Core, error) {
	ret := _m.Called(ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(string) domain.Core); ok {
		r0 = rf(ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *ServiceInterface) GetAll() ([]domain.Core, error) {
	ret := _m.Called()

	var r0 []domain.Core
	if rf, ok := ret.Get(0).(func() []domain.Core); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Core)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Insert provides a mock function with given fields: newData
func (_m *ServiceInterface) Insert(newData domain.Core) (domain.Core, error) {
	ret := _m.Called(newData)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core) domain.Core); ok {
		r0 = rf(newData)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthorized provides a mock function with given fields: c
func (_m *ServiceInterface) IsAuthorized(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: updatedData, ID
func (_m *ServiceInterface) Update(updatedData domain.Core, ID string) (domain.Core, error) {
	ret := _m.Called(updatedData, ID)

	var r0 domain.Core
	if rf, ok := ret.Get(0).(func(domain.Core, string) domain.Core); ok {
		r0 = rf(updatedData, ID)
	} else {
		r0 = ret.Get(0).(domain.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(domain.Core, string) error); ok {
		r1 = rf(updatedData, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewServiceInterface creates a new instance of ServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewServiceInterface(t mockConstructorTestingTNewServiceInterface) *ServiceInterface {
	mock := &ServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
