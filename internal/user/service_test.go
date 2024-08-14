package user

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockRepo = new(MockUserRepository)
var service = NewUserService(mockRepo)

func TestFindUserById(t *testing.T) {

	t.Run("Find user by id user service: Successfully fetch user by id", func(t *testing.T) {
		userId := uuid.New()

		expectedUser := &User{
			ID:        userId,
			Username:  "Test User",
			Email:     "testuser@mail.com",
			CreatedAt: time.Now(),
		}

		mockRepo.On("FindById", userId.String()).Return(expectedUser, nil)
		result, err := service.FindById(userId.String())

		require.NoError(t, err)
		assert.Equal(t, expectedUser.ID, result.ID)
		assert.Equal(t, expectedUser.Username, result.Username)
		assert.Equal(t, expectedUser.Email, result.Email)
		assert.NotEmpty(t, result.CreatedAt)

	})

	t.Run("Find user by id user service: User not found", func(t *testing.T) {

		userId := uuid.New()

		mockRepo.On("FindById", userId.String()).Return((*User)(nil), common.ErrNotFound)

		result, err := service.FindById(userId.String())

		require.Error(t, err)
		assert.Equal(t, common.ErrNotFound, err)

		assert.Nil(t, result)

		mockRepo.AssertExpectations(t)

	})

}

func TestUpdateUser(t *testing.T) {

	t.Run("Update user by id user service: Successfully update user", func(t *testing.T) {

		userId := uuid.New()

		updateUserReq := &UpdateUserRequest{
			Username: "Updated User",
			Email:    "updatedtest@mail.com",
		}

		existingUser := &User{
			ID:        userId,
			Username:  "Test User",
			Email:     "test@mail.com",
			CreatedAt: time.Now(),
		}

		updatedUser := &User{
			ID:        userId,
			Username:  updateUserReq.Username,
			Email:     updateUserReq.Email,
			UpdatedAt: time.Now(),
		}

		mockRepo.On("FindById", userId.String()).Return(existingUser, nil)
		mockRepo.On("Update", mock.AnythingOfType("*user.User")).Return(updatedUser, nil)

		result, err := service.Update(userId.String(), updateUserReq)

		require.NoError(t, err)
		assert.Equal(t, updatedUser.ID, result.ID)
		assert.Equal(t, updatedUser.Username, result.Username)
		assert.Equal(t, updatedUser.Email, result.Email)
		assert.NotEmpty(t, result.UpdatedAt)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Update user by id user service: User not found", func(t *testing.T) {

		userId := uuid.New()

		updateUserReq := &UpdateUserRequest{
			Username: "Updated User",
			Email:    "updatedtest@mail.com",
		}

		mockRepo.On("FindById", userId.String()).Return((*User)(nil), common.ErrNotFound)
		result, err := service.Update(userId.String(), updateUserReq)
		require.Error(t, err)

		assert.Equal(t, common.ErrNotFound, err)
		assert.Nil(t, result)
		mockRepo.AssertExpectations(t)

	})

}

func TestDeleteUser(t *testing.T) {

	t.Run("Delete user by id user service: Successfully delete user", func(t *testing.T) {
		userId := uuid.New()

		existingUser := &User{
			ID:        userId,
			Username:  "Test User",
			Email:     "test@mail.com",
			CreatedAt: time.Now(),
		}

		mockRepo.On("FindById", userId.String()).Return(existingUser, nil)
		mockRepo.On("Delete", userId.String()).Return(nil)

		err := service.Delete(userId.String())

		require.NoError(t, err)

		mockRepo.AssertExpectations(t)

	})

	t.Run("Delete user by id user service: User not found", func(t *testing.T) {
		userId := uuid.New()

		mockRepo.On("FindById", userId.String()).Return((*User)(nil), common.ErrNotFound)

		err := service.Delete(userId.String())

		require.Error(t, err)
		assert.Equal(t, common.ErrNotFound, err)

		mockRepo.AssertExpectations(t)

	})

}