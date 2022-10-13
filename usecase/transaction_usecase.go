package usecase

import (
	"assignment-golang-backend/customerrors.go"
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
)

type TransactionUsecase interface {
	TopUpAmount(e *entity.Transaction) (*entity.Transaction, error)
	Transfer(e *entity.Transaction) (*entity.Transaction, error)
	GetAllTransactionById(walletid int, params map[string]string) ([]*entity.Transaction, error)
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
		SourceID:    wallet.WalletID,
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

func (u *transactionUsecase) Transfer(e *entity.Transaction) (*entity.Transaction, error) {
	source, sourceRow, sourceErr := u.userRepo.GetUserByID(e.WalletID)
	if sourceRow == 0 {
		return nil, &customerrors.NoDataFoundError{}
	}
	if sourceErr != nil {
		return nil, sourceErr
	}

	target, targetRow, targetErr := u.userRepo.GetUserByID(e.TargetID)
	if targetRow == 0 {
		return nil, &customerrors.TargetNotFoundError{}
	}
	if targetErr != nil {
		return nil, targetErr
	}

	if e.TargetID == e.WalletID {
		return nil, &customerrors.TransferFailed{}
	}
	if source.Balance < e.Amount {
		return nil, &customerrors.InsufficientBalanceError{}
	}

	transaction1 := &entity.Transaction{
		WalletID:    source.WalletID,
		TransType:   "TRANSFER",
		TargetID:    e.TargetID,
		Amount:      e.Amount,
		Description: e.Description,
	}

	transaction2 := &entity.Transaction{
		WalletID:    e.TargetID,
		TransType:   "RECEIVED TRANSFER",
		SourceID:    source.WalletID,
		Amount:      e.Amount,
		Description: e.Description,
	}

	res1 := u.transactionRepo.DoTransaction(transaction1)
	if res1 != nil {
		return nil, res1
	}

	res2 := u.transactionRepo.DoTransaction(transaction2)
	if res2 != nil {
		return nil, res2
	}

	_, errSource := u.userRepo.UpdateBalanceByWalletID(source.WalletID, -e.Amount)
	if errSource != nil {
		return nil, errSource
	}

	_, errTarget := u.userRepo.UpdateBalanceByWalletID(target.WalletID, e.Amount)
	if errTarget != nil {
		return nil, errTarget
	}

	return transaction1, nil
}

func (u *transactionUsecase) GetAllTransactionById(walletid int, params map[string]string) ([]*entity.Transaction, error) {
	tl, err := u.transactionRepo.GetAllTransactionById(walletid, params)
	if err != nil {
		return nil, err
	}
	return tl, nil
}

func (u *transactionUsecase) GenerateDescription(fundId int) string {
	source, err := u.transactionRepo.GetFundNameById(fundId)
	if err != nil {
		return ""
	}
	return "Top up from " + source.FundName
}
