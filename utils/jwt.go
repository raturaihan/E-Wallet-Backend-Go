package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secret      string        = "topsecretveryconfidential"
	jwtDuration time.Duration = 1 * time.Hour
)

type CustomClaim struct {
	Email    string `json:"email"`
	WalletID int    `json:"wallet_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(email string, wallet_id int) (string, error) {
	now := time.Now()

	claims := CustomClaim{
		email,
		wallet_id,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(jwtDuration)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func CheckToken(input string) (string, int, error) {
	token, err := jwt.ParseWithClaims(input, &CustomClaim{}, func(tkn *jwt.Token) (any, error) {
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claims.Email, claims.WalletID, nil
	} else {
		return "", 0, err
	}
}
