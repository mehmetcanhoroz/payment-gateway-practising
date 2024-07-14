package handlers

import (
	"encoding/json"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/logger"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/dtos"
	"net/http"

	"github.com/go-chi/chi/v5"

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
		println(chi.URLParam(r, "id"))
		id := chi.URLParam(r, "id")
		payment, err := h.service.GetPayment(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			logger.Error("error getting payment: %v\n", err)
			return
		}

		if payment != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(payment); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			w.WriteHeader(http.StatusNoContent)
		}

		println(chi.URLParam(r, "id"))
	}
}

func (h *PaymentsHandler) GetHandlerV2(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	payment, err := h.service.GetPayment(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logger.Error("error getting payment: %v\n", err)
		return
	}

	if payment != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err = json.NewEncoder(w).Encode(payment); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

// PostHandler returns an http.HandlerFunc that handles HTTP POST requests.
// It retrieves a payment object and passes it to service layer to save and proceed
func (h *PaymentsHandler) PostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var paymentRequest dtos.PostPaymentRequest
		err := decoder.Decode(&paymentRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			err = json.NewEncoder(w).Encode("wrong request body")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		payment, err := h.service.MakePayment(paymentRequest)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err = json.NewEncoder(w).Encode(payment); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
