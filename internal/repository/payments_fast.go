package repository

import (
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/entities"
)

type PaymentsFastRepository struct {
	payments []entities.Payment
}

func NewPaymentsFastRepository() *PaymentsFastRepository {
	return &PaymentsFastRepository{
		payments: []entities.Payment{},
	}
}

func (ps *PaymentsFastRepository) GetPayment(id string) *entities.Payment {
	for _, element := range ps.payments {
		if element.Id == id {
			return &element
		}
	}
	return nil
}

func (ps *PaymentsFastRepository) AddPayment(payment entities.Payment) {
	ps.payments = append(ps.payments, payment)
}
