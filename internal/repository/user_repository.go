package repository

import (
	"simpro/config"
	"simpro/internal/model"
)

func FindByUsername(username string) (*model.User, error) {
	var user model.User

	err := config.DB.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUser(user *model.User) error {
	return config.DB.Create(user).Error
}
