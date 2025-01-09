package main

import (
	"fmt"
	"net/http"

	"github.com/10Narratives/todo-list/database"
	"github.com/10Narratives/todo-list/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	database.ConnectDB()

	router := chi.NewRouter()

	router.Get("/tasks", handlers.GetTasks)
	router.Post("/tasks", handlers.PostTask)
	router.Get("/tasks/{id}", handlers.GetTaskByID)
	router.Delete("/tasks/{id}", handlers.DeleteTask)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Привет, программист!"))
	})

	if err := http.ListenAndServe(":3000", router); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}
