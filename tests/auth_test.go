package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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

	require.Equal(t, http.StatusCreated, res.StatusCode)

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

	require.Equal(t, http.StatusOK, res.StatusCode)

}
