package entity

import (
	"time"
)

type User struct {
	WalletID  int    `gorm:"primaryKey;column:wallet_id" json:"wallet_id"`
	Password  string `json:"password,omitempty"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Balance   int    `json:"balance"`
	CreatedAt time.Time
}

type UserToken struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
