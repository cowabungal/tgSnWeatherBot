package repository

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"tgSnWeatherBot"
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

func (r *UserRepository) AddName(userId int, name string) (string, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, user_id) values ($1, $2) RETURNING name", namesTable)

	var nameAdded string

	row := r.db.QueryRow(query, name, userId)
	err := row.Scan(&nameAdded)

	return nameAdded, err
}

func (r *UserRepository) DeleteName(userId int, name string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1 and name=$2", namesTable)

	_, err := r.db.Query(query, userId, name)

	return err
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

func (r *UserRepository) State(userId int) (string, error) {
	var state string

	query := fmt.Sprintf("SELECT state from %s WHERE user_id=$1;", usersTable)
	err := r.db.Get(&state, query, userId)

	return state, err
}

func (r *UserRepository) ChangeState(userId int, newState string) (string, error) {
	var state string

	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE user_id=$2 RETURNING state;", usersTable, stateColumn)
	err := r.db.Get(&state, query, newState, userId)

	return state, err
}


func (r *UserRepository) Info (userId int) (*tgSnWeatherBot.User, error) {
	var list tgSnWeatherBot.User

	query := fmt.Sprintf("SELECT username, user_id, city FROM %s WHERE user_id=$1",
		usersTable)
	err := r.db.Get(&list, query, userId)
	if err != nil {
		return nil, errors.New("UserRepository: Info: Get main info: " + err.Error())
	}

	query = fmt.Sprintf("SELECT name FROM %s WHERE user_id=$1",
		namesTable)
	err = r.db.Select(&list.Names, query, userId)

	if err != nil {
		return nil, errors.New("UserRepository: Info: Select name: " + err.Error())
	}

	return &list, nil
}

func (r *UserRepository) AddCallbackId(userId int, callbackId string) error {
	query := fmt.Sprintf("INSERT INTO %s (user_id, callback_id) values ($1, $2) RETURNING user_id", callbacksTable)

	var tmp string

	row := r.db.QueryRow(query, userId, callbackId)
	err := row.Scan(&tmp)

	return err
}

func (r *UserRepository) AddCallbackData(callbackId, callbackData string) error {
	var tmp string

	query := fmt.Sprintf("UPDATE %s SET %s=$1 WHERE callback_id=$2 RETURNING callback_data;", callbacksTable, callbackDataColumn)
	err := r.db.Get(&tmp, query, callbackData, callbackId)

	return err
}

func (r *UserRepository) GetCallbackData(userId int) (string, error) {
	var data string

	query := fmt.Sprintf("SELECT callback_data from %s WHERE user_id=$1;", callbacksTable)
	err := r.db.Get(&data, query, userId)

	return data, err
}

func (r *UserRepository) GetCallbackId(userId int) (int, error) {
	var id int

	query := fmt.Sprintf("SELECT callback_id from %s WHERE user_id=$1;", callbacksTable)
	err := r.db.Get(&id, query, userId)

	return id, err
}

func (r *UserRepository) DeleteCallback(userId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id=$1", callbacksTable)

	_, err := r.db.Query(query, userId)

	return err
}
