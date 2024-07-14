package dtos

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

const (
	PAYMENT_REJECTED   = "rejected"
	PAYMENT_AUTHORIZED = "authorized"
	PAYMENT_DECLINED   = "declined"
)

type PostPaymentRequest struct {
	CardNumber  int    `json:"card_number"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	Currency    string `json:"currency"`
	Amount      int    `json:"amount"`
	Cvv         int    `json:"cvv"`
}

func (dto PostPaymentRequest) Validate() error {
	// Validate Card Number: 14-19 numeric characters
	cardStr := fmt.Sprintf("%d", dto.CardNumber)
	if len(cardStr) < 14 || len(cardStr) > 19 {
		return errors.New("card number must be between 14-19 numeric characters")
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(cardStr) {
		return errors.New("card number must only contain numeric characters")
	}

	// Validate Expiry Month: Between 1-12
	if dto.ExpiryMonth < 1 || dto.ExpiryMonth > 12 {
		return errors.New("expiry month must be between 1 and 12")
	}
	// Validate Expiry Year: Must be in the future
	currentYear := time.Now().Year()
	if dto.ExpiryYear < currentYear {
		return errors.New("expiry year must be in the future")
	}
	// Ensure Expiry Month and Year Combination is in the future
	currentMonth := time.Now().Month()
	if dto.ExpiryYear == currentYear && dto.ExpiryMonth < int(currentMonth) {
		return errors.New("expiry month and year combination must be in the future")
	}

	// Validate Currency: Must be one of the allowed ISO currency codes (example: USD, EUR, GBP)
	validCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"GBP": true,
	}
	if !validCurrencies[dto.Currency] {
		return errors.New("currency must be one of 'USD', 'EUR', 'GBP'")
	}

	// Validate Currency Format: Exactly 3 characters
	if len(dto.Currency) != 3 {
		return errors.New("currency must be exactly 3 characters")
	}

	// Validate Amount: Must be an integer
	if dto.Amount <= 0 {
		return errors.New("amount must be a positive integer representing the amount in the minor currency unit")
	}

	// Validate CVV: 3-4 numeric characters
	cvvStr := fmt.Sprintf("%d", dto.Cvv)
	if len(cvvStr) < 3 || len(cvvStr) > 4 {
		return errors.New("CVV must be between 3-4 numeric characters")
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(cvvStr) {
		return errors.New("CVV must only contain numeric characters")
	}

	return nil
}

type PostPaymentResponse struct {
	Id                 string `json:"id,omitempty"`
	PaymentStatus      string `json:"payment_status"`
	CardNumberLastFour int    `json:"card_number_last_four,omitempty"`
	ExpiryMonth        int    `json:"expiry_month,omitempty"`
	ExpiryYear         int    `json:"expiry_year,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Amount             int    `json:"amount,omitempty"`
}

type GetPaymentResponse struct {
	Id                 string `json:"id,omitempty"`
	PaymentStatus      string `json:"payment_status"`
	CardNumberLastFour int    `json:"card_number_last_four,omitempty"`
	ExpiryMonth        int    `json:"expiry_month,omitempty"`
	ExpiryYear         int    `json:"expiry_year,omitempty"`
	Currency           string `json:"currency,omitempty"`
	Amount             int    `json:"amount,omitempty"`
}
