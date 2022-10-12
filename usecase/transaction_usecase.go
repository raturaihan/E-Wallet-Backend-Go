package usecase

import (
	"assignment-golang-backend/customerrors.go"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	TopUpAmount(e *entity.Transaction) (*entity.Transaction, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
}

func NewTransactionUsecase(transrepo repository.TransactionRepository, userrepo repository.UserRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: transrepo,
		userRepo:        userrepo,
	}
}

func (u *transactionUsecase) TopUpAmount(e *entity.Transaction) (*entity.Transaction, error) {
	wallet, rowsAffected, err := u.userRepo.GetUserByID(e.WalletID)

	if rowsAffected == 0 {
		return nil, &customerrors.NoDataFoundError{}
	}
	if err != nil {
		return nil, err
	}

	transaction := &entity.Transaction{
		WalletID:    wallet.WalletID,
		TransType:   "TOPUP",
		Amount:      e.Amount,
		FundID:      e.FundID,
		Description: u.GenerateDescription(e.FundID),
	}

	res := u.transactionRepo.DoTransaction(transaction)
	if res != nil {
		return nil, res
	}

	_, err = u.userRepo.UpdateBalanceByWalletID(e.WalletID, e.Amount)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (u *transactionUsecase) GenerateDescription(fundId int) string {
	if fundId == 1 {
		return "Top up from bank transfer"
	}
	if fundId == 2 {
		return "Top up from credit card"
	}
	if fundId == 3 {
		return "Top up from cash"
	}
	return ""
}
