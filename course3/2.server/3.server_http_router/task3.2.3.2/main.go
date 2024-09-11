package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	// Первая группа маршрутов
	r.Route("/group1", func(r chi.Router) {
		r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 1 Привет, мир 1"))
		})
		r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 1 Привет, мир 2"))
		})
		r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 1 Привет, мир 3"))
		})
	})

	// Вторая группа маршрутов
	r.Route("/group2", func(r chi.Router) {
		r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 2 Привет, мир 1"))
		})
		r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 2 Привет, мир 2"))
		})
		r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 2 Привет, мир 3"))
		})
	})

	// Третья группа маршрутов
	r.Route("/group3", func(r chi.Router) {
		r.Get("/1", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 3 Привет, мир 1"))
		})
		r.Get("/2", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 3 Привет, мир 2"))
		})
		r.Get("/3", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Group 3 Привет, мир 3"))
		})
	})

	http.ListenAndServe(":8080", r)
}
