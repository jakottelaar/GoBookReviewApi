package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var AccessToken string

func TestRegisterUser(t *testing.T) {

	reqBody := `{
        "username": "test_username",
        "email": "test@mail.com",
        "password": "secretpass1234!"
    }`

	req, err := http.NewRequest("POST", baseAuthEndpointURL+"register", strings.NewReader(reqBody))
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusCreated, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	user, ok := response["user"].(map[string]interface{})
	require.True(t, ok)

	assert.NotEmpty(t, user["id"])
	assert.Equal(t, "test_username", user["username"].(string))
	assert.Equal(t, "test_username", user["username"].(string))
	assert.NotEmpty(t, user["created_at"])

}

func TestLoginUser(t *testing.T) {

	reqBody := `{
		"email": "test@mail.com",
        "password": "secretpass1234!"
	}`

	req, err := http.NewRequest("POST", baseAuthEndpointURL+"login", strings.NewReader(reqBody))
	require.NoError(t, err)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	token, ok := response["access_token"].(string)
	require.True(t, ok)

	AccessToken = token

}

func TestGetUser(t *testing.T) {

	req, err := http.NewRequest("GET", baseUserEndpointURL+"profile", nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	user, ok := response["user"].(map[string]interface{})
	require.True(t, ok)

	assert.Equal(t, "test_username", user["username"])
	assert.Equal(t, "test@mail.com", user["email"])
	assert.NotEmpty(t, user["created_at"])

}

func TestUpdateUser(t *testing.T) {

	reqBody := `{
		"username": "updated_username",
		"email": "updatedtest@mail.com"
	}`

	req, err := http.NewRequest("PUT", baseUserEndpointURL+"profile", strings.NewReader(reqBody))
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	user, ok := response["user"].(map[string]interface{})
	require.True(t, ok)

	assert.Equal(t, "updated_username", user["username"])
	assert.Equal(t, "updatedtest@mail.com", user["email"])
	assert.NotEmpty(t, user["updated_at"])

}

func TestDeleteUser(t *testing.T) {

	req, err := http.NewRequest("DELETE", baseUserEndpointURL, nil)
	require.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+AccessToken)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

}
