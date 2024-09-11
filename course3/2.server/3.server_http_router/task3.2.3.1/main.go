package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.HandleFunc("/1", hw1Handler)
	r.HandleFunc("/2", hw2Handler)
	r.HandleFunc("/3", hw3Handler)

	r.HandleFunc("/", othersHandler)

	http.ListenAndServe(":8080", r)
}

func hw1Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello world")
}

func hw2Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello world 2")
}

func hw3Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Hello world 3")
}

func othersHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintf(w, "404 Not Found")
}
