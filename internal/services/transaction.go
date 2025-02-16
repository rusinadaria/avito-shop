package services

import (
	"avito-shop/internal/repository"
	"avito-shop/models"
	"fmt"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (t *TransactionService) SendCoin(senderId int, username string, amount int) error {
	recipient := models.AuthRequest {
		Username: username,
	}
	_, err := t.repo.GetUserById(recipient)
	if err != nil {
		return err
	}

	req := models.SendCoinRequest {
		ToUser: username,
		Amount: amount,
	}

	err = t.repo.SendCoin(senderId, req)
	if err != nil {
		return err
	}

	return nil
}

func (t *TransactionService) BuyItem(userId int, name string) error {
	fmt.Println(userId, name)
	err := t.repo.BuyItem(userId, name)
	if err != nil {
		return err
	}
	return nil
}