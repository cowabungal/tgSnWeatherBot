package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) IsUser(userId int) error {
	var username string

	query := fmt.Sprintf("SELECT username from %s WHERE user_id=$1", usersTable)
	err := r.db.Get(&username, query, userId)

	return err
}

func (r *AuthRepository) CreateUser(username string, userId int) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (username, user_id) values ($1, $2)", usersTable)

	row := r.db.QueryRow(query, username, userId)
	err := row.Scan(&id)
	return err
}
