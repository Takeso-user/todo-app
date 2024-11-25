package repository

import (
	"fmt"
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
func (r *AuthPostgres) CreateUser(user todoapp.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	var id int
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (todoapp.User, error) {
	var user todoapp.User
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
