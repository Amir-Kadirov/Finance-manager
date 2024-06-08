package models

type Card struct{
	CustomerId string `json:"customer_id"`
	Id string `json:"id"`
	CardHolderName string `json:"card_holder_name"`
	Cvv int `json:"cvv"`
	ExpityDate string `json:"expity_date"`
	Password string `json:"password"`
	Balance uint `json:"balance"`
}

type CreateCard struct{
	CustomerId string `json:"customer_id"`
	Password string `json:"password"`
}