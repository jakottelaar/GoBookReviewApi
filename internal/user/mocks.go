package user

import "github.com/stretchr/testify/mock"

type MockUserRepository struct {
	mock.Mock
}

type MockUserService struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Update(user *User) (*User, error) {
	args := m.Called(user)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) GetUserByEmail(email string) (*User, error) {
	args := m.Called(email)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) Update(id string, req *UpdateUserRequest) (*User, error) {
	args := m.Called(id, req)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserService) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
