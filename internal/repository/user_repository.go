package repository

import "simpro/internal/model"

func GetAllUsers() []model.User {
	return []model.User{
		{
			ID:    1,
			Name:  "John Doe",
			Email: "JohnDoe@gmail.com",
		},
		{
			ID:    2,
			Name:  "Muharafi Dalilah",
			Email: "rafidalilah@gmail.com",
		},
	}
}
