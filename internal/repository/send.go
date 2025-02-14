package repository

import (
	"avito-shop/models"
	"database/sql"
	"log"
	"errors"
)

type TransactionPostgres struct {
	db *sql.DB
}

func NewTransactionPostgres(db *sql.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (t *TransactionPostgres) GetUserById(user models.AuthRequest) (int, error) {
	var id int
	err := t.db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&id)
	if err != nil {
		log.Println("Пользователь с таким именем не существует:", err)
		return 0, err
	}
	return id, nil
}

func (t *TransactionPostgres) SendCoin(fromUserId int, req models.SendCoinRequest) error {
	tx, err := t.db.Begin()
	if err != nil {
		log.Println(err)
		log.Printf("Ошибка при получении монет отправителя: %v", err)
		return err
	}

	var senderCoins int
	err = tx.QueryRow("SELECT coins FROM wallet WHERE employee_id = $1", fromUserId).Scan(&senderCoins)
	if err != nil {
		log.Println(err)
		log.Printf("Ошибка при получении монет отправителя: %v", err)
		tx.Rollback()
		return err
	}

	if senderCoins < req.Amount {
		tx.Rollback()
		log.Printf("Ошибка при получении монет отправителя: %v", err)
		return errors.New("недостаточно монет для отправки")
	}

	_, err = tx.Exec("UPDATE wallet SET coins = coins - $1 WHERE employee_id = $2", req.Amount, fromUserId)
	if err != nil {
		log.Println(err)
		log.Printf("Ошибка при получении монет отправителя: %v", err)
		tx.Rollback()
		return err
	}

	var toUserId int
	err = tx.QueryRow(`SELECT employee_id FROM wallet
						WHERE employee_id = (SELECT id FROM users WHERE username = $1)`, req.ToUser).Scan(&toUserId)
	if err != nil {
		log.Println(err)
		log.Printf("Ошибка при получении ID получателя: %v", err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET coins = coins + $1 WHERE employee_id = $2", req.Amount, toUserId)
	if err != nil {
		log.Println(err)
		log.Printf("Ошибка при получении монет отправителя: %v", err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transaction (from_user, to_user, amount) VALUES ($1, $2, $3)", fromUserId, toUserId, req.Amount)
	if err != nil {
		log.Println(err)
		log.Printf("Ошибка при получении монет отправителя: %v", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()

}
