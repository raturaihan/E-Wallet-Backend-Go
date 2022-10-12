package repository

import (
	"assignment-golang-backend/entity"
	"fmt"

	"gorm.io/gorm"
)

const WalletID = 100000

type UserRepository interface {
	CreateUser(e *entity.User) (*entity.User, int, error)
	GetUser(name string) ([]*entity.User, int, error)
	GetUserByEmail(email string) (*entity.User, int, error)
	IsUserExist(email string) (bool, error)
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
	result := r.db.Create(&e)
	return e, int(result.RowsAffected), result.Error
}

func (r *userRepository) GetUser(name string) ([]*entity.User, int, error) {
	var users []*entity.User
	result := r.db.Where("name = ?", name).Find(&users)
	return users, int(result.RowsAffected), result.Error
}

func (r *userRepository) GetUserByEmail(email string) (*entity.User, int, error) {
	var user *entity.User
	result := r.db.Where("email = ?", email).Find(&user)
	return user, int(result.RowsAffected), result.Error
}

func (r *userRepository) IsUserExist(email string) (bool, error) {
	var count int64

	err := r.db.Model(&entity.User{}).
		Where("email ILIKE ?", fmt.Sprintf("%%%s%%", email)).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
