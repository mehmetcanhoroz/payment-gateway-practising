package services

import (
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/dtos"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/mappers"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/repository"
)

type PaymentsService struct {
	storage     *repository.PaymentsRepository
	fastStorage *repository.PaymentsFastRepository
}

func NewPaymentsService(storage *repository.PaymentsRepository, fastStorage *repository.PaymentsFastRepository) *PaymentsService {
	return &PaymentsService{
		storage:     storage,
		fastStorage: fastStorage,
	}
}

// GetPayment ...
func (h *PaymentsService) GetPayment(id string) (*dtos.GetPaymentResponse, error) {
	payment, err := h.storage.GetPayment(id)
	if err != nil {
		return nil, err
	}

	if payment == nil {
		return nil, nil
	}

	dtoPayment := mappers.MapperPaymentsToGetPaymentResponse(payment)

	return &dtoPayment, nil
}
