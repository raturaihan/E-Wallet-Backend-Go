package utils

import (
	"assignment-golang-backend/customerrors.go"
	"assignment-golang-backend/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	secret      string        = "topsecretveryconfidential"
	jwtDuration time.Duration = 1 * time.Hour
)

type CustomClaim struct {
	User *entity.UserToken `json:"user"`
	jwt.RegisteredClaims
}

func GenerateJWT(user *entity.UserToken) (string, error) {
	now := time.Now()

	claims := &CustomClaim{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
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

func CheckToken(input string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(input, &CustomClaim{}, func(input *jwt.Token) (interface{}, error) {
		if _, ok := input.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, &customerrors.InvalidTokenError{}
		}

		return []byte(secret), nil
	})

}

func ParseAuthorizationHeader(authHeader string) (string, error) {
	authHeaderSplit := strings.Split(authHeader, "Bearer ")
	if len(authHeaderSplit) != 2 {
		return "", &customerrors.AuthHeaderUnavailable{}
	}

	return strings.TrimSpace(authHeaderSplit[1]), nil
}
