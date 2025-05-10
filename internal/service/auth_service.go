package service

import (
	"errors"

	"github.com/AndikaSaputra27/booking-system/internal/entity"
	"github.com/AndikaSaputra27/booking-system/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user *entity.User) (*entity.User, error)
	Login(email, password string) (*entity.User, error)
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Register(user *entity.User) (*entity.User, error) {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)
	return s.userRepo.CreateUser(user)
}

func (s *authService) Login(email, password string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("email not found")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
