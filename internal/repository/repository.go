package repository

import (
	// "github.com/jmoiron/sqlx"
	"database/sql"
	"avito-shop/models"
	// "errors"
	// "fmt"
	// "log"
)

type Authorization interface {
	FindUser(user models.AuthRequest) (int, error)
	CreateUser(user models.AuthRequest) (int, error)
	AddCoins(user_wallet models.Wallet) error
}

type Transaction interface {
	GetUserById(user models.AuthRequest) (int, error)
	SendCoin(id int, req models.SendCoinRequest) error
}


type Repository struct {
	Authorization
	Transaction
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transaction: NewTransactionPostgres(db),
	}
}