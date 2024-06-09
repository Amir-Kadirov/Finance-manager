package service

import (
	"context"
	"finance/api/models"
	"finance/storage"
	"fmt"
)

type customerService struct {
	storage storage.IStorage
}

func NewCustomerService(storage storage.IStorage) customerService {
	return customerService{storage: storage}
}

func (s customerService) Create(ctx context.Context,customer models.Customer) (string, error) {

	id, err := s.storage.CustomerStorage().CreateCustomer(ctx,customer)
	if err != nil {
		fmt.Println("error while creating customer, err: ", err)
		return "", err
	}

	return id, nil
}

func (s customerService) CustomerGetById(ctx context.Context,id string) (models.CustomerGet,error) {
	resp,err:=s.storage.CustomerStorage().CustomerGetById(ctx,id)
	if err!=nil {
		return resp,err
	}

	return resp,nil
}

func (s customerService) Delete(ctx context.Context,id string) error {
	err:=s.storage.CustomerStorage().Delete(ctx,id)
	if err != nil {
		return err
	}

	return nil
}

func (s customerService) PaymentHistory(ctx context.Context,req models.PaymentHistoryRequest) (models.PaymentHistory,error) {
	resp,err:=s.storage.CustomerStorage().PaymentHistory(ctx,req)
	if err!=nil {
		return resp,err
	}

	return resp,nil
}

func (s customerService) ExpenseCalculator(ctx context.Context,id string) (models.ExpenseCalculator,error) {
	resp,err:=s.storage.CustomerStorage().ExpenseCalculator(ctx,id)
	if err!=nil {
		return resp,err
	}

	return resp,nil
}