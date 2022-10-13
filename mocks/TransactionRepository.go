// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "assignment-golang-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionRepository is an autogenerated mock type for the TransactionRepository type
type TransactionRepository struct {
	mock.Mock
}

// DoTransaction provides a mock function with given fields: e
func (_m *TransactionRepository) DoTransaction(e *entity.Transaction) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Transaction) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllTransactionById provides a mock function with given fields: walletid, params
func (_m *TransactionRepository) GetAllTransactionById(walletid int, params map[string]string) ([]*entity.Transaction, error) {
	ret := _m.Called(walletid, params)

	var r0 []*entity.Transaction
	if rf, ok := ret.Get(0).(func(int, map[string]string) []*entity.Transaction); ok {
		r0 = rf(walletid, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, map[string]string) error); ok {
		r1 = rf(walletid, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFundNameById provides a mock function with given fields: id
func (_m *TransactionRepository) GetFundNameById(id int) (*entity.Fund, error) {
	ret := _m.Called(id)

	var r0 *entity.Fund
	if rf, ok := ret.Get(0).(func(int) *entity.Fund); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Fund)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionRepository creates a new instance of TransactionRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionRepository(t mockConstructorTestingTNewTransactionRepository) *TransactionRepository {
	mock := &TransactionRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
