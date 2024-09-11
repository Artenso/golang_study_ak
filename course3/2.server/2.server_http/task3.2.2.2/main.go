package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Ошибка загрузки файла .env: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Переменная PORT не установлена")
	}

	http.HandleFunc("/", handler)

	log.Printf("Сервер запущен на порту %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
