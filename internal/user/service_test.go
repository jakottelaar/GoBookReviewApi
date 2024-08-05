package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateUserService(t *testing.T) {

	mockRepo := new(MockUserRepository)
	service := NewUserService(mockRepo)

	t.Run("Create user service: Successfully create a user", func(t *testing.T) {
		createReq := &CreateUserRequest{
			Username: "Test User",
			Email:    "testuser@mail.com",
			Password: "password",
		}

		expectedUser := &User{
			ID:       uuid.New(),
			Username: createReq.Username,
			Email:    createReq.Email,
			Password: createReq.Password,
		}

		mockRepo.On("Save", mock.AnythingOfType("*user.User")).Return(expectedUser, nil)

		result, err := service.Create(createReq)

		require.NoError(t, err)
		assert.Equal(t, expectedUser, result)

		mockRepo.AssertExpectations(t)

	})

}
