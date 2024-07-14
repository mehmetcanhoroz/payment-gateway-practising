package mappers

import (
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/dtos"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/entities"
)

func MapperPaymentsToGetPaymentResponse(payment *entities.Payment) dtos.GetPaymentResponse {
	return dtos.GetPaymentResponse{
		Amount:      payment.Amount,
		Currency:    payment.Currency,
		ExpiryMonth: payment.ExpiryMonth,
		ExpiryYear:  payment.ExpiryYear,
	}
}
