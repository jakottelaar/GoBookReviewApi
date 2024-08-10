package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jakottelaar/gobookreviewapp/pkg/common"
)

type UserRepository interface {
	FindByEmail(email string) (*User, error)
	FindById(id string) (*User, error)
	Save(user *User) (*User, error)
	Update(user *User) (*User, error)
	Delete(id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (r *userRepository) FindByEmail(email string) (*User, error) {

	var user User

	query := `
	SELECT id, password FROM users
	WHERE email = $1`

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Password)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindById(id string) (*User, error) {
	var user User

	query := `
	SELECT id, username, email, created_at FROM users
	WHERE id = $1`

	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Save(user *User) (*User, error) {
	query :=
		`INSERT INTO users (id, username, email, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, user.ID, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		switch {
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_email_key\"":
			return nil, common.ErrEmailAlreadyExists
		case err.Error() == "pq: duplicate key value violates unique constraint \"users_username_key\"":
			return nil, common.ErrUsernameAlreadyExists
		default:
			return nil, err
		}
	}

	return user, nil

}

func (r *userRepository) Update(user *User) (*User, error) {

	query := `
		UPDATE users 
		SET username = $1, email = $2
		WHERE id = $3
		RETURNING updated_at`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, common.ErrNotFound
		default:
			return nil, err
		}
	}

	return user, nil

}

func (r *userRepository) Delete(id string) error {

	query := `
		DELETE FROM users 
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
