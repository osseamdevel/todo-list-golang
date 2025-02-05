package main

import (
	//"database/sql"
	"log"
	"net/http"

	"github.com/osseamdevel/todo-list-golang/database"
	"github.com/osseamdevel/todo-list-golang/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	http.HandleFunc("/todos", handlers.GetTodos(database.DB))
	http.HandleFunc("/todos/create", handlers.CreateTodo(database.DB))
	http.HandleFunc("/todos/update", handlers.UpdateTodo(database.DB))
	http.HandleFunc("/todos/delete", handlers.DeleteTodo(database.DB))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
