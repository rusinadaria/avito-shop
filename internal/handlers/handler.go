package handlers

import (
	"avito-shop/internal/services"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	// "encoding/json"
	// "avito-shop/models"
	// "log"
)

type Handler struct {
	services *services.Service
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth_token")
		if err != nil {
			http.Error(w, "Токен не найден", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("qrkjk#4#%35FSFJlja#4353KSFjH"), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Необходима аутентификация!", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Post("/api/auth", h.addUserHandler)

		r.With(h.AuthMiddleware).Group(func(r chi.Router) {
			r.Get("/api/info", h.InfoHandler)
			// r.Get("/api/buy/{item}", )
			r.Post("/api/sendCoin", h.SendHandler)
		})
	})
	return r
}