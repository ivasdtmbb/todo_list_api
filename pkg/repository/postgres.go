package repository

// sudo docker run --name=todo_db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres
// sudo docker exec -it todo_db /bin/bash
// psgl -U postgres
//  migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up


import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
	todoListsTable = "todo_lists"
	usersListsTable = "users_lists"
	todoItemsTable = "todo_items"
	listsItemsTable = "lists_items"
)

type Config struct {
	Host 		string
	Port 		string
	Username	string
	Password	string
	DBName 		string
	SSLMode 	string
}

func NewPostgressDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}