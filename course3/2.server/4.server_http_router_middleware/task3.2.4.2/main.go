package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	var err error
	logger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	defer logger.Sync()

	r := chi.NewRouter()

	// Применение middleware для логирования с помощью zap logger
	r.Use(LoggerMiddleware)

	// Здесь можно добавить ваши маршруты с различными методами
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

	http.ListenAndServe(":8080", r)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Request received",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("client_ip", r.RemoteAddr),
		)

		// Продолжение выполнения следующего обработчика
		next.ServeHTTP(w, r)
	})
}
