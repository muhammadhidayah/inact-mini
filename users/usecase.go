package users

import "github.com/muhammadhidayah/inact-mini/models"

type Usecase interface {
	GetUserById(id int64) (*models.Users, error)
	DeleteUserById(id int64) error
	UpdateUserById(id int64, data *models.Users) error
	InsertUser(data *models.Users) (int, error)
}
