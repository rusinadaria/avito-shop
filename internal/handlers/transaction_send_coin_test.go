package handlers

import (
	mock_services "avito-shop/internal/services/mocks"
	"avito-shop/internal/services"
	"avito-shop/models"
	"bytes"
	"net/http/httptest"
	"testing"
	"github.com/go-chi/chi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
)

func TestHandler_SendHandler(t *testing.T) {
	type mockBehavior func(s *mock_services.MockTransaction, userId int, req models.SendCoinRequest)

	testTable := []struct {
		name string
		inputBody string
		inputUser models.SendCoinRequest
		mockBehavior mockBehavior
		expectedStatusCode int
		expectedRequestBody string
	} {
		{
			name: "OK",
			inputBody: `{"toUser":"test_user", "amount":50}`,
			inputUser: models.SendCoinRequest {
				ToUser: "test_user",
				Amount: 50,
			},
			mockBehavior: func(s *mock_services.MockTransaction, userId int, req models.SendCoinRequest) {
				s.EXPECT().SendCoin(userId, req.ToUser, req.Amount).Return(nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: ``,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			transact := mock_services.NewMockTransaction(c)
			testCase.mockBehavior(transact, 1, testCase.inputUser)

			auth := mock_services.NewMockAuth(c)
			auth.EXPECT().ParseToken("some_valid_token").Return(1, nil)

			services := &services.Service{Transaction: transact, Auth: auth}
			handler := NewHandler(services)

			r := chi.NewRouter()
			r.Post("/api/sendCoin", handler.SendHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/sendCoin", bytes.NewBufferString(testCase.inputBody))

			req.AddCookie(&http.Cookie{Name: "auth_token", Value: "some_valid_token"})

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}