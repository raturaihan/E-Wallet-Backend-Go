package repository

import (
	"assignment-golang-backend/entity"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetAllTransactionById(walletid int, params map[string]string) ([]*entity.Transaction, error)
	DoTransaction(e *entity.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{
		db: db,
	}
}

func (r *transactionRepository) GetAllTransactionById(walletid int, params map[string]string) ([]*entity.Transaction, error) {
	transactions := make([]*entity.Transaction, 0)

	query := []string{}
	values := []interface{}{}

	query = append(query, "wallet_id = ?")
	values = append(values, walletid)

	if val, ok := params["trans_type"]; ok {
		val = "%" + val + "%"
		query = append(query, "trans_type = ?")
		values = append(values, val)
	}

	temp := r.db.Where(strings.Join(query, " AND"), values...)

	if val, ok := params["limit"]; ok {
		id, _ := strconv.Atoi(val)
		temp.Limit(id)
	}
	if val, ok := params["page"]; ok {
		id, _ := strconv.Atoi(val)
		temp.Offset(id)
	}

	temp.Find(&transactions)
	return transactions, nil
}

func (r *transactionRepository) DoTransaction(e *entity.Transaction) error {
	result := r.db.Create(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
