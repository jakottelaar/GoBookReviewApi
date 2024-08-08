package auth

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockRepo = new(user.MockUserRepository)
var service = NewAuthService(mockRepo)

func TestRegisterService(t *testing.T) {

	t.Run("Register auth service: Successfully register a user", func(t *testing.T) {
		createReq := &RegisterRequest{
			Username: "Test User",
			Email:    "testuser@mail.com",
			Password: "password",
		}

		expectedUser := &user.User{
			ID:       uuid.New(),
			Username: createReq.Username,
			Email:    createReq.Email,
			Password: createReq.Password,
		}

		mockRepo.On("Save", mock.AnythingOfType("*user.User")).Return(expectedUser, nil)

		result, err := service.Register(createReq)

		require.NoError(t, err)
		assert.Equal(t, expectedUser, result)
		assert.Equal(t, createReq.Username, result.Username)
		assert.Equal(t, createReq.Email, result.Email)
		assert.NotEmpty(t, result.Password)

		mockRepo.AssertExpectations(t)

	})

}

func TestLoginService(t *testing.T) {

	t.Run("Login auth service: Successfully login a user", func(t *testing.T) {
		loginReq := &LoginRequest{
			Email:    "testuser@mail.com",
			Password: "password",
		}

		expectedUser := &user.User{
			ID:       uuid.New(),
			Password: "$argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG",
		}

		mockRepo.On("FindByEmail", "testuser@mail.com").Return(expectedUser, nil)

		result, err := service.Login(loginReq)

		require.NoError(t, err)
		assert.NotEmpty(t, result)

		mockRepo.AssertExpectations(t)
	})
}
