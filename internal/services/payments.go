package services

import (
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/logger"
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
	logger.Debug("DEBUG: getting payment with id %s in service level", id)
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

// MakePayment ...
func (h *PaymentsService) MakePayment(paymentRequest dtos.PostPaymentRequest) (*dtos.PostPaymentResponse, error) {
	logger.Debug("DEBUG: payment with id %s in service level")
	paymentEntity := mappers.MapperPostPaymentRequestToPayments(&paymentRequest)
	payment, err := h.storage.AddPayment(paymentEntity)
	if err != nil {
		return nil, err
	}

	dtoPayment := mappers.MapperPaymentsToPostPaymentResponse(&payment)

	return &dtoPayment, nil
}
