package user

import (
	"database/sql"
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

	err := r.db.QueryRow(query, user.ID, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *userRepository) Update(user *User) (*User, error) {
	return nil, nil
}

func (r *userRepository) Delete(id string) error {
	return nil
}
