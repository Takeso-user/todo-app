package repository

import (
	"fmt"
	todoapp "github.com/Takeso-user/todo-app"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func (r *TodoItemPostgres) Update(userId int, itemId int, input todoapp.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}
	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++

	}
	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, *input.Done)
		argId++

	}
	setQuery := strings.Join(setValues, ", ")
	query := fmt.Sprintf(`UPDATE %s ti SET %s FROM %s li, %s ul	 WHERE 
                     ti.id = li.item_id AND 
                     li.list_id = ul.list_id AND 
                     ul.user_id = $%d AND 
                     ti.id = $%d`,
		todoItemsTable, setQuery, listsItemsTable, usersListsTable, argId, argId+1)
	args = append(args, userId, itemId)
	logrus.Infof("query: %s", query)
	logrus.Infof("args: %s", args)

	res, err := r.db.Exec(query, args...)
	logrus.Infof("res: %s", res)
	return err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`DELETE FROM %s ti USING %s li, %s ul 
       WHERE ti.id = li.item_id AND li.list_id = ul.list_id AND ul.user_id = $1 AND ti.id = $2`,
		todoItemsTable, listsItemsTable, usersListsTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *TodoItemPostgres) GetById(userId, itemId int) (todoapp.TodoItem, error) {
	var items todoapp.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
         		INNER JOIN %s li ON ti.id = li.item_id  
             	inner join %s ul on li.list_id = ul.list_id 
                WHERE ul.user_id = $2 AND ti.id = $1`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Get(&items, query, itemId, userId); err != nil {
		return items, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetAll(userId int, listId int) ([]todoapp.TodoItem, error) {
	var items []todoapp.TodoItem
	query := fmt.Sprintf(`SELECT ti.id, ti.title, ti.description, ti.done FROM %s ti 
         INNER JOIN %s li ON ti.id = li.item_id  
             inner join %s ul on li.list_id = ul.list_id 
                          WHERE ul.user_id = $2 AND ul.list_id = $1`,
		todoItemsTable, listsItemsTable, usersListsTable)

	if err := r.db.Select(&items, query, listId, userId); err != nil {
		return nil, err
	}
	return items, nil
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item todoapp.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoItemsTable)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	err = row.Scan(&itemId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	createListItemsQuery := fmt.Sprintf("INSERT INTO %s (list_id, item_id) VALUES ($1, $2)", listsItemsTable)
	_, err = tx.Exec(createListItemsQuery, listId, itemId)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			return 0, err
		}
	}
	return itemId, tx.Commit()

}
