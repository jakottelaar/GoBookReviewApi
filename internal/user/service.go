package user

import "github.com/google/uuid"

type UserService interface {
	Create(user *CreateUserRequest) (*User, error)
	GetUserByEmail(email string) (*User, error)
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

func (s *userService) Create(user *CreateUserRequest) (*User, error) {
	newId := uuid.New()

	newUser := &User{
		ID:       newId,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	savedUser, err := s.repo.Save(newUser)

	if err != nil {
		return nil, err
	}

	return savedUser, nil

}
func (u *userService) GetUserByEmail(email string) (*User, error) {
	panic("unimplemented")
}

func (u *userService) Update(id string, user *UpdateUserRequest) (*User, error) {
	panic("unimplemented")
}

func (u *userService) Delete(id string) error {
	panic("unimplemented")
}
