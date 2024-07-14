package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/services"
)

type PaymentsHandler struct {
	service *services.PaymentsService
}

func NewPaymentsHandler(paymentsService *services.PaymentsService) *PaymentsHandler {
	return &PaymentsHandler{
		service: paymentsService,
	}
}

// GetHandler returns a http.HandlerFunc that handles HTTP GET requests.
// It retrieves a payment record by its ID from the storage.
// The ID is expected to be part of the URL.
func (h *PaymentsHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		payment, err := h.service.GetPayment(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if payment != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(payment); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

// PostHandler returns an http.HandlerFunc that handles HTTP POST requests.
// It retrieves a payment object and passes it to service layer to save and proceed
func (ph *PaymentsHandler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).Encode("post method")
		if err != nil {
			return
		}
	}
}
