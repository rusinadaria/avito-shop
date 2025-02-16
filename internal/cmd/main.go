package main

import (
	"log/slog"
	"net/http"
	"os"
	"log"
	_"github.com/lib/pq"
	"avito-shop/internal/handlers"
	"avito-shop/internal/services"
	"avito-shop/internal/repository"
	"fmt"
	// "github.com/joho/godotenv"
)

func main() {
	logger := configLogger()

	// err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Ошибка загрузки .env файла")
    // }

	storage_path := os.Getenv("DB_PATH")
	// fmt.Println("DB_PATH:", storage_path)
	db, err := repository.ConnectDatabase(storage_path, logger)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	repo := repository.NewRepository(db)
	srv := services.NewService(repo)
	handler := handlers.NewHandler(srv)

	port := os.Getenv("PORT")
	fmt.Println("PORT:", port)
	if port == "" {
		port = ":8080"
	}

	err = http.ListenAndServe(port, handler.InitRoutes(logger))
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



