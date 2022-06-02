// Code generated by mockery v2.12.0. DO NOT EDIT.

package mocks

import (
	model "ilanver/internal/model"

	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	repository "ilanver/internal/repository"

	testing "testing"
)

// IAddressRepo is an autogenerated mock type for the IAddressRepo type
type IAddressRepo struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: id
func (_m *IAddressRepo) GetByID(id uint) (model.Adress, error) {
	ret := _m.Called(id)

	var r0 model.Adress
	if rf, ok := ret.Get(0).(func(uint) model.Adress); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Adress)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: address
func (_m *IAddressRepo) Save(address *model.Adress) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Adress) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: address
func (_m *IAddressRepo) Update(address *model.Adress) error {
	ret := _m.Called(address)

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Adress) error); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithTx provides a mock function with given fields: db
func (_m *IAddressRepo) WithTx(db *gorm.DB) repository.IAddressRepo {
	ret := _m.Called(db)

	var r0 repository.IAddressRepo
	if rf, ok := ret.Get(0).(func(*gorm.DB) repository.IAddressRepo); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.IAddressRepo)
		}
	}

	return r0
}

// NewIAddressRepo creates a new instance of IAddressRepo. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewIAddressRepo(t testing.TB) *IAddressRepo {
	mock := &IAddressRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}