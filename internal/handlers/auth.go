package handlers

import (
	"avito-shop/models"
	"encoding/json"
	// "fmt"
	"net/http"
	"log"
	"time"
)

func setTokenCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,     
		Secure:   true,
	}

	http.SetCookie(w, cookie)
}

func (h *Handler) addUserHandler (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не разрешен", http.StatusMethodNotAllowed)
		return
	}

	var user models.AuthRequest
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	_, err := h.services.GetUser(user.Username, user.Password)
	if err == nil {
		token, err := h.services.GenerateToken(user.Username, user.Password)
		if err != nil {
			http.Error(w, "Не удалось сгенерировать токен для пользователя", http.StatusBadRequest)
		}
		setTokenCookie(w, token)
	}

	id, err := h.services.Auth.CreateUser(user.Username, user.Password)
	if err != nil {
		log.Println("Ошибка при добавлении пользователя:", err)
		http.Error(w, "Ошибка при добавлении пользователя", http.StatusInternalServerError)
		return
	}
	user.Id = id

	token, err := h.services.GenerateToken(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Не удалось сгенерировать токен для пользователя", http.StatusBadRequest)
	}
	setTokenCookie(w, token)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
