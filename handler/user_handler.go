package handler

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/usecase"
	"assignment-golang-backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(userusecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		usecase: userusecase,
	}
}

func (h *UserHandler) Login(c *gin.Context) {
	var input entity.UserLogin
	err := c.ShouldBindJSON(&input)

	if err != nil {
		utils.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	res, err := h.usecase.Login(strings.ToLower(input.Email), strings.ToLower(input.Password))

	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (h *UserHandler) Register(c *gin.Context) {
	var input entity.UserRegister
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	newUser := &entity.User{
		Name:     strings.ToLower(input.Name),
		Email:    strings.ToLower(input.Email),
		Password: input.Password,
	}

	res, err := h.usecase.Register(newUser)
	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, http.StatusText(http.StatusOK), res)
}
