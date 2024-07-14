package repository

import (
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/entities"
)

type PaymentsRepository struct {
	payments []entities.Payment
}

func NewPaymentsRepository() *PaymentsRepository {
	return &PaymentsRepository{
		payments: []entities.Payment{},
	}
}

func (ps *PaymentsRepository) GetPayment(id string) (*entities.Payment, error) {
	for _, element := range ps.payments {
		if element.Id == id {
			return &element, nil
		}
	}
	return nil, nil
}

func (ps *PaymentsRepository) AddPayment(payment entities.Payment) {
	ps.payments = append(ps.payments, payment)
}
