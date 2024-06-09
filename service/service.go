package service

import (
	"finance/storage"
)


type IServiceManager interface {
	Customer() customerService
	Card() cardService
	Transaction() transactionService
	Currency() currencyService
}

type Service struct {
	customerService customerService
	cardService cardService
	transactionService transactionService
	currencyService currencyService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.customerService = NewCustomerService(storage)
	services.cardService = NewCardService(storage)
	services.transactionService=NewTransactionService(storage)
	services.currencyService=NewCurrencyService(storage)

	return services
}

func (s Service) Customer() customerService {
	return s.customerService
}


func (s Service) Card() cardService {
	return s.cardService
}

func (s Service) Transaction() transactionService {
	return s.transactionService
}

func (s Service) Currency() currencyService {
	return s.currencyService
}