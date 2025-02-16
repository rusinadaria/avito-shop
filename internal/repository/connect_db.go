package repository

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
	"log/slog"
	"fmt"
)

func ConnectDatabase(storage_path string, logger *slog.Logger) (*sql.DB, error) {
	connStr := "user=postgres password=root dbname=shop sslmode=disable"
	// connStr := os.Getenv("user=postgres password=root dbname=shop sslmode=disable")
	fmt.Println("DB_PATH:", storage_path)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed connect database")
		return nil, err
	}
	logger.Info("Connect database")
	return db, nil
}
