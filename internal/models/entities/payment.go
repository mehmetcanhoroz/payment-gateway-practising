package entities

type Payment struct {
	Id          string `json:"id"`
	CardNumber  int    `json:"card_number"`
	ExpiryMonth int    `json:"expiry_month"`
	ExpiryYear  int    `json:"expiry_year"`
	Currency    string `json:"currency"`
	Amount      int    `json:"amount"`
	Cvv         int    `json:"cvv"`
}
