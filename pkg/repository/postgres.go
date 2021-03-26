package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable  = "todo_items"
	listsItemsTable = "lists_items"
)

func NewPostgressDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	logrus.Info("db connecting ...")
	if err != nil {
		logrus.Info("db connecting failed")
		return nil, err
	}

	logrus.Info("db connecting Done!")

	logrus.Info("Pinging ...")
	if err = db.Ping(); err != nil {
		logrus.Info("Pinging failed")
		return nil, err
	}

	logrus.Info("Pinging done!")

	return db, nil
}
