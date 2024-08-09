package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var bookUserToken string
var bookId string

func TestCreateUser(t *testing.T) {

	t.Run("Register new book user", func(t *testing.T) {
		reqBody := `{
			"username": "test_book_user",
			"email": "test_book@mail.com",
			"password": "password1234!"
		}`

		req, err := http.NewRequest("POST", baseAuthEndpointURL+"register", strings.NewReader(reqBody))
		require.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer res.Body.Close()

		require.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("Login book user", func(t *testing.T) {
		reqBody := `{
			"email": "test_book@mail.com",
			"password": "password1234!"
		}`

		req, err := http.NewRequest("POST", baseAuthEndpointURL+"login", strings.NewReader(reqBody))
		require.NoError(t, err)

		res, err := http.DefaultClient.Do(req)
		require.NoError(t, err)
		defer res.Body.Close()

		require.Equal(t, http.StatusOK, res.StatusCode)

		var response map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&response)
		require.NoError(t, err)

		token, ok := response["access_token"].(string)
		assert.True(t, ok)

		bookUserToken = token
	})
}

func TestCreateBookRequest(t *testing.T) {
	reqBody := `{
		"title": "Book Title",
		"author": "Book Author",
		"published_year": 2020,
		"isbn": "9780743273565"
	}`

	req, err := http.NewRequest("POST", baseBooksEndpointURL, strings.NewReader(reqBody))
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+bookUserToken)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusCreated, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	book, ok := response["book"].(map[string]interface{})
	require.True(t, ok)

	bookId = book["id"].(string)
	assert.NotEmpty(t, bookId)
	assert.Equal(t, "Book Title", book["title"].(string))
	assert.Equal(t, "Book Author", book["author"].(string))
	assert.Equal(t, float64(2020), book["published_year"].(float64))
	assert.Equal(t, "9780743273565", book["isbn"].(string))
	assert.Contains(t, book, "created_at")
}

func TestGetBookByIdRequest(t *testing.T) {

	req, err := http.NewRequest("GET", baseBooksEndpointURL+bookId, nil)
	require.NoError(t, err)
	req.Header.Set("Authorization", "Bearer "+bookUserToken)

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	book, ok := response["book"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, bookId, book["id"].(string))
	assert.Equal(t, "Book Title", book["title"].(string))
	assert.Equal(t, "Book Author", book["author"].(string))
	assert.Equal(t, float64(2020), book["published_year"].(float64))
	assert.Equal(t, "9780743273565", book["isbn"].(string))
	assert.Contains(t, book, "created_at")
	assert.Contains(t, book, "user_id")
}

// func TestUpdateBookById(t *testing.T) {
// 	updateReqBody := `{
// 		"title": "Updated Book Title",
// 		"author": "Updated Book Author",
// 		"published_year": 2021,
// 		"isbn": "9780743273565"
// 	}`

// 	req, err := http.NewRequest("PUT", baseBooksEndpointURL+bookId, strings.NewReader(updateReqBody))
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v", err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	require.NoError(t, err)
// 	defer res.Body.Close()

// 	assert.Equal(t, http.StatusOK, res.StatusCode)

// 	var response map[string]interface{}
// 	err = json.NewDecoder(res.Body).Decode(&response)
// 	require.NoError(t, err)

// 	updatedBook, ok := response["book"].(map[string]interface{})

// 	assert.True(t, ok)
// 	assert.Equal(t, bookId, updatedBook["id"].(string))
// 	assert.Equal(t, "Updated Book Title", updatedBook["title"].(string))
// 	assert.Equal(t, "Updated Book Author", updatedBook["author"].(string))
// 	assert.Equal(t, float64(2021), updatedBook["published_year"].(float64))
// 	assert.Equal(t, "9780743273565", updatedBook["isbn"].(string))
// 	assert.Contains(t, updatedBook, "updated_at")
// }

// func TestDeleteBookById(t *testing.T) {

// 	req, err := http.NewRequest("DELETE", baseBooksEndpointURL+bookId, nil)
// 	if err != nil {
// 		t.Fatalf("Could not create request: %v", err)
// 	}

// 	res, err := http.DefaultClient.Do(req)
// 	require.NoError(t, err)
// 	defer res.Body.Close()

// 	require.Equal(t, http.StatusOK, res.StatusCode)

// 	var response map[string]interface{}
// 	err = json.NewDecoder(res.Body).Decode(&response)
// 	require.NoError(t, err)

// 	message, ok := response["message"]
// 	assert.True(t, ok)
// 	assert.Equal(t, "book deleted", message)

// }
