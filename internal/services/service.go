package services

import (
	"avito-shop/internal/repository"
)

type Auth interface {
	CreateUser(username string, password string) (int, error)
	GetUser(username string, password string) (int, error)
	GenerateToken(username string, password string) (string, error)
}

type Transaction interface {
	SendCoin(senderId string, username string, amount int) error
}

type Service struct {
	Auth
	Transaction
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Authorization),
		Transaction: NewTransactionService(repos.Transaction),
	}
}
