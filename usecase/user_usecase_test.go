package usecase

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/mocks"
	"errors"
	"testing"
)

func TestGetUserDetails_NoError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	s := NewUserUsecase(user)

	entityUser := entity.User{
		WalletID: 100001,
		Email:    "test2@shopee.com",
		Name:     "test2",
		Balance:  50000,
	}

	user.On("GetUserDetails", 100001).
		Return(&entityUser, nil)

	e, err := s.GetUserDetails(100001)
	if err != nil {
		t.Errorf("expected no error, got error")
	}
	if e == nil {
		t.Errorf("expected user, got none")
	}
}

func TestGetUserDetails_Error(t *testing.T) {
	user := mocks.NewUserRepository(t)
	s := NewUserUsecase(user)

	user.On("GetUserDetails", 100001).
		Return(nil, errors.New("user not found"))

	e, err := s.GetUserDetails(100001)
	if err == nil {
		t.Errorf("expected error, got none")
	}
	if e != nil {
		t.Errorf("expected none, got user")
	}
}

func TestLoginUser_NoError(t *testing.T) {
	user := mocks.NewUserRepository(t)
	s := NewUserUsecase(user)

	entityToken := entity.Token{
		TokenID: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3Q4QHNob3BlZS5jb20iLCJ3YWxsZXRfaWQiOjEwMDAwNywiZXhwIjoxNjY1NjM1OTAwLCJpYXQiOjE2NjU2MzIzMDB9.ujBs3esW74aRSKhfy6jQZV_B5Bbx6flZFzgfigjpGtc",
	}

	user.On("Login", "test2@shopee.com", "1234").
		Return(&entityToken, nil)

	e, err := s.Login("test2@shopee.com", "1234")
	if err != nil {
		t.Errorf("expected no error, got error")
	}
	if e == nil {
		t.Errorf("expected token, got none")
	}
}
