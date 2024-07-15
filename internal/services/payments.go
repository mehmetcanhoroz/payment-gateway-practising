package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/logger"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/dtos"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/mappers"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/pkg/imposter_bank"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/repository"
)

type PaymentsService struct {
	storage       *repository.PaymentsRepository
	fastStorage   *repository.PaymentsFastRepository
	bankConnector *imposter_bank.Connector
}

func NewPaymentsService(
	storage *repository.PaymentsRepository,
	fastStorage *repository.PaymentsFastRepository,
	bankConn *imposter_bank.Connector,
) *PaymentsService {
	return &PaymentsService{
		storage:       storage,
		fastStorage:   fastStorage,
		bankConnector: bankConn,
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
	err := paymentRequest.Validate()
	if err != nil {
		return &dtos.PostPaymentResponse{PaymentStatus: dtos.PAYMENT_REJECTED}, err
	}

	expiryMonth := ""
	if paymentRequest.ExpiryMonth < 10 {
		expiryMonth = fmt.Sprintf("0%v", paymentRequest.ExpiryMonth)
	}
	bankReq := dtos.BankPaymentRequest{
		CardNumber: paymentRequest.CardNumber,
		ExpiryDate: fmt.Sprintf("%v/%v", expiryMonth, paymentRequest.ExpiryYear),
		Currency:   paymentRequest.Currency,
		Amount:     paymentRequest.Amount,
		Cvv:        paymentRequest.Cvv,
	}

	bankResponse, err := h.bankConnector.MakePayment(context.TODO(), bankReq)
	if err != nil {
		return nil, err
	}

	if bankResponse.Authorized {
		paymentEntity := mappers.MapperPostPaymentRequestToPayments(&paymentRequest)
		payment, err := h.storage.AddPayment(paymentEntity)
		if err != nil {
			return nil, err
		}

		dtoPayment := mappers.MapperPaymentsToPostPaymentResponse(&payment)
		return &dtoPayment, nil
	}

	return nil, errors.New("could not make payment")
}
