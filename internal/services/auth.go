package services

import (
	"avito-shop/internal/repository"
	"avito-shop/models"
	// "errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"github.com/golang-jwt/jwt"
	// "fmt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) GetUser(username string, password string) (int, error) {
	user := models.AuthRequest {
		Username: username,
		Password: password,
	}
	id, err := s.repo.FindUser(user)
	if err != nil {
		log.Println("Невозможно найти пользователя с таким именем")
		return 0, err
	}

	if !checkPasswordHash(password, user.Password) {
		log.Println("Неверный пароль для пользователя:", username)
		return 0, err
	}
	return id, nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) CreateUser(username string, password string) (int, error) {
	password, err := hashPassword(password)
	if err != nil {
		log.Println("Не удалось захэшировать пароль")
		return 0, err 
	}
	user := models.AuthRequest {
		Username: username,
		Password: password,
	}

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	user_wallet := models.Wallet {
		Employee_id: id,
	}
	err = s.repo.AddCoins(user_wallet)
	if err != nil {
		log.Println("Ошибка при зачислении коинов")
		return 0, nil
	}

	return id, nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка при хешировании пароля:", err)
		return "", err
	}
	return string(hash), nil
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}


func (s *AuthService) GenerateToken(username string, password string) (string, error) {
	user := models.AuthRequest {
		Username: username,
		Password: password,
	}
	id, err := s.repo.FindUser(user)
	if err != nil {
		log.Println("Пользователя с таким именем не существует")
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
	})
	signingKey := "qrkjk#4#%35FSFJlja#4353KSFjH"

	return token.SignedString([]byte(signingKey))
}