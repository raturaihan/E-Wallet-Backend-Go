package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllTransactionById_NoError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	params := make(map[string]string)
	entityTrans := []*entity.Transaction{
		{
			WalletID:  100001,
			TransType: "TOPUP",
			Amount:    50000,
			SourceID:  100001,
			FundID:    1,
		},
		{
			WalletID:  100001,
			TransType: "TOPUP",
			Amount:    50000,
			SourceID:  100001,
			FundID:    1,
		},
	}
	trans.On("GetAllTransactionById", 100001, params).
		Return(entityTrans, nil)

	e, err := s.GetAllTransactionById(100001, params)
	if err != nil {
		t.Errorf("expected no error, got error")
	}
	if e == nil {
		t.Errorf("expected transactions, got none")
	}

}

func TestGetAllTransactionById_Error(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	params := make(map[string]string)

	trans.On("GetAllTransactionById", 100001, params).
		Return(nil, errors.New("id not found"))

	e, err := s.GetAllTransactionById(100001, params)
	if err == nil {
		t.Errorf("expected error, got none")
	}
	if e != nil {
		t.Errorf("expected none, got transactions")
	}

}

func TestGenerateDescription_NoError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	entityFund := entity.Fund{
		FundName: "bank transfer",
	}

	trans.On("GetFundNameById", 1).
		Return(&entityFund, nil)

	res := s.GenerateDescription(1)
	if res == "" {
		t.Errorf("expected string, got none")
	}
}

func TestGenerateDescription_Error(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	trans.On("GetFundNameById", 1).
		Return(nil, errors.New("id not found"))

	res := s.GenerateDescription(1)
	if res != "" {
		t.Errorf("expected not string, got one")
	}
}

func TestTopUpAmount_ErrorIDNotFound(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	entityTrans := entity.Transaction{
		WalletID:  100001,
		SourceID:  100001,
		TransType: "TOPUP",
		Amount:    50000,
		FundID:    1,
	}

	user.On("GetUserByID", 100001).
		Return(nil, 0, errors.New("id not found"))

	trans.On("DoTransaction", &entityTrans).
		Return(errors.New("failed top up"))

	user.On("UpdateBalanceByWalletID", 100001, 50000).
		Return(nil, errors.New("error update balance"))

	res1, rows, err1 := user.GetUserByID(100001)
	if res1 != nil {
		t.Errorf("expected none, got one transaction")
	}
	if rows != 0 {
		t.Errorf("expected 0 rows affected, got one")
	}
	if err1 == nil {
		t.Errorf("expected error, got none")
	}

	res2, err2 := s.TopUpAmount(&entityTrans)
	if err2 == nil {
		t.Errorf("expected error, got none")
	}
	if res2 != nil {
		t.Errorf("expected no transaction, one")
	}

	err3 := trans.DoTransaction(&entityTrans)
	if err3 == nil {
		t.Errorf("expected error, got none")
	}

	_, err4 := user.UpdateBalanceByWalletID(100001, 50000)
	if err4 == nil {
		t.Errorf("expected error, got none")
	}
}

func TestTopUpAmount_NoError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	entityTrans := entity.Transaction{
		WalletID:    100001,
		SourceID:    100001,
		TransType:   "TOPUP",
		Amount:      50000,
		FundID:      1,
		Description: "Top up from bank transfer",
	}

	entityUser := entity.User{
		WalletID: 100001,
		Email:    "test2@shopee.com",
		Name:     "test2",
		Balance:  50000,
	}

	entityFund := entity.Fund{
		FundName: "bank transfer",
	}

	user.On("GetUserByID", 100001).
		Return(&entityUser, 1, nil)

	trans.On("GetFundNameById", 1).
		Return(&entityFund, nil)

	trans.On("DoTransaction", &entityTrans).
		Return(nil)

	user.On("UpdateBalanceByWalletID", 100001, 50000).
		Return(&entityUser, nil)

	res1, rows, err1 := user.GetUserByID(100001)
	if res1 == nil {
		t.Errorf("expected one transaction, got none")
	}
	if rows == 0 {
		t.Errorf("expected 1 rows affected, got 0")
	}
	if err1 != nil {
		t.Errorf("expected nil, got error")
	}

	res2, err2 := s.TopUpAmount(&entityTrans)
	if err2 != nil {
		t.Errorf("expected nil, got error")
	}
	if res2 == nil {
		t.Errorf("expected transaction, none")
	}

	err3 := trans.DoTransaction(&entityTrans)
	if err3 != nil {
		t.Errorf("expected nil, got error")
	}

	_, err4 := user.UpdateBalanceByWalletID(100001, 50000)
	if err4 != nil {
		t.Errorf("expected no error, got one")
	}
}

func TestTransfer_NoError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	transaction1 := entity.Transaction{
		WalletID:    100001,
		TransType:   "TRANSFER",
		TargetID:    100002,
		Amount:      10000,
		Description: "sedekah",
	}

	transaction2 := entity.Transaction{
		WalletID:    100002,
		TransType:   "RECEIVED TRANSFER",
		SourceID:    100001,
		Amount:      10000,
		Description: "sedekah",
	}

	entityUser1 := entity.User{
		WalletID: 100001,
		Email:    "test2@shopee.com",
		Name:     "test2",
		Balance:  100000,
	}

	entityUser2 := entity.User{
		WalletID: 100002,
		Email:    "test3@shopee.com",
		Name:     "test3",
	}

	user.On("GetUserByID", 100001).
		Return(&entityUser1, 1, nil)

	user.On("GetUserByID", 100002).
		Return(&entityUser2, 1, nil)

	trans.On("DoTransaction", &transaction1).
		Return(nil)

	trans.On("DoTransaction", &transaction2).
		Return(nil)

	user.On("UpdateBalanceByWalletID", 100001, -10000).
		Return(&entityUser1, nil)

	user.On("UpdateBalanceByWalletID", 100002, 10000).
		Return(&entityUser2, nil)

	source, row, err := user.GetUserByID(100001)
	assert.NotNil(t, source)
	assert.Equal(t, 1, row)
	assert.Nil(t, err)

	source2, row2, err2 := user.GetUserByID(100002)
	assert.NotNil(t, source2)
	assert.Equal(t, 1, row2)
	assert.Nil(t, err2)

	err3 := trans.DoTransaction(&transaction1)
	assert.Nil(t, err3)

	err4 := trans.DoTransaction(&transaction2)
	assert.Nil(t, err4)

	_, err5 := user.UpdateBalanceByWalletID(100002, 10000)
	assert.Nil(t, err5)

	_, err6 := user.UpdateBalanceByWalletID(100001, -10000)
	assert.Nil(t, err6)

	res, err := s.Transfer(&transaction1)
	assert.NotNil(t, res)
	assert.Nil(t, err)

}

func TestTransfer_Error(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	transaction := entity.Transaction{
		WalletID:    100001,
		TransType:   "TRANSFER",
		TargetID:    100002,
		Amount:      10000,
		Description: "utang traktiran",
	}

	user.On("GetUserByID", 100001).
		Return(nil, 0, errors.New("user not found"))

	user.On("GetUserByID", 100002).
		Return(nil, 0, errors.New("target not found"))

	trans.On("DoTransaction", &transaction).
		Return(errors.New("failed transfer"))

	user.On("UpdateBalanceByWalletID", 100001, -10000).
		Return(nil, errors.New("failed transfer"))

	user.On("UpdateBalanceByWalletID", 100002, 10000).
		Return(nil, errors.New("failed transfer"))

	source, row, err := user.GetUserByID(100001)
	assert.Nil(t, source)
	assert.Equal(t, 0, row)
	assert.Error(t, err)

	source2, row2, err2 := user.GetUserByID(100002)
	assert.Nil(t, source2)
	assert.Equal(t, 0, row2)
	assert.Error(t, err2)

	err3 := trans.DoTransaction(&transaction)
	assert.Error(t, err3)

	res, err4 := s.Transfer(&transaction)
	assert.Nil(t, res)
	assert.Error(t, err4)

	_, err5 := user.UpdateBalanceByWalletID(100002, 10000)
	assert.Error(t, err5)

	_, err6 := user.UpdateBalanceByWalletID(100001, -10000)
	assert.Error(t, err6)
}

func TestTransfer_InsufficientBalanceError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	trans := mocks.NewTransactionRepository(t)
	s := NewTransactionUsecase(trans, user)

	transaction1 := entity.Transaction{
		WalletID:    100001,
		TransType:   "TRANSFER",
		TargetID:    100002,
		Amount:      10000,
		Description: "sedekah",
	}

	transaction2 := entity.Transaction{
		WalletID:    100002,
		TransType:   "RECEIVED TRANSFER",
		SourceID:    100001,
		Amount:      10000,
		Description: "sedekah",
	}

	entityUser1 := entity.User{
		WalletID: 100001,
		Email:    "test2@shopee.com",
		Name:     "test2",
		Balance:  5000,
	}

	entityUser2 := entity.User{
		WalletID: 100002,
		Email:    "test3@shopee.com",
		Name:     "test3",
	}

	user.On("GetUserByID", 100001).
		Return(&entityUser1, 1, nil)

	user.On("GetUserByID", 100002).
		Return(&entityUser2, 1, nil)

	trans.On("DoTransaction", &transaction1).
		Return(nil)

	trans.On("DoTransaction", &transaction2).
		Return(nil)

	user.On("UpdateBalanceByWalletID", 100001, -10000).
		Return(&entityUser1, nil)

	user.On("UpdateBalanceByWalletID", 100002, 10000).
		Return(&entityUser2, nil)

	source, row, err := user.GetUserByID(100001)
	assert.NotNil(t, source)
	assert.Equal(t, 1, row)
	assert.Nil(t, err)

	source2, row2, err2 := user.GetUserByID(100002)
	assert.NotNil(t, source2)
	assert.Equal(t, 1, row2)
	assert.Nil(t, err2)

	err3 := trans.DoTransaction(&transaction1)
	assert.Nil(t, err3)

	err4 := trans.DoTransaction(&transaction2)
	assert.Nil(t, err4)

	_, err5 := user.UpdateBalanceByWalletID(100002, 10000)
	assert.Nil(t, err5)

	_, err6 := user.UpdateBalanceByWalletID(100001, -10000)
	assert.Nil(t, err6)

	res, err := s.Transfer(&transaction1)
	assert.Nil(t, res)
	assert.NotNil(t, err)

}
