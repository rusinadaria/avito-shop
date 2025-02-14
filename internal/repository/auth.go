package repository

import (
	// "github.com/jmoiron/sqlx"
	"avito-shop/models"
	"database/sql"
	"log"
	// "errors"
	// "fmt"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) FindUser(user models.AuthRequest) (int, error) {
	var id int
	err := r.db.QueryRow("SELECT id FROM users WHERE username = $1", user.Username).Scan(&id)
	if err != nil {
		log.Println("Пользователь с таким именем не существует:", err)
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) CreateUser(user models.AuthRequest) (int, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`
	var id int
	err := r.db.QueryRow(query, user.Username, user.Password).Scan(&id)
	if err != nil {
		log.Println("Ошибка при добавлении пользователя в базу данных:", err)
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) AddCoins(user_wallet models.Wallet) error {
	coins := 1000
	query := `INSERT INTO wallet (employee_id, coins) VALUES ($1, $2)`
	_, err := r.db.Exec(query, user_wallet.Employee_id, coins)
	if err != nil {
		log.Println("Ошибка при начислении коинов:", err)
		return err
	}
	return nil
}