package handlers

import (
	"avito-shop/models"
	"encoding/json"
	"net/http"
	"time"
)

func setTokenCookie(w http.ResponseWriter, token string) {
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,     
		Secure:   false,
	}

	http.SetCookie(w, cookie)
}

func (h *Handler) AddUserHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")

    var user models.AuthRequest
	var userID int
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        writeErrorResponse(w, http.StatusBadRequest, "Неверный запрос")
        return
    }

    _, err := h.services.FindUser(user.Username)
    if err != nil {
        userID, err = h.services.CreateUser(user.Username, user.Password)
        if err != nil {
            writeErrorResponse(w, http.StatusInternalServerError, "Не удалось создать пользователя")
            return
        }
    } else {
        userID, err = h.services.SignIn(user.Username, user.Password)
        if err != nil {
            writeErrorResponse(w, http.StatusUnauthorized, "Неавторизован")
            return
        }
    }

    token, err := h.services.GenerateToken(userID)
    if err != nil {
        writeErrorResponse(w, http.StatusInternalServerError, "Не удалось сгенерировать токен для пользователя")
        return
    }
    setTokenCookie(w, token)

    w.WriteHeader(http.StatusOK)
}