package book

import (
	"time"
)

type CreateBookRequest struct {
	Title         string `json:"title" validate:"required" example:"The Great Gatsby"`
	Author        string `json:"author" validate:"required" example:"F. Scott Fitzgerald"`
	PublishedYear int    `json:"published_year" validate:"required,gt=0" example:"1925"`
	ISBN          string `json:"isbn" validate:"required,isbn13" example:"9780743273565"`
}

type UpdateBookRequest struct {
	Title         string `json:"title" validate:"required" example:"The Great Gatsby"`
	Author        string `json:"author" validate:"required" example:"F. Scott Fitzgerald"`
	PublishedYear int    `json:"published_year" validate:"required" example:"1925"`
	ISBN          string `json:"isbn" validate:"required" example:"9780743273565"`
}

type CreateBookResponse struct {
	ID            string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
	Title         string    `json:"title" example:"The Great Gatsby"`
	Author        string    `json:"author" example:"F. Scott Fitzgerald"`
	PublishedYear int       `json:"published_year" example:"1925"`
	ISBN          string    `json:"isbn" example:"9780743273565"`
	CreatedAt     time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UserID        string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
}

type UpdateBookResponse struct {
	ID            string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
	Title         string    `json:"title" example:"The Great Gatsby"`
	Author        string    `json:"author" example:"F. Scott Fitzgerald"`
	PublishedYear int       `json:"published_year" example:"1925"`
	ISBN          string    `json:"isbn" example:"9780743273565"`
	UpdatedAt     time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	UserID        string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
}

type GetBookResponse struct {
	ID            string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
	Title         string    `json:"title" example:"The Great Gatsby"`
	Author        string    `json:"author" example:"F. Scott Fitzgerald"`
	PublishedYear int       `json:"published_year" example:"1925"`
	ISBN          string    `json:"isbn" example:"9780743273565"`
	CreatedAt     time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
	UpdatedAt     time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
	UserID        string    `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
}
