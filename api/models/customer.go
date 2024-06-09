package models

import "time"

type Customer struct {
	Id        string `json:"id"`
	Phone     string `json:"phone"`
	Passowrd  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gmail     string `json:"gmail"`
}

type CustomerGet struct {
	Phone     string            `json:"phone"`
	FirstName string            `json:"first_name"`
	Gmail     string            `json:"gmail"`
	Card      []CustomerCard    `json:"customer_card"`
	Info      []TransactionInfo `json:"info"`
}

type CustomerCard struct {
	CardHolderName string `json:"card_holder_name"`
	Balance        uint   `json:"balance"`
}

type TransactionInfo struct {
	History    string    `json:"history"`
	CreartedAt time.Time `json:"created_at"`
	CardId     string    `json:"card_id"`
}

type PaymentHistoryRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
	Id string `json:"id"`
}

type PaymentHistory struct {
	TransactionInfo []TransactionInfo `json:"payment_history"`
}

type ExpenseCalculator struct {
	TotalExpense       float64 `json:"total_expense"`
	Restaurant         float64 `json:"restaurant"`
	Supermarkets       float64 `json:"supermarkets"`
	BeautyMedecine     float64 `json:"beauty_medecine"`
	EnerprainmentSport float64 `json:"entertaintment_sport"`
}
