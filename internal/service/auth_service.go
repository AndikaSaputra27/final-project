package service

import (
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(secretKey string) (string, error) {
	// Buat token JWT
	claims := jwt.MapClaims{
		"user_id": 1,                                    // Misalnya ID user
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Expired setelah 1 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Menandatangani token dengan secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return signedToken, nil
}
