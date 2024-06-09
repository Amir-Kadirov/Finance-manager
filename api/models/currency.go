package models

type ExchangeRateRequest struct {
	Amount int `json:"amount"`
	From  string `json:"from"`
	To string	`json:"to"`
}