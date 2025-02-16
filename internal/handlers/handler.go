package handlers

import (
	"avito-shop/internal/services"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"log/slog"
	// "avito-shop/internal/handlers/middleware"
	"avito-shop/internal/common"
)

const signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"

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
			common.WriteErrorResponse(w, http.StatusUnauthorized, "Токен не найден")
			return
		}

		tokenString := cookie.Value
		claims := &jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		})

		if err != nil || !token.Valid {
			common.WriteErrorResponse(w, http.StatusUnauthorized, "Неавторизован")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) InitRoutes(logger *slog.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(LoggerMiddlewareWrapper(logger))

	r.Group(func(r chi.Router) {
		r.Post("/api/auth", h.AddUserHandler)

		r.With(h.AuthMiddleware).Group(func(r chi.Router) {
		// r.With(middleware.AuthMiddleware(h)).Group(func(r chi.Router) {
			r.Get("/api/info", h.InfoHandler)
			r.Get("/api/buy/{item}", h.BuyItemHandler)
			r.Post("/api/sendCoin", h.SendHandler)
		})
	})
	return r
}