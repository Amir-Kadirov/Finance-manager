package service

import (
	"context"
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
	Data, err := s.storage.TransactionStorage().TransactionToCard(ctx, transaction)
	if err != nil {
		log.Fatal("error while transaction card crud, err %w: ", err)
		return "", err
	}

	// logic
	db := s.storage.GetDB()

	msg := fmt.Sprintf("From : %s To : %s amount: %d",
		transaction.FromCard, transaction.ToCard, transaction.Bill)

	queryPaymentInfo := `INSERT INTO payment_info (customer_id, card_id, history,category) VALUES ($1, $2, $3,$4)`
	_, err = db.Exec(ctx, queryPaymentInfo, Data.CustomerIdSend, transaction.FromCard, msg, transaction.Category)
	if err != nil {
		return "", err
	}

	queryExpense := `UPDATE expense SET total_expense=total_expense+$1,` + transaction.Category + `=` + transaction.Category + `+$1 WHERE customer_id=$2`
	_, err = db.Exec(ctx, queryExpense, transaction.Bill, Data.CustomerIdSend)
	if err != nil {
		return "", err
	}

	return msg, nil
}
