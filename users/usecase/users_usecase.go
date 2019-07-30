package usecase

import (
	"github.com/muhammadhidayah/inact-mini/models"
	"github.com/muhammadhidayah/inact-mini/users"
)

type userUsecase struct {
	userRepo users.Repository
}

func NewUserUsecase(userRepo users.Repository) users.Usecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) GetUserById(id int64) (*models.Users, error) {
	res, err := u.userRepo.GetUserById(id)
	if err != nil {
		return nil, err

	}
	return res, nil
}

func (u *userUsecase) InsertUser(data *models.Users) (int, error) {
	lastInserID, err := u.userRepo.InsertUser(data)
	return lastInserID, err
}

func (u *userUsecase) DeleteUserById(id int64) error {
	return nil
}

func (u *userUsecase) UpdateUserById(id int64, data *models.Users) error {
	return nil
}

func (u *userUsecase) GetUserByUsername(username string) (*models.Users, error) {
	res, err := u.userRepo.GetUserByUsername(username)

	if err != nil {
		return nil, err
	}

	return res, nil
}
