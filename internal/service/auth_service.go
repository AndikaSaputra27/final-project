package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(secretKey string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": 1,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
