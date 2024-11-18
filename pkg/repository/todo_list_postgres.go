package repository

import (
	"fmt"
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func (r *TodoListPostgres) GetById(userIdd int, listId int) (todoapp.TodoList, error) {
	var list todoapp.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1 AND tl.id = $2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userIdd, listId)
	return list, err
}

func (r *TodoListPostgres) GetAll(id int) ([]todoapp.TodoList, error) {
	var lists []todoapp.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul ON tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, id)
	return lists, err
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}
func (r *TodoListPostgres) Create(userId int, list todoapp.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) values ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createListQuery = fmt.Sprintf("INSERT INTO %s (user_id, list_id) values ($1, $2)", usersListsTable)
	_, err = tx.Exec(createListQuery, userId, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
