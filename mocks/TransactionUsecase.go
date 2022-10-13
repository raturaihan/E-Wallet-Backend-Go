// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "assignment-golang-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// TransactionUsecase is an autogenerated mock type for the TransactionUsecase type
type TransactionUsecase struct {
	mock.Mock
}

// GenerateDescription provides a mock function with given fields: fundId
func (_m *TransactionUsecase) GenerateDescription(fundId int) string {
	ret := _m.Called(fundId)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(fundId)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetAllTransactionById provides a mock function with given fields: walletid, params
func (_m *TransactionUsecase) GetAllTransactionById(walletid int, params map[string]string) ([]*entity.Transaction, error) {
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

// TopUpAmount provides a mock function with given fields: e
func (_m *TransactionUsecase) TopUpAmount(e *entity.Transaction) (*entity.Transaction, error) {
	ret := _m.Called(e)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(*entity.Transaction) *entity.Transaction); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Transaction) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Transfer provides a mock function with given fields: e
func (_m *TransactionUsecase) Transfer(e *entity.Transaction) (*entity.Transaction, error) {
	ret := _m.Called(e)

	var r0 *entity.Transaction
	if rf, ok := ret.Get(0).(func(*entity.Transaction) *entity.Transaction); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*entity.Transaction) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTransactionUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewTransactionUsecase creates a new instance of TransactionUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTransactionUsecase(t mockConstructorTestingTNewTransactionUsecase) *TransactionUsecase {
	mock := &TransactionUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
