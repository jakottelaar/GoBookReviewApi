package user

import (
	"database/sql"
)

type UserRepository interface {
	FindByEmail(email string) (*User, error)
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

func (ur *userRepository) FindByEmail(email string) (*User, error) {
	return nil, nil
}

func (ur *userRepository) Save(user *User) (*User, error) {
	query :=
		`INSERT INTO users (id, username, email, password)
        VALUES ($1, $2, $3, $4)
        RETURNING id, created_at`

	err := ur.db.QueryRow(query, user.ID, user.Username, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return user, nil

}

func (ur *userRepository) Update(user *User) (*User, error) {
	return nil, nil
}

func (ur *userRepository) Delete(id string) error {
	return nil
}
