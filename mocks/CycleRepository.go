// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	dto "financial/src/dto"
	model "financial/src/model"

	mock "github.com/stretchr/testify/mock"
)

// CycleRepository is an autogenerated mock type for the CycleRepository type
type CycleRepository struct {
	mock.Mock
}

// CreateCycle provides a mock function with given fields: cycle
func (_m *CycleRepository) CreateCycle(cycle dto.CreateCycleRequest) (model.Cycle, error) {
	ret := _m.Called(cycle)

	if len(ret) == 0 {
		panic("no return value specified for CreateCycle")
	}

	var r0 model.Cycle
	var r1 error
	if rf, ok := ret.Get(0).(func(dto.CreateCycleRequest) (model.Cycle, error)); ok {
		return rf(cycle)
	}
	if rf, ok := ret.Get(0).(func(dto.CreateCycleRequest) model.Cycle); ok {
		r0 = rf(cycle)
	} else {
		r0 = ret.Get(0).(model.Cycle)
	}

	if rf, ok := ret.Get(1).(func(dto.CreateCycleRequest) error); ok {
		r1 = rf(cycle)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewCycleRepository creates a new instance of CycleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCycleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *CycleRepository {
	mock := &CycleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
