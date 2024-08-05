package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockService = new(MockUserService)
var handler = NewUserHandler(mockService)

func TestCreateUserHandler(t *testing.T) {

	t.Run("POST User handler: Successfully create a user", func(t *testing.T) {
		reqBody := CreateUserRequest{
			Username: "Test User",
			Email:    "testuser@mail.com",
			Password: "password",
		}

		expectedUser := &User{
			ID:        uuid.New(),
			Username:  "Test User",
			Email:     "testuser@mail.com",
			CreatedAt: time.Now(),
		}

		mockService.On("Create", mock.AnythingOfType("*user.CreateUserRequest")).Return(expectedUser, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/v1/api/users", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handler.CreateUser(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		user, ok := response["user"].(map[string]interface{})

		assert.True(t, ok)
		assert.Equal(t, expectedUser.ID.String(), user["id"])
		assert.Equal(t, expectedUser.Username, user["username"])
		assert.Equal(t, expectedUser.Email, user["email"])
		assert.NotEmpty(t, user["created_at"])

	})

}
