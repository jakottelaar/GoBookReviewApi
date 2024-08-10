package book

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type BookService interface {
	GetBookById(id string) (*Book, error)
	Create(userId string, book *CreateBookRequest) (*Book, error)
	Update(id string, userId string, book *UpdateBookRequest) (*Book, error)
	Delete(id string, userId string) error
}

type bookService struct {
	repo BookRepository
}

func NewBookService(repo BookRepository) BookService {
	return &bookService{
		repo: repo,
	}
}

func (s *bookService) Create(userId string, book *CreateBookRequest) (*Book, error) {

	newId := uuid.New()

	newBook := &Book{
		ID:            newId,
		Title:         book.Title,
		Author:        book.Author,
		PublishedYear: book.PublishedYear,
		ISBN:          book.ISBN,
		UserId:        uuid.MustParse(userId),
	}

	savedBook, err := s.repo.Save(newBook)

	if err != nil {
		return nil, err
	}

	return savedBook, nil
}

func (s *bookService) GetBookById(id string) (*Book, error) {

	book, err := s.repo.FindById(id)

	if err != nil {
		switch {
		case errors.Is(err, common.ErrNotFound):
			return nil, common.ErrNotFound

		default:
			return nil, err
		}
	}

	return book, nil

}

func (s *bookService) Update(id string, userId string, updateReq *UpdateBookRequest) (*Book, error) {

	_, err := s.repo.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrNotFound):
			return nil, common.ErrNotFound
		default:
			return nil, err
		}
	}

	updatedBook := &Book{
		ID:            uuid.MustParse(id),
		Title:         updateReq.Title,
		Author:        updateReq.Author,
		PublishedYear: updateReq.PublishedYear,
		ISBN:          updateReq.ISBN,
		UserId:        uuid.MustParse(userId),
	}

	book, err := s.repo.Update(updatedBook)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, common.ErrNotFound
		default:
			return nil, err
		}
	}

	return book, nil

}

func (s *bookService) Delete(id string, userId string) error {

	_, err := s.repo.FindById(id)
	if err != nil {
		switch {
		case errors.Is(err, common.ErrNotFound):
			return common.ErrNotFound
		default:
			return err

		}

	}

	err = s.repo.Delete(id, userId)

	if err != nil {
		return err
	}

	return nil

}
