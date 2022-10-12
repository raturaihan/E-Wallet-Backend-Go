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

func (h *TransactionHandler) GetAllTransaction(c *gin.Context) {
	params := make(map[string]string)
	walletId := c.MustGet("wallet_id")
	walletIdInt := walletId.(int)

	if c.Query("limit") != "" {
		params["limit"] = c.Query("limit")
	}

	if c.Query("page") != "" {
		params["page"] = c.Query("page")
	}

	if c.Query("trans_type") != "" {
		params["trans_type"] = c.Query("trans_type")
	}

	if c.Query("description") != "" {
		params["description"] = c.Query("description")
	}

	if c.Query("sortBy") != "" {
		params["sortBy"] = c.Query("sortBy")
	}

	tl, err := h.usecase.GetAllTransactionById(walletIdInt, params)

	if err != nil {
		utils.WriteErrorResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	utils.WriteErrorResponse(c, http.StatusOK, http.StatusText(http.StatusOK), tl)
}
