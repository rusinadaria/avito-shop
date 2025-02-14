package handlers

import (
	"avito-shop/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func (h *Handler) SendHandler (w http.ResponseWriter, r *http.Request) { // Отправить монеты другому пользователю.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := map[string]string{
		"message": "Send page",
	}

	var req models.SendCoinRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("user_id")
	if err != nil {
		http.Error(w, "Не удалось получить куку", http.StatusUnauthorized)
		return
	}

	senderId := cookie.Value

	err = h.services.SendCoin(senderId, req.ToUser, req.Amount)
	fmt.Println(req.ToUser)
	if err != nil {
		log.Println("Ошибка при попытке отправить коины")
	}


	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}