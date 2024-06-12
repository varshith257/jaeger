// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	es "github.com/jaegertracing/jaeger/pkg/es"
	mock "github.com/stretchr/testify/mock"
)

// IndexService is an autogenerated mock type for the IndexService type
type IndexService struct {
	mock.Mock
}

// Add provides a mock function with given fields:
func (_m *IndexService) Add() {
	_m.Called()
}

// BodyJson provides a mock function with given fields: body
func (_m *IndexService) BodyJson(body interface{}) es.IndexService {
	ret := _m.Called(body)

	if len(ret) == 0 {
		panic("no return value specified for BodyJson")
	}

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(interface{}) es.IndexService); ok {
		r0 = rf(body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Id provides a mock function with given fields: id
func (_m *IndexService) Id(id string) es.IndexService {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Id")
	}

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Index provides a mock function with given fields: index
func (_m *IndexService) Index(index string) es.IndexService {
	ret := _m.Called(index)

	if len(ret) == 0 {
		panic("no return value specified for Index")
	}

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(index)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// Type provides a mock function with given fields: typ
func (_m *IndexService) Type(typ string) es.IndexService {
	ret := _m.Called(typ)

	if len(ret) == 0 {
		panic("no return value specified for Type")
	}

	var r0 es.IndexService
	if rf, ok := ret.Get(0).(func(string) es.IndexService); ok {
		r0 = rf(typ)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(es.IndexService)
		}
	}

	return r0
}

// NewIndexService creates a new instance of IndexService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIndexService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IndexService {
	mock := &IndexService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
