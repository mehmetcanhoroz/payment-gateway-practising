package mappers

import (
	"strconv"

	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/dtos"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/entities"
)

func MapperPaymentsToGetPaymentResponse(payment *entities.Payment) dtos.GetPaymentResponse {
	return dtos.GetPaymentResponse{
		Id:                 payment.Id,
		CardNumberLastFour: last4DigitsOfCard(payment.CardNumber),
		Amount:             payment.Amount,
		Currency:           payment.Currency,
		ExpiryMonth:        payment.ExpiryMonth,
		ExpiryYear:         payment.ExpiryYear,
	}
}

func MapperPostPaymentRequestToPayments(payment *dtos.PostPaymentRequest) entities.Payment {
	return entities.Payment{
		Amount:      payment.Amount,
		Currency:    payment.Currency,
		ExpiryMonth: payment.ExpiryMonth,
		ExpiryYear:  payment.ExpiryYear,
		Cvv:         payment.Cvv,
		CardNumber:  payment.CardNumber,
	}
}

func MapperPaymentsToPostPaymentResponse(payment *entities.Payment) dtos.PostPaymentResponse {
	return dtos.PostPaymentResponse{
		Amount:             payment.Amount,
		PaymentStatus:      dtos.PAYMENT_AUTHORIZED,
		Currency:           payment.Currency,
		ExpiryMonth:        payment.ExpiryMonth,
		ExpiryYear:         payment.ExpiryYear,
		CardNumberLastFour: last4DigitsOfCard(payment.CardNumber),
		Id:                 payment.Id,
	}
}

func last4DigitsOfCard(cardNumber int) int {
	stringCard := strconv.Itoa(cardNumber)
	lastFourDigit, _ := strconv.Atoi(stringCard[len(stringCard)-4:])
	return lastFourDigit
}
