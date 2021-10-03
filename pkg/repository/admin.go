package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"tgSnWeatherBot"
)

type AdminRepository struct {
	db *sqlx.DB
}

func NewAdminRepository(db *sqlx.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (r *AdminRepository) UsersList() ([]tgSnWeatherBot.User, error) {
	var list []tgSnWeatherBot.User

	query := fmt.Sprintf("SELECT users.username, users.user_id, users.city FROM %s",
		usersTable)
	err := r.db.Select(&list, query)

	return list, err
}
