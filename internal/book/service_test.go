//go:build unit
// +build unit

package book

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateBookService(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("Create book service: Successfully create a book", func(t *testing.T) {
		createReq := &CreateBookRequest{
			Title:         "Test Book",
			Author:        "Test Author",
			PublishedYear: 2004,
			ISBN:          "6940-550830956-8450",
		}

		expectedBook := &Book{
			ID:            uuid.New(),
			Title:         createReq.Title,
			Author:        createReq.Author,
			PublishedYear: createReq.PublishedYear,
			ISBN:          createReq.ISBN,
			CreatedAt:     time.Now(),
		}

		mockRepo.On("Save", mock.AnythingOfType("*book.Book")).Return(expectedBook, nil)

		result, err := service.Create(createReq)

		require.NoError(t, err)
		assert.Equal(t, expectedBook, result)
		mockRepo.AssertExpectations(t)

	})
}

func TestGetBookByIdService(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("Get book by id service: Successfully get a book", func(t *testing.T) {
		bookID := uuid.New()

		expectedBook := &Book{
			ID:            bookID,
			Title:         "Test Book",
			Author:        "Test Author",
			PublishedYear: 2004,
			ISBN:          "6940-550830956-8450",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}

		mockRepo.On("FindById", bookID.String()).Return(expectedBook, nil)

		result, err := service.GetBookById(bookID.String())

		require.NoError(t, err)

		assert.Equal(t, expectedBook, result)
		assert.Equal(t, expectedBook.ID, result.ID)
		assert.Equal(t, expectedBook.Title, result.Title)
		assert.Equal(t, expectedBook.Author, result.Author)
		assert.Equal(t, expectedBook.PublishedYear, result.PublishedYear)
		assert.Equal(t, expectedBook.ISBN, result.ISBN)
		assert.Equal(t, expectedBook.CreatedAt, result.CreatedAt)
		assert.Equal(t, expectedBook.UpdatedAt, result.UpdatedAt)

		mockRepo.AssertExpectations(t)

	})

	t.Run("Get book by id service: Book not found", func(t *testing.T) {
		bookID := uuid.New()

		mockRepo.On("FindById", bookID.String()).Return((*Book)(nil), common.ErrNotFound)

		result, err := service.GetBookById(bookID.String())

		require.Error(t, err)
		require.Equal(t, common.ErrNotFound, err)

		assert.Nil(t, result)

		mockRepo.AssertExpectations(t)

	})

}

func TestUpdateBookService(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("Update book service: Successfully update a book", func(t *testing.T) {
		bookID := uuid.New()

		updatedBook := &UpdateBookRequest{
			Title:         "Updated Book",
			Author:        "Updated Author",
			PublishedYear: 2005,
			ISBN:          "6940-550830956-8450",
		}

		existingBook := &Book{
			ID:            bookID,
			Title:         "Original Title",
			Author:        "Original Author",
			PublishedYear: 2000,
			ISBN:          "1234567890",
			CreatedAt:     time.Now(),
		}

		expectedBook := &Book{
			ID:            bookID,
			Title:         "Updated Book",
			Author:        "Updated Author",
			PublishedYear: 2005,
			ISBN:          "6940-550830956-8450",
			CreatedAt:     time.Now(),
		}

		// Set up the mock expectations
		mockRepo.On("FindById", bookID.String()).Return(existingBook, nil)
		mockRepo.On("Update", mock.AnythingOfType("*book.Book")).Return(expectedBook, nil)

		// Invoke the service method
		result, err := service.Update(bookID.String(), updatedBook)

		// Verify results
		require.NoError(t, err)
		assert.Equal(t, expectedBook.Title, result.Title)
		mockRepo.AssertExpectations(t)

	})

	t.Run("Update book service: Book not found", func(t *testing.T) {
		bookID := uuid.New()

		updatedBook := &UpdateBookRequest{
			Title:         "Updated Book",
			Author:        "Updated Author",
			PublishedYear: 2005,
			ISBN:          "6940-550830956-8450",
		}

		mockRepo.On("FindById", bookID.String()).Return((*Book)(nil), common.ErrNotFound)

		book, err := service.Update(bookID.String(), updatedBook)

		require.Error(t, err)
		assert.Equal(t, common.ErrNotFound, err)
		assert.Nil(t, book)

		mockRepo.AssertExpectations(t)
	})
}

func TestDeleteBookService(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := NewBookService(mockRepo)

	t.Run("Delete book service: Successfully delete a book", func(t *testing.T) {

		bookID := uuid.New()

		existingBook := &Book{
			ID:            bookID,
			Title:         "Original Title",
			Author:        "Original Author",
			PublishedYear: 2000,
			ISBN:          "1234567890",
			CreatedAt:     time.Now(),
		}

		mockRepo.On("FindById", bookID.String()).Return(existingBook, nil)
		mockRepo.On("Delete", bookID.String()).Return(nil)

		err := service.Delete(bookID.String())

		require.NoError(t, err)

		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete book service: Book not found", func(t *testing.T) {
		bookID := uuid.New()

		mockRepo.On("FindById", bookID.String()).Return((*Book)(nil), common.ErrNotFound)

		err := service.Delete(bookID.String())

		require.Error(t, err)
		assert.Equal(t, common.ErrNotFound, err)

		mockRepo.AssertExpectations(t)
	})
}
