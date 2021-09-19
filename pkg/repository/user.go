package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Name(userId int) (string, error) {
	var name string

	query := fmt.Sprintf("SELECT name from %s WHERE user_id=$1 ORDER BY random() LIMIT 1;", namesTable)
	err := r.db.Get(&name, query, userId)

	return name, err
}
