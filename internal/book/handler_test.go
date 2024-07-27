package book

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var mockService = new(MockBookService)
var handler = NewBookHandler(mockService)

func TestCreateBookHandler(t *testing.T) {

	t.Run("POST Book handler: Successfully create a book", func(t *testing.T) {
		reqBody := CreateBookRequest{
			Title:         "Test Book",
			Author:        "Test Author",
			PublishedYear: 2004,
			ISBN:          "9780743273565",
		}

		expectedBook := &Book{
			ID:            uuid.New(),
			Title:         reqBody.Title,
			Author:        reqBody.Author,
			PublishedYear: reqBody.PublishedYear,
			ISBN:          reqBody.ISBN,
			CreatedAt:     time.Now(),
		}

		mockService.On("Create", mock.AnythingOfType("*book.CreateBookRequest")).Return(expectedBook, nil)

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/v1/api/books", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		handler.CreateBook(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		book, ok := response["book"].(map[string]interface{})

		assert.True(t, ok)
		assert.Equal(t, expectedBook.ID.String(), book["id"])
		assert.Equal(t, expectedBook.Title, book["title"])
		assert.Equal(t, expectedBook.Author, book["author"])
		assert.Equal(t, expectedBook.ISBN, book["isbn"])
		assert.Contains(t, book, "published_year")
		assert.Contains(t, book, "created_at")

		mockService.AssertExpectations(t)
	})

	t.Run("POST Book handler: Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/v1/api/books", bytes.NewReader([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		handler.CreateBook(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("POST Book handler: Missing title field", func(t *testing.T) {
		reqBody := CreateBookRequest{
			Author:        "Test Author",
			PublishedYear: 2004,
			ISBN:          "9780743273565",
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/v1/api/books", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		handler.CreateBook(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

		var response map[string]interface{}

		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

		require.NoError(t, err)
		assert.Contains(t, response, "error")
		assert.Contains(t, response["error"], "Title")

	})

	t.Run("POST Book handler: Missing multiple required fields", func(t *testing.T) {
		reqBody := CreateBookRequest{
			Title: "Test Book",
		}

		body, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/v1/api/books", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		handler.CreateBook(w, req)

		var response map[string]interface{}

		err := json.Unmarshal(w.Body.Bytes(), &response)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

		assert.NoError(t, err)
		assert.Contains(t, response, "error")
		assert.Contains(t, response["error"], "Author")
		assert.Contains(t, response["error"], "PublishedYear")
		assert.Contains(t, response["error"], "ISBN")
	})
}

func TestGetBookHandler(t *testing.T) {
	mockService := new(MockBookService)
	handler := NewBookHandler(mockService)

	t.Run("GET Book by id handler: Successfully get a book", func(t *testing.T) {
		bookID := uuid.New()

		expectedBook := &Book{
			ID:            bookID,
			Title:         "Test Book",
			Author:        "Test Author",
			PublishedYear: 2004,
			ISBN:          "9780743273565",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		mockService.On("GetBookById", bookID.String()).Return(expectedBook, nil)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/books/"+bookID.String(), nil)

		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Get("/v1/api/books/{id}", handler.GetBookById)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		book, ok := response["book"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, expectedBook.ID.String(), book["id"])
		assert.Equal(t, expectedBook.Title, book["title"])
		assert.Equal(t, expectedBook.Author, book["author"])
		assert.Equal(t, expectedBook.ISBN, book["isbn"])
		assert.Contains(t, book, "published_year")
		assert.Contains(t, book, "created_at")
		assert.Contains(t, book, "updated_at")

		mockService.AssertExpectations(t)
	})

	t.Run("GET Book by id handler: Book not found", func(t *testing.T) {
		bookID := uuid.New()

		mockService.On("GetBookById", bookID.String()).Return((*Book)(nil), common.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/books/"+bookID.String(), nil)
		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Get("/v1/api/books/{id}", handler.GetBookById)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		mockService.AssertExpectations(t)
	})
}

func TestUpdateBookHandler(t *testing.T) {
	mockService := new(MockBookService)
	handler := NewBookHandler(mockService)

	t.Run("PUT Book handler: Successfully update a book", func(t *testing.T) {
		bookID := uuid.New()

		updateReq := UpdateBookRequest{
			Title:         "Updated Book",
			Author:        "Updated Author",
			PublishedYear: 2005,
			ISBN:          "0987654321",
		}

		existingBook := &Book{
			ID:            bookID,
			Title:         "Test Book",
			Author:        "Test Author",
			PublishedYear: 2004,
			ISBN:          "9780743273565",
			CreatedAt:     time.Now(),
		}

		expectedBook := &Book{
			ID:            bookID,
			Title:         updateReq.Title,
			Author:        updateReq.Author,
			PublishedYear: updateReq.PublishedYear,
			ISBN:          updateReq.ISBN,
			CreatedAt:     existingBook.CreatedAt,
		}

		mockService.On("Update", bookID.String(), mock.AnythingOfType("*book.UpdateBookRequest")).Return(expectedBook, nil)

		body, _ := json.Marshal(updateReq)

		req := httptest.NewRequest(http.MethodPut, "/v1/api/books/"+bookID.String(), bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Put("/v1/api/books/{id}", handler.UpdateBook)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		book, ok := response["book"].(map[string]interface{})
		require.True(t, ok)
		assert.Equal(t, expectedBook.ID.String(), book["id"])
		assert.Equal(t, expectedBook.Title, book["title"])
		assert.Equal(t, expectedBook.Author, book["author"])

		mockService.AssertExpectations(t)
	})

	t.Run("PUT Book handler: Empty body request", func(t *testing.T) {
		bookID := uuid.New()

		req := httptest.NewRequest(http.MethodPut, "/v1/api/books/"+bookID.String(), bytes.NewReader([]byte("")))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Put("/v1/api/books/{id}", handler.UpdateBook)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")

		assert.Equal(t, "body must not be empty", response["error"])

		mockService.AssertExpectations(t)

	})
}

func TestDeleteBookHandler(t *testing.T) {

	mockService := new(MockBookService)
	handler := NewBookHandler(mockService)

	t.Run("DELETE handler: Successfully delete a book", func(t *testing.T) {

		bookID := uuid.New()

		mockService.On("Delete", bookID.String()).Return(nil)

		req := httptest.NewRequest(http.MethodDelete, "/v1/api/books/"+bookID.String(), nil)
		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Delete("/v1/api/books/{id}", handler.DeleteBook)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "message")
		assert.Equal(t, "Successfully deleted book", response["message"])

		mockService.AssertExpectations(t)
	})

	t.Run("DELETE handler: Book not found", func(t *testing.T) {
		bookID := uuid.New()

		mockService.On("GetBookById", bookID.String()).Return((*Book)(nil), common.ErrNotFound)

		req := httptest.NewRequest(http.MethodGet, "/v1/api/books/"+bookID.String(), nil)
		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Get("/v1/api/books/{id}", handler.GetBookById)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		mockService.AssertExpectations(t)
	})

	t.Run("DELETE handler: No book with id", func(t *testing.T) {
		bookID := uuid.New()

		mockService.On("Delete", bookID.String()).Return(common.ErrNotFound)

		req := httptest.NewRequest(http.MethodDelete, "/v1/api/books/"+bookID.String(), nil)
		w := httptest.NewRecorder()

		r := chi.NewRouter()
		r.Delete("/v1/api/books/{id}", handler.DeleteBook)
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Contains(t, response, "error")
		mockService.AssertExpectations(t)
	})

}
