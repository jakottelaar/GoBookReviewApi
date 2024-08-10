package user

import (
	"time"
)

type CreateUserRequest struct {
	Username string `json:"username" validate:"required" example:"exampleuser"`
	Email    string `json:"email" validate:"required,email" example:"exampleuser@mail.com"`
	Password string `json:"password" validate:"required" example:"password"`
}

type CreateUserResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
	Username  string    `json:"username" example:"exampleuser"`
	Email     string    `json:"email" example:"exampleuser@mail.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
}

type UpdateUserRequest struct {
	Username string `json:"username" validate:"required" example:"exampleuser"`
	Email    string `json:"email" validate:"required,email" example:"exampleuser@mail.com"`
}

type UpdateUserResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
	Username  string    `json:"username" example:"exampleuser"`
	Email     string    `json:"email" example:"exampleuser@mail.com"`
	UpdatedAt time.Time `json:"updated_at" example:"2024-01-01T00:00:00Z"`
}

type GetUserProfileResponse struct {
	ID        string    `json:"id" example:"123e4567-e89b-12d3-a456-426614174000" format:"uuid"`
	Username  string    `json:"username" example:"exampleuser"`
	Email     string    `json:"email" example:"exampleuser@mail.com"`
	CreatedAt time.Time `json:"created_at" example:"2024-01-01T00:00:00Z"`
}
