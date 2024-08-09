package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockService = new(MockAuthService)
var handler = NewAuthHandler(mockService)

func TestRegisterUserHandler(t *testing.T) {

	t.Run("POST Register handler: Successfully register a user", func(t *testing.T) {
		reqBody := RegisterRequest{
			Username: "Test User",
			Email:    "testuser@mail.com",
			Password: "password",
		}

		expectedUser := &user.User{
			ID:        uuid.New(),
			Username:  "Test User",
			Email:     "testuser@mail.com",
			CreatedAt: time.Now(),
		}

		mockService.On("Register", mock.AnythingOfType("*auth.RegisterRequest")).Return(expectedUser, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handler.Register(w, req)

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

func TestLoginUserHandler(t *testing.T) {

	t.Run("POST Login handler: Successfully login a user", func(t *testing.T) {
		reqBody := LoginRequest{
			Email:    "testuser@mail.com",
			Password: "password",
		}

		expected := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

		mockService.On("Login", mock.AnythingOfType("*auth.LoginRequest")).Return(expected, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest("POST", "/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		handler.Login(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockService.AssertExpectations(t)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, expected, response["access_token"])
	})

}
