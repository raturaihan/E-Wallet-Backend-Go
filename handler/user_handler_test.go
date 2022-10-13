package handler

import (
	"assignment-golang-backend/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUserLogin_NoError(t *testing.T) {
	rr := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodPost, "/login", nil)
	c, _ := gin.CreateTestContext(rr)
	c.Request = r

	s := mocks.NewUserUsecase(t)
	h := NewUserHandler(s)

	h.Login(c)

	if rr.Code == http.StatusOK {
		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
	}

}

// func TestUserRegister_NoError(t *testing.T) {
// 	rr := httptest.NewRecorder()
// 	r := httptest.NewRequest(http.MethodPost, "/login", nil)
// 	c, _ := gin.CreateTestContext(rr)
// 	c.Request = r

// 	s := mocks.NewUserUsecase(t)
// 	h := NewUserHandler(s)

// 	h.Register(c)

// 	if rr.Code == http.StatusOK {
// 		t.Errorf("expected %d, got code %d", rr.Code, http.StatusOK)
// 	}

// }
