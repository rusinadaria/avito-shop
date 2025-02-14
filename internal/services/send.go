package services

import (
	"avito-shop/internal/repository"
	"avito-shop/models"
	"fmt"
	"log"
	"strconv"
)

type TransactionService struct {
	repo repository.Transaction
}

func NewTransactionService(repo repository.Transaction) *TransactionService {
	return &TransactionService{repo: repo}
}

func (t *TransactionService) SendCoin(senderId string, username string, amount int) error {
	fmt.Println("Получатель: ",username)

	recipient := models.AuthRequest {
		Username: username,
	}
	_, err := t.repo.GetUserById(recipient)

	fmt.Println("Отправитель: ", senderId)
	if err != nil {
		log.Println("Невозможно найти пользователя с таким именем")
		return err
	}

	req := models.SendCoinRequest {
		ToUser: username,
		Amount: amount,
	}

	numId, err := strconv.Atoi(senderId)
    if err != nil {
        fmt.Println("Ошибка конвертации:", err)
        return err
    }
	err = t.repo.SendCoin(numId, req)
	if err != nil {
		log.Println("Ошибка в сервисе, функции SendCoin")
		log.Printf("Ошибка в сервисе, функции SendCoin: %v", err)
	}

	return nil
}