package auth

import (
	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type AuthService interface {
	Login(req *LoginRequest) (*LoginResponse, error)
	Register(req *RegisterRequest) (*user.User, error)
}

type authService struct {
	repo user.UserRepository
}

func NewAuthService(repo user.UserRepository) AuthService {
	return &authService{
		repo: repo,
	}
}

func (s *authService) Login(req *LoginRequest) (*LoginResponse, error) {
	panic("unimplemented")
}

func (s *authService) Register(userReq *RegisterRequest) (*user.User, error) {

	newId := uuid.New()

	// Add password hashing here
	hashedPassword, err := common.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}

	newUser := &user.User{
		ID:       newId,
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: string(hashedPassword),
	}

	savedUser, err := s.repo.Save(newUser)

	if err != nil {
		return nil, err
	}

	return savedUser, nil

}
