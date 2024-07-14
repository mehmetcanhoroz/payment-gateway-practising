package repository

import (
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/logger"
	"github.com/cko-recruitment/payment-gateway-challenge-go/internal/models/entities"
	"github.com/google/uuid"
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
	logger.Debug("DEBUG: getting payment with id %s in Repository level", id)
	for _, element := range ps.payments {
		if element.Id == id {
			return &element, nil
		}
	}
	return nil, nil
}

func (ps *PaymentsRepository) AddPayment(payment entities.Payment) (entities.Payment, error) {
	uuidGenerator, _ := uuid.NewV7()
	payment.Id = uuidGenerator.String()
	ps.payments = append(ps.payments, payment)
	return payment, nil
}
