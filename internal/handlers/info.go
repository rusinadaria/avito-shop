package handlers

import (
	"net/http"
	"encoding/json"
)

func (h *Handler) InfoHandler (w http.ResponseWriter, r *http.Request) { // Получить информацию о монетах, инвентаре и истории транзакций.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	response := map[string]string{
		"message": "Добро пожаловать в наше приложение! Info page",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}