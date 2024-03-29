package repository

import (
	"assignment-golang-backend/customerrors.go"
	"assignment-golang-backend/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(e *entity.User) (*entity.User, int, error)
	GetUserDetails(walletid int) (*entity.User, error)
	GetUserByEmail(email string) (*entity.User, int, error)
	UpdateBalanceByWalletID(id, amount int) (*entity.User, error)
	GetUserByID(walletID int) (*entity.User, int, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(e *entity.User) (*entity.User, int, error) {
	result := r.db.Omit("wallet_id", "balance", "created_at").Create(&e)
	return e, int(result.RowsAffected), result.Error
}

func (r *userRepository) GetUserDetails(walletid int) (*entity.User, error) {
	var user *entity.User
	result := r.db.Where("wallet_id = ?", walletid).Find(&user)
	return user, result.Error
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, int, error) {
	var user *entity.User
	result := r.db.Where("email = ?", email).Find(&user)
	return user, int(result.RowsAffected), result.Error
}

func (r *userRepository) UpdateBalanceByWalletID(id, amount int) (*entity.User, error) {
	var user entity.User
	r.db.First(&user, id)

	result := r.db.Model(&user).Update("balance", user.Balance+amount)

	if result.RowsAffected == 0 {
		return nil, &customerrors.NoDataUpdatedError{}
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) GetUserByID(walletID int) (*entity.User, int, error) {
	var user *entity.User
	result := r.db.Where("wallet_id = ?", walletID).Find(&user)
	return user, int(result.RowsAffected), result.Error
}
