package postgres

import (
	"context"
	"errors"
	"finance/api/models"		
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type transactionRepo struct {
	db *pgxpool.Pool
}

func NewTransaction(db *pgxpool.Pool) transactionRepo {
	return transactionRepo{
		db: db,
	}
}

func (t *transactionRepo) TransactionToCard(ctx context.Context, transaction models.TransactionToCard) (models.TransactionServer, error) {
	resp:=models.TransactionServer{TransactionToCard: transaction}
	tx, err := t.db.Begin(ctx)
	if err != nil {
		return resp, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			log.Fatalf("panic occurred: %v", p)
		} else if err != nil {
			tx.Rollback(ctx)
			log.Println("rolled back due to error:", err)
		}
	}()


	querySelect := `SELECT balance FROM card WHERE id=$1`
	err = tx.QueryRow(ctx, querySelect, transaction.FromCard).Scan(&resp.Balance)
	if err != nil {
		return resp, err
	}
	if resp.Balance < transaction.Bill {
		return resp, errors.New("not enough amount of money")
	}

	queryP1 := `UPDATE card SET balance=balance-$1 WHERE id=$2 RETURNING customer_id`
	err = tx.QueryRow(ctx, queryP1, transaction.Bill, transaction.FromCard).Scan(&resp.CustomerIdSend)
	if err != nil {
		return resp, err
	}

	queryP2 := `UPDATE card SET balance=balance+$1 WHERE id=$2 RETURNING customer_id`
	err = tx.QueryRow(ctx, queryP2, transaction.Bill, transaction.ToCard).Scan(&resp.CustomerIdReceive)
	if err != nil {
		return resp, err
	}


	if err := tx.Commit(ctx); err != nil {
		return resp, err
	}
	
	return resp, nil
}