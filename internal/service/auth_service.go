package service

import (
	"errors"
)

// User adalah struct untuk mendefinisikan user yang melakukan login
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthService adalah service yang menangani autentikasi pengguna
type AuthService struct {
	// Bisa ditambah field lain jika diperlukan, misalnya koneksi ke database
}

// NewAuthService membuat instance baru dari AuthService
func NewAuthService() *AuthService {
	return &AuthService{}
}

// Authenticate memeriksa username dan password untuk autentikasi
func (s *AuthService) Authenticate(username, password string) (*User, error) {
	// Ini hanya contoh sederhana, kamu bisa mengganti logika ini dengan pengecekan ke database
	if username == "admin" && password == "password" {
		return &User{
			Username: username,
			Password: password,
		}, nil
	}
	return nil, errors.New("invalid credentials")
}
