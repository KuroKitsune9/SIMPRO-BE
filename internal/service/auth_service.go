package service

import (
	"errors"
	dto "simpro/internal/dto/auth"
	"simpro/internal/model"
	"simpro/internal/repository"
	"simpro/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(req dto.RegisterRequest) error {

	existingUser, _ := repository.FindByUsername(req.Username)

	if existingUser != nil {
		return errors.New("username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user := model.User{
		Name:       req.Name,
		Username:   req.Username,
		Password:   string(hashedPassword),
		Role:       req.Role,
		BranchCode: req.BranchCode,
	}

	return repository.CreateUser(&user)
}

func Login(req dto.LoginRequest) (string, error) {
	user, err := repository.FindByUsername(req.Username)

	if err != nil {
		return "", errors.New("Username atau password salah")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return "", errors.New("Username atau password salah")
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.Username,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}
