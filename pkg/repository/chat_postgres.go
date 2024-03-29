package repository

import (
	"fmt"
	_"github.com/jmoiron/sqlx"
	"github.com/VladMak/ServiceAPI"
)

func (r *AuthPostgres) SendMessage(user ServiceAPI.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", messagesTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *AuthPostgres) SendMessage1(user ServiceAPI.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}