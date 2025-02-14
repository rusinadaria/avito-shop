package main

import (
	// "fmt"
	"log/slog"
	"net/http"
	"os"
	"log"
	// "database/sql"
	_"github.com/lib/pq"
		// "encoding/json"
			"avito-shop/internal/handlers"
	"avito-shop/internal/services"
		"avito-shop/internal/repository"

)

func main() {
	logger := configLogger()

	db, err := repository.ConnectDatabase(logger)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	repo := repository.NewRepository(db)
	srv := services.NewService(repo)
	handler := handlers.NewHandler(srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	err = http.ListenAndServe(port, handler.InitRoutes())
	if err != nil {
		log.Fatal("Не удалось запустить сервер:", err)
	}
}

func configLogger() *slog.Logger {
	var logger *slog.Logger

	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
        slog.Error("Unable to open a file for writing")
    }

	opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }

	logger = slog.New(slog.NewJSONHandler(f, opts))
	return logger
}



