// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	histories "daily-tracker-calories/bussiness/histories"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// CreateHistories provides a mock function with given fields: _a0
func (_m *Service) CreateHistories(_a0 *histories.Domain) (*histories.Domain, error) {
	ret := _m.Called(_a0)

	var r0 *histories.Domain
	if rf, ok := ret.Get(0).(func(*histories.Domain) *histories.Domain); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*histories.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*histories.Domain) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllHistoriesByUserID provides a mock function with given fields: userid
func (_m *Service) GetAllHistoriesByUserID(userid int) (*[]histories.Domain, error) {
	ret := _m.Called(userid)

	var r0 *[]histories.Domain
	if rf, ok := ret.Get(0).(func(int) *[]histories.Domain); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]histories.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserStat provides a mock function with given fields: userid
func (_m *Service) UserStat(userid int) (float64, float64, string, string, error) {
	ret := _m.Called(userid)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int) float64); ok {
		r0 = rf(userid)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 float64
	if rf, ok := ret.Get(1).(func(int) float64); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Get(1).(float64)
	}

	var r2 string
	if rf, ok := ret.Get(2).(func(int) string); ok {
		r2 = rf(userid)
	} else {
		r2 = ret.Get(2).(string)
	}

	var r3 string
	if rf, ok := ret.Get(3).(func(int) string); ok {
		r3 = rf(userid)
	} else {
		r3 = ret.Get(3).(string)
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(int) error); ok {
		r4 = rf(userid)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}
