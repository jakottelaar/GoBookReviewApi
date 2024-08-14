package user

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type UserService interface {
	FindById(id string) (*User, error)
	Update(id string, user *UpdateUserRequest) (*User, error)
	Delete(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) FindById(id string) (*User, error) {

	user, err := u.repo.FindById(id)

	if err != nil {
		switch {
		case errors.Is(err, common.ErrNotFound):
			return nil, common.ErrNotFound

		default:
			return nil, err
		}
	}

	return user, nil

}

func (u *userService) Update(id string, user *UpdateUserRequest) (*User, error) {
	_, err := u.repo.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrNotFound):
			return nil, common.ErrNotFound
		default:
			return nil, err
		}
	}

	updatedUser := &User{
		ID:       uuid.MustParse(id),
		Username: user.Username,
		Email:    user.Email,
	}

	updatedUser, err = u.repo.Update(updatedUser)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, common.ErrNotFound
		default:
			return nil, err
		}
	}

	return updatedUser, nil

}

func (u *userService) Delete(id string) error {

	_, err := u.repo.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrNotFound):
			return common.ErrNotFound
		default:
			return err
		}
	}

	err = u.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil

}