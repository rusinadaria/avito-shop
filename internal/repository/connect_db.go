package repository

import (
	"database/sql"
	_"github.com/lib/pq"
	"log"
	"log/slog"
)

func ConnectDatabase(storage_path string, logger *slog.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", storage_path)
	if err != nil {
		log.Fatal("Failed connect database")
		return nil, err
	}
	logger.Info("Connect database")
	return db, nil
}
