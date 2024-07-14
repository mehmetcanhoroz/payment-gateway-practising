package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/repository"
	"github.com/go-chi/chi"
)

type PaymentsHandler struct {
	storage     *repository.PaymentsRepository
	fastStorage *repository.PaymentsFastRepository
}

func NewPaymentsHandler(storage *repository.PaymentsRepository) *PaymentsHandler {
	return &PaymentsHandler{
		storage: storage,
	}
}

// GetHandler returns an http.HandlerFunc that handles HTTP GET requests.
// It retrieves a payment record by its ID from the storage.
// The ID is expected to be part of the URL.
func (h *PaymentsHandler) GetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		payment := h.storage.GetPayment(id)

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
