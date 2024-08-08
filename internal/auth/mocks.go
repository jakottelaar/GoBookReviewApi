package auth

import (
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(userReq *RegisterRequest) (*user.User, error) {
	args := m.Called(userReq)
	return args.Get(0).(*user.User), args.Error(1)
}

func (m *MockAuthService) Login(loginReq *LoginRequest) (string, error) {
	args := m.Called(loginReq)
	return args.Get(0).(string), args.Error(1)
}
