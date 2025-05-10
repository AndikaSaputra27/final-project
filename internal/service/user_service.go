package service

import (
	"errors"

	"github.com/AndikaSaputra27/booking-system/internal/config"
	"github.com/AndikaSaputra27/booking-system/internal/entity"
)

func GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
