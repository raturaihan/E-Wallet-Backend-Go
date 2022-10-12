package handler

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/usecase"
	"assignment-golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
}

func NewTransactionHandler(transusecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		usecase: transusecase,
	}
}

func (h *TransactionHandler) TopUpAmount(c *gin.Context) {
	var input entity.TopUpInput
	walletId := c.MustGet("wallet_id")
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	walletIdInt := walletId.(int)
	newTransaction := &entity.Transaction{
		WalletID: walletIdInt,
		Amount:   input.Amount,
		FundID:   input.FundID,
	}

	res, err := h.usecase.TopUpAmount(newTransaction)

	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, http.StatusText(http.StatusOK), res)
}

func (h *TransactionHandler) Transfer(c *gin.Context) {
	var input entity.TransferInput
	walletId := c.MustGet("wallet_id")
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utils.WriteResponse(c, http.StatusBadRequest, err.Error(), nil)
	}

	walletIdInt := walletId.(int)

	newTransaction := &entity.Transaction{
		WalletID:    walletIdInt,
		TargetID:    input.TargetID,
		Amount:      input.Amount,
		Description: input.Description,
	}

	res, err := h.usecase.Transfer(newTransaction)

	if err != nil {
		utils.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.WriteResponse(c, http.StatusOK, http.StatusText(http.StatusOK), res)
}
