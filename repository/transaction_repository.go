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
	GetFundNameById(id int) (*entity.Fund, error)
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
	sort := "DESC"
	order := "created_at"
	limit := 10

	query = append(query, "wallet_id = ?")
	values = append(values, walletid)

	if val, ok := params["trans_type"]; ok {
		val = "%%" + val + "%%"
		query = append(query, "trans_type ILIKE ?")
		values = append(values, val)
	}
	if val, ok := params["description"]; ok {
		val = "%%" + val + "%%"
		query = append(query, "description ILIKE ?")
		values = append(values, val)
	}

	temp := r.db.Where(strings.Join(query, " AND "), values...)

	if val, ok := params["limit"]; ok {
		limit, _ = strconv.Atoi(val)
	}
	if val, ok := params["page"]; ok {
		id, _ := strconv.Atoi(val)
		temp.Offset((id - 1) * limit)
	}

	if val, ok := params["sortBy"]; ok {
		order = val
	}
	if val, ok := params["sort"]; ok {
		sort = val
	}

	temp.Preload("Fund").Limit(limit).Order(order + " " + sort).Find(&transactions)
	return transactions, nil
}

func (r *transactionRepository) DoTransaction(e *entity.Transaction) error {
	result := r.db.Create(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *transactionRepository) GetFundNameById(id int) (*entity.Fund, error) {
	var fund *entity.Fund
	result := r.db.Where("fund_id = ?", id).Find(&fund)

	return fund, result.Error
}
