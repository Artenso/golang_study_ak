// Пример кода для создания сервера на go-chi
package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/1", handleRoute1)
	r.Get("/2", handleRoute2)
	r.Get("/3", handleRoute3)

	http.ListenAndServe(":8080", r)
}

func handleRoute1(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello world")
}

func handleRoute2(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello world 2")
}

func handleRoute3(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello world 3")
}
