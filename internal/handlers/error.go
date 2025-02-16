package handlers

import (
    "encoding/json"
    "net/http"
	"avito-shop/models"
)

func writeErrorResponse(w http.ResponseWriter, statusCode int, message string) {
    w.WriteHeader(statusCode)
    errorResponse := models.ErrorResponse{Errors: message}
    json.NewEncoder(w).Encode(errorResponse)
}
