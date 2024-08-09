package book

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type BookRepository interface {
	FindById(id string) (*Book, error)
	Save(book *Book) (*Book, error)
	Update(book *Book) (*Book, error)
	Delete(id string) error
}

type bookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) BookRepository {
	return &bookRepository{
		db: db,
	}
}

type Book struct {
	ID            uuid.UUID
	Title         string
	Author        string
	PublishedYear int
	ISBN          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
	UserId        uuid.UUID
}

func (r *bookRepository) Save(book *Book) (*Book, error) {
	query := `
		INSERT INTO books (id, title, author, published_year, isbn, user_id) 
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at`

	err := r.db.QueryRow(query, book.ID, book.Title, book.Author, book.PublishedYear, book.ISBN, book.UserId).Scan(&book.ID, &book.CreatedAt)

	if err != nil {
		return nil, err
	}

	return book, nil

}

func (r *bookRepository) FindById(id string) (*Book, error) {
	query := `
		SELECT id, title, author, published_year, isbn, created_at, updated_at, user_id
		FROM books
		WHERE id = $1`

	var book Book

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, id).Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.ISBN, &book.CreatedAt, &book.UpdatedAt, &book.UserId)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, common.ErrNotFound
		default:
			return nil, err
		}
	}

	return &book, nil
}

func (r *bookRepository) Update(book *Book) (*Book, error) {
	query := `
		UPDATE books
		SET title = $1, author = $2, published_year = $3, isbn = $4
		WHERE id = $5
		RETURNING updated_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, book.Title, book.Author, book.PublishedYear, book.ISBN, book.ID).Scan(&book.UpdatedAt)

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

func (r *bookRepository) Delete(id string) error {
	query := `
		DELETE FROM books 
		WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return common.ErrNotFound
	}

	return nil
}
