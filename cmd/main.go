package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Привет, программист!"))
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		fmt.Printf("Start server error: %s", err.Error())
		return
	}
}
