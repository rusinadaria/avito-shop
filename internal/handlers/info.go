package handlers

import (
	"net/http"
	"encoding/json"
)

func (h *Handler) InfoHandler (w http.ResponseWriter, r *http.Request) { // Получить информацию о монетках, инвентаре и истории транзакций.
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	cookie, err := r.Cookie("auth_token")
	if err != nil {
		writeErrorResponse(w, http.StatusBadRequest, "Ошибка при попытке получить cookies")
		return
	}

	id, err := h.services.ParseToken(cookie.Value)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Ошибка при попытке распарсить cookies")
		return
	}

	info, err := h.services.UserInfo(id)
	if err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Не удалось получить информацию о пользователе")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(info)
}