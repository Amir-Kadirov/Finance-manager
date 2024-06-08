package storage

import (
	"context"
	"finance/api/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IStorage interface {
	CloseDB()
	CustomerStorage() CustomerStorage
	CardStorage() CardStorage
	TransactionStorage() TransactionStorage
	GetDB() *pgxpool.Pool
}

type CustomerStorage interface {
	CreateCustomer(ctx context.Context, customer models.Customer) (string, error)
	CustomerGetById(ctx context.Context,id string) (models.CustomerGet,error)
	PaymentHistory(ctx context.Context,id string) (models.PaymentHistory,error)
	ExpenseCalculator(ctx context.Context, id string) (models.ExpenseCalculator, error)
}

type CardStorage interface {
	CreateCard(ctx context.Context, card models.CreateCard) (string, error)
}

type TransactionStorage interface {
	TransactionToCard(ctx context.Context, transaction models.TransactionToCard) (models.TransactionServer, error)
}
