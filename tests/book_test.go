//go:build integration
// +build integration

package tests

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var bookId string

func TestCreateBookRequest(t *testing.T) {
	reqBody := `{
		"title": "Book Title",
		"author": "Book Author",
		"published_year": 2020,
		"isbn": "9780743273565"
	}`

	req, err := http.NewRequest("POST", baseBooksEndpointUrl, strings.NewReader(reqBody))
	require.NoError(t, err)

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
}

func TestGetBookByIdRequest(t *testing.T) {

	req, err := http.NewRequest("GET", baseBooksEndpointUrl+bookId, nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

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
}

func TestUpdateBookById(t *testing.T) {
	updateReqBody := `{
		"title": "Updated Book Title",
		"author": "Updated Book Author",
		"published_year": 2021,
		"isbn": "9780743273565"
	}`

	req, err := http.NewRequest("PUT", baseBooksEndpointUrl+bookId, strings.NewReader(updateReqBody))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	updatedBook, ok := response["book"].(map[string]interface{})

	assert.True(t, ok)
	assert.Equal(t, bookId, updatedBook["id"].(string))
	assert.Equal(t, "Updated Book Title", updatedBook["title"].(string))
	assert.Equal(t, "Updated Book Author", updatedBook["author"].(string))
	assert.Equal(t, float64(2021), updatedBook["published_year"].(float64))
	assert.Equal(t, "9780743273565", updatedBook["isbn"].(string))
}

func TestDeleteBookById(t *testing.T) {

	req, err := http.NewRequest("DELETE", baseBooksEndpointUrl+bookId, nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	require.NoError(t, err)
	defer res.Body.Close()

	require.Equal(t, http.StatusOK, res.StatusCode)

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	require.NoError(t, err)

	message, ok := response["message"]
	assert.True(t, ok)
	assert.Equal(t, "Successfully deleted book", message)

}
