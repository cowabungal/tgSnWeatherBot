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

func (r *UserRepository) City(userId int) (string, error) {
	var city string

	query := fmt.Sprintf("SELECT city from %s WHERE user_id=$1;", usersTable)
	err := r.db.Get(&city, query, userId)

	return city, err
}

func (r *UserRepository) ChangeCity(userId int, newCity string) (string, error) {
	var city string

	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE user_id=$2 RETURNING city", usersTable, cityColumn)
	err := r.db.Get(&city, query, newCity, userId)

	return city, err
}
