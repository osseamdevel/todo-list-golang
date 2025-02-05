package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/osseamdevel/todo-list-golang/models"
)

func GetTodos(db *sql.DB) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		todos, err := models.GetAllTodos(db)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.Header().Set("Content-Type", "application/json")
		json.NewEncoder(response).Encode(todos)
	}
}

func CreateTodo(db *sql.DB) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		var todo struct {
			Task string `json:"task"`
		}
		if err := json.NewDecoder(request.Body).Decode(&todo); err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.CreateTodo(db, todo.Task); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.WriteHeader(http.StatusCreated)
	}
}

func UpdateTodo(db *sql.DB) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		idStr := request.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(response, "Invalid ID", http.StatusBadRequest)
			return
		}

		var todo struct {
			Completed bool `json:"completed"`
		}
		if err := json.NewDecoder(request.Body).Decode(&todo); err != nil {
			http.Error(response, err.Error(), http.StatusBadRequest)
			return
		}

		if err := models.UpdateTodo(db, id, todo.Completed); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.WriteHeader(http.StatusOK)
	}
}

func DeleteTodo(db *sql.DB) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		idStr := request.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(response, "Invalid ID", http.StatusBadRequest)
			return
		}

		if err := models.DeleteTodo(db, id); err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		response.WriteHeader(http.StatusOK)
	}
}
