// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by mockery v2.15.0. DO NOT EDIT.

package awslib

import mock "github.com/stretchr/testify/mock"

// MockCrossRegionFetcher is an autogenerated mock type for the CrossRegionFetcher type
type MockCrossRegionFetcher[T interface{}] struct {
	mock.Mock
}

type MockCrossRegionFetcher_Expecter[T interface{}] struct {
	mock *mock.Mock
}

func (_m *MockCrossRegionFetcher[T]) EXPECT() *MockCrossRegionFetcher_Expecter[T] {
	return &MockCrossRegionFetcher_Expecter[T]{mock: &_m.Mock}
}

// GetMultiRegionsClientMap provides a mock function with given fields:
func (_m *MockCrossRegionFetcher[T]) GetMultiRegionsClientMap() map[string]T {
	ret := _m.Called()

	var r0 map[string]T
	if rf, ok := ret.Get(0).(func() map[string]T); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]T)
		}
	}

	return r0
}

// MockCrossRegionFetcher_GetMultiRegionsClientMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMultiRegionsClientMap'
type MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T interface{}] struct {
	*mock.Call
}

// GetMultiRegionsClientMap is a helper method to define mock.On call
func (_e *MockCrossRegionFetcher_Expecter[T]) GetMultiRegionsClientMap() *MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T] {
	return &MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T]{Call: _e.mock.On("GetMultiRegionsClientMap")}
}

func (_c *MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T]) Run(run func()) *MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T] {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T]) Return(_a0 map[string]T) *MockCrossRegionFetcher_GetMultiRegionsClientMap_Call[T] {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockCrossRegionFetcher interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCrossRegionFetcher creates a new instance of MockCrossRegionFetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCrossRegionFetcher[T interface{}](t mockConstructorTestingTNewMockCrossRegionFetcher) *MockCrossRegionFetcher[T] {
	mock := &MockCrossRegionFetcher[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}