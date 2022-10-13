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

	entityUser := entity.User{
		WalletID: 100001,
		Password: "$2a$10$SWvAZsaXq4bAlqQCeF6JjeBZqc6vU7OUS29ELmlbIIV07tfQjlkLq",
		Name:     "test2",
		Email:    "test2@shopee.com",
	}

	user.On("GetUserByEmail", "test2@shopee.com").
		Return(&entityUser, 1, nil)

	e, err := s.Login("test2@shopee.com", "1234")
	if err != nil {
		t.Errorf("expected no error, got error")
	}
	if e == nil {
		t.Errorf("expected token, got none")
	}
}

func TestLoginUser_ErrorIDNotFound(t *testing.T) {
	user := mocks.NewUserRepository(t)
	s := NewUserUsecase(user)

	user.On("GetUserByEmail", "test2@shopee.com").
		Return(nil, 0, errors.New("id not found"))

	e, err := s.Login("test2@shopee.com", "1234")
	if err == nil {
		t.Errorf("expected error, got none")
	}
	if e != nil {
		t.Errorf("expected no token, got one")
	}
}

func TestLoginUser_ErrorWrongPassword(t *testing.T) {
	user := mocks.NewUserRepository(t)
	s := NewUserUsecase(user)

	entityUser := entity.User{
		WalletID: 100001,
		Email:    "test2@shopee.com",
		Name:     "test2",
		Balance:  50000,
		Password: "12345",
	}

	user.On("GetUserByEmail", "test2@shopee.com").
		Return(&entityUser, 1, nil)

	e, err := s.Login("test2@shopee.com", "1234")
	if err == nil {
		t.Errorf("expected error, got none")
	}
	if e != nil {
		t.Errorf("expected no token, got one")
	}
}
