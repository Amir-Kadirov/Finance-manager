package service

import (
	"context"
	"errors"
	"finance/api/models"
	"finance/storage"
	"fmt"
	"log"
)

type transactionService struct {
	storage storage.IStorage
}

func NewTransactionService(storage storage.IStorage) transactionService {
	return transactionService{storage: storage}
}

func (s transactionService) TransactionToCard(ctx context.Context, transaction models.TransactionToCard) (string, error) {

	// crud
	Data, err := s.storage.TransactionStorage().TransactionToCard(ctx,transaction)
	if err != nil {
		log.Fatalf("Error while transaction card crud, err: ", err)
		return "", err
	}

	// logic
	db:=s.storage.GetDB()
	
	msg := fmt.Sprintf("from card: %s to card: %s amount: %d",
		transaction.FromCard, transaction.ToCard, transaction.Bill)

	queryPaymentInfo := `INSERT INTO payment_info (customer_id, card_id, history) VALUES ($1, $2, $3)`
	_, err = db.Exec(ctx, queryPaymentInfo, Data.CustomerIdSend, transaction.FromCard, msg)
	if err != nil {
		return "", errors.New("error while inserting to history")
	}

	queryExpense := `UPDATE expense SET total_expense=total_expense+$1,` + transaction.Category + `=` + transaction.Category +`+$1 WHERE customer_id=$2`
	_, err = db.Exec(ctx, queryExpense, transaction.Bill, Data.CustomerIdSend)
	if err != nil {
		return "", err
	}

	
	fmt.Println(queryExpense)

	return msg, nil
}