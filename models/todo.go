package models

import (
	"database/sql"
	"log"
)

type Todo struct {
	ID        int    `json:"id"`
	TASK      string `json:"task"`
	Completed bool   `json:"Completed"`
}

func GetAllTodos(db *sql.DB) ([]Todo, error) {
	rows, err := db.Query("SELECT id, task, completed FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.TASK, &t.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func CreateTodo(db *sql.DB, task string) error {
	_, err := db.Exec("INSERT INTO todos (task, completed) VALUES (?, ?)", task, false)
	return err
}

func UpdateTodo(db *sql.DB, id int, completed bool) error {
	_, err := db.Exec("UPDATE todos SET completed = ? WHERE id = ?", completed, id)
	return err
}

func DeleteTodo(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
