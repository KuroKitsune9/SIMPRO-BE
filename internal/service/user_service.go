package service

import (
	"simpro/internal/model"
	"simpro/internal/repository"
)

func GetUsers() []model.User {
	return repository.GetAllUsers()
}
