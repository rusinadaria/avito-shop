package tests

import (
	"net/http/httptest"
	"net/http"
	"encoding/json"
	"avito-shop/models"
	// "avito-shop/internal/repository"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func (s *APITestSuite) TestLoginAndGetToken() {
	loginRequest := models.AuthRequest{
		Username: "sender_username",
		Password: "12345",
	}

	requestBody, err := json.Marshal(loginRequest)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest("POST", "/api/auth", bytes.NewBuffer(requestBody))
	assert.NoError(s.T(), err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.AddUserHandler)

	handler.ServeHTTP(rr, req)

	assert.Equal(s.T(), http.StatusOK, rr.Code)

	var response struct {
		Token string `json:"token"`
	}
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(s.T(), err)

	// return response.Token
	s.token = response.Token
}

func (s *APITestSuite) createTestUsers() {
	sender := models.AuthRequest{
		Username: "sender_username",
		Password: "12345",
	}
	// Создание пользователя в базе
	_, err := s.repos.CreateUser(sender)
	assert.NoError(s.T(), err)

	recipient := models.AuthRequest{
		Username: "recipient_username",
		Password: "12345",
	}
	_, err = s.repos.CreateUser(recipient)
	assert.NoError(s.T(), err)
}

func (s *APITestSuite) TestSendCoins() {
	sendCoinRequest := models.SendCoinRequest{
		ToUser: "recipient_username",
		Amount: 100,
	}

	requestBody, err := json.Marshal(sendCoinRequest)
	assert.NoError(s.T(), err)

	req, err := http.NewRequest("POST", "/api/sendCoin", bytes.NewBuffer(requestBody))
	assert.NoError(s.T(), err)

	req.AddCookie(&http.Cookie{Name: "auth_token", Value: s.token})


	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.handler.SendHandler)


	handler.ServeHTTP(rr, req)


	assert.Equal(s.T(), http.StatusOK, rr.Code)

	s.checkCoinBalances("sender_username", "recipient_username", 100)
}

func (s *APITestSuite) checkCoinBalances(senderUsername, recipientUsername string, amount int) {
	senderBalance, err := s.getUserBalance(senderUsername)
	assert.NoError(s.T(), err, "Failed to get sender balance")

	recipientBalance, err := s.getUserBalance(recipientUsername)
	assert.NoError(s.T(), err, "Failed to get recipient balance")

	assert.Equal(s.T(), 1000-amount, senderBalance, "Sender balance is incorrect")
	assert.Equal(s.T(), 1000+amount, recipientBalance, "Recipient balance is incorrect")
}

func (s *APITestSuite) getUserBalance(username string) (int, error) {
	var balance int
	query := `SELECT coins FROM wallet WHERE user_id = (SELECT id FROM "user" WHERE username = $1)`
	err := s.db.QueryRow(query, username).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}