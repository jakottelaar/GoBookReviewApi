package user

type UserService interface {
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

func (u *userService) GetUserByEmail(email string) (*User, error) {
	panic("unimplemented")
}

func (u *userService) Update(id string, user *UpdateUserRequest) (*User, error) {
	panic("unimplemented")
}

func (u *userService) Delete(id string) error {
	panic("unimplemented")
}
