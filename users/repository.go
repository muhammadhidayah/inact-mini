package users

import (
	"github.com/muhammadhidayah/inact-mini/models"
)

type Repository interface {
	GetUserById(int64) (*models.Users, error)
	GetUserByUsername(string) (*models.Users, error)
	DeleteUserById(int64) error
	UpdateUserById(int64, *models.Users) error
	InsertUser(*models.Users) (int, error)
}
