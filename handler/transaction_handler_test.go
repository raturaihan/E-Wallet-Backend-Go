package handler

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/mocks"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestTopUpAmount_NoError(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonBody := `{
		"wallet_id":100001,
		"amount": 100000,
		"fund_id":1
    }`

	body := strings.NewReader(jsonBody)

	r := httptest.NewRequest(http.MethodPost, "/user/topup", body)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	s := mocks.NewTransactionUsecase(t)
	h := NewTransactionHandler(s)

	entTrans := entity.Transaction{
		WalletID: 100001,
		Amount:   100000,
		FundID:   1,
	}

	s.On("TopUpAmount", &entTrans).
		Return(&entTrans, nil)
	h.TopUpAmount(c)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}

}

func TestTopUpAmount_AmountError(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonBody := `{
		"wallet_id":100001,
		"amount": 10000,
		"fund_id":1
    }`

	body := strings.NewReader(jsonBody)

	r := httptest.NewRequest(http.MethodPost, "/user/topup", body)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	s := mocks.NewTransactionUsecase(t)
	h := NewTransactionHandler(s)

	entTrans := entity.Transaction{
		WalletID: 100001,
		Amount:   10000,
		FundID:   1,
	}

	s.On("TopUpAmount", &entTrans).
		Return(nil, errors.New("failed top up"))
	h.TopUpAmount(c)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusInternalServerError)
	}

}

func TestTransfer_NoError(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonBody := `{
		"wallet_id":100001,
		"amount": 100000,
		"target_id":100002
    }`

	body := strings.NewReader(jsonBody)

	r := httptest.NewRequest(http.MethodPost, "/user/transfer", body)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	s := mocks.NewTransactionUsecase(t)
	h := NewTransactionHandler(s)

	entTrans := entity.Transaction{
		WalletID: 100001,
		Amount:   100000,
		TargetID: 100002,
	}

	s.On("Transfer", &entTrans).
		Return(&entTrans, nil)
	h.Transfer(c)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}

}

func TestTransfer_AmountError(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonBody := `{
		"wallet_id":100001,
		"amount": 500,
		"target_id":100002
    }`

	body := strings.NewReader(jsonBody)

	r := httptest.NewRequest(http.MethodPost, "/user/transfer", body)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	s := mocks.NewTransactionUsecase(t)
	h := NewTransactionHandler(s)

	entTrans := entity.Transaction{
		WalletID: 100001,
		Amount:   500,
		TargetID: 100002,
	}

	s.On("Transfer", &entTrans).
		Return(nil, errors.New("failed transfer"))
	h.Transfer(c)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusInternalServerError)
	}

}

func TestGetAllTransaction_NoError(t *testing.T) {
	params := make(map[string]string)
	rr := httptest.NewRecorder()

	r := httptest.NewRequest(http.MethodPost, "/user/transaction", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	s := mocks.NewTransactionUsecase(t)
	h := NewTransactionHandler(s)

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

	s.On("GetAllTransactionById", 100001, params).
		Return(entityTrans, nil)
	h.GetAllTransaction(c)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}

}

func TestGetAllTransaction_Error(t *testing.T) {
	params := make(map[string]string)
	rr := httptest.NewRecorder()

	r := httptest.NewRequest(http.MethodPost, "/user/transaction", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	s := mocks.NewTransactionUsecase(t)
	h := NewTransactionHandler(s)

	s.On("GetAllTransactionById", 100001, params).
		Return(nil, errors.New("id not found"))
	h.GetAllTransaction(c)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusInternalServerError)
	}

}
