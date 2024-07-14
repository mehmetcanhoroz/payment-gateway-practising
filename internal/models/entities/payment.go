package entities

type Payment struct {
	Id                 string `json:"id"`
	CardNumberLastFour int    `json:"card_number_last_four"`
	ExpiryMonth        int    `json:"expiry_month"`
	ExpiryYear         int    `json:"expiry_year"`
	Currency           string `json:"currency"`
	Amount             int    `json:"amount"`
	Cvv                int    `json:"cvv"`
}
