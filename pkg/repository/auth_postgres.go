package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mohito22/todo-app-golang/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}
