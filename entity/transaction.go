package entity

import "time"

type Transaction struct {
	TransactionID int    `gorm:"primaryKey;column:transaction_id"`
	WalletID      int    `json:"wallet_id"`
	TransType     string `json:"trans_type"`
	Amount        int    `json:"amount"`
	SourceID      int    `json:"source_id"`
	TargetID      int    `json:"target_id"`
	FundID        int    `json:"fund_id"`
	Description   string
	CreatedAt     time.Time `json:"created_at"`
}

type Fund struct {
	FundID   int    `gorm:"primaryKey;column:fund_id" json:"fund_id"`
	FundName string `json:"fund_name"`
}

type TopUpInput struct {
	WalletID int `json:"wallet_id"`
	Amount   int `json:"amount"`
	FundID   int `json:"fund_id"`
}

type TransferInput struct {
	WalletID    int    `json:"wallet_id"`
	TargetID    int    `json:"target_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
