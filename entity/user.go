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
	//Transaction []Transaction `gorm:"foreignKey:WalletID;references:WalletID"`
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

type UserDetails struct {
	WalletID int    `json:"wallet_id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Balance  int    `json:"balance"`
}
