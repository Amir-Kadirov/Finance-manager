package models

type TransactionToCard struct {
	FromCard string `json:"from_card"`
	ToCard   string `json:"to_card"`
	Bill     uint   `json:"bill"`
	Category string `json:"category"`
}

type TransactionServer struct {
	CustomerIdSend    string
	Balance           uint
	CustomerIdReceive string
	TransactionToCard TransactionToCard
}