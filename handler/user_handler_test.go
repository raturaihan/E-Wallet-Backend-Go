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

func TestUserLogin_ErrorNoInput(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	s := mocks.NewUserUsecase(t)
	h := NewUserHandler(s)

	h.Login(c)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusBadRequest)
	}

}

func TestUserLogin_NoError(t *testing.T) {
	rr := httptest.NewRecorder()
	jsonBody := `{
        "email":"test@test.com",
        "password":"1234"
        }`

	entityToken := entity.Token{
		TokenID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHNob3BlZS5jb20iLCJ3YWxsZXRfaWQiOjEwMDAwMCwiZXhwIjoxNjY1NjU5ODE3LCJpYXQiOjE2NjU2NTYyMTd9.9qyPyGr4NWYfErr27Ot47TLQlx_LqcluD3xfmeYRwX8",
	}

	body := strings.NewReader(jsonBody)
	r := httptest.NewRequest(http.MethodPost, "/login", body)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	s := mocks.NewUserUsecase(t)
	h := NewUserHandler(s)

	s.On("Login", "test@test.com", "1234").
		Return(&entityToken, nil)
	h.Login(c)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}

}

func TestUserRegister_Error(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/register", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	s := mocks.NewUserUsecase(t)
	h := NewUserHandler(s)

	newUser := &entity.User{
		Name:     "",
		Email:    "",
		Password: "",
	}

	s.On("Register", newUser).
		Return(nil, errors.New("empty body"))
	h.Register(c)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusBadRequest)
	}

}

func TestUserRegister_NoError(t *testing.T) {
	rr := httptest.NewRecorder()
	newUser := &entity.User{
		Password: "1234",
		Email:    "tafia@test.com",
		Name:     "tafia",
	}

	jsonBody := `{
		"password":"1234",
		"email": "tafia@test.com",
		"name":"tafia"
    }`

	entityToken := entity.Token{
		TokenID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3QxQHNob3BlZS5jb20iLCJ3YWxsZXRfaWQiOjEwMDAwMCwiZXhwIjoxNjY1NjU5ODE3LCJpYXQiOjE2NjU2NTYyMTd9.9qyPyGr4NWYfErr27Ot47TLQlx_LqcluD3xfmeYRwX8",
	}

	body := strings.NewReader(jsonBody)
	r := httptest.NewRequest(http.MethodPost, "/register", body)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	s := mocks.NewUserUsecase(t)
	h := NewUserHandler(s)

	s.On("Register", newUser).
		Return(&entityToken, nil)
	h.Register(c)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}

}

func TestGetUserDetails_NoError(t *testing.T) {
	rr := httptest.NewRecorder()

	r := httptest.NewRequest(http.MethodPost, "/user", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	c.Set("wallet_id", 100001)

	entityUserToken := entity.User{
		Name:     "test1",
		Email:    "test1@shopee.com",
		Password: "1234",
		WalletID: 100001,
	}

	s := mocks.NewUserUsecase(t)
	h := NewUserHandler(s)

	s.On("GetUserDetails", 100001).
		Return(&entityUserToken, nil)
	h.GetUserDetails(c)

	if rr.Code != http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}
}
