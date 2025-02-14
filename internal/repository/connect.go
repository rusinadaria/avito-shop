package repository

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
	"log/slog"
)

func ConnectDatabase(logger *slog.Logger) (*sql.DB, error) {
	connStr := "user=postgres password=root dbname=shop sslmode=disable"
	// connStr := os.Getenv("user=postgres password=root dbname=shop sslmode=disable")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed connect database")
		return nil, err
	}
	logger.Info("Connect database")
	return db, nil
}
