package auth

import (
	"testing"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRegisterService(t *testing.T) {

	mockRepo := new(user.MockUserRepository)
	service := NewAuthService(mockRepo)

	t.Run("Register user service: Successfully register a user", func(t *testing.T) {
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
