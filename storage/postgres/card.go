package postgres

import (
	"context"
	"errors"
	"finance/api/models"
	"finance/pkg/hash"
	"fmt"
	"math/rand"
	"time"

	"github.com/nsuprun/ccgen"

	"github.com/jackc/pgx/v5/pgxpool"
)

type cardRepo struct {
	db *pgxpool.Pool
}

func NewCard(db *pgxpool.Pool) cardRepo {
	return cardRepo{
		db: db,
	}
}

func (c *cardRepo) CreateCard(ctx context.Context, card models.CreateCard) (string, error) {
	cardId := ccgen.Visa.GenerateOfLength(15)
	cvv := fmt.Sprintf("%03d", rand.Intn(999)+1)
	expiryDate := time.Now().AddDate(3, 0, 0)
	hashedPassword, err := hash.HashPassword(card.Password)

	if err != nil {
		return "", errors.New("error while hashing password")
	}

	var name string

	queryCustomer := `UPDATE customer SET card_id=$1 WHERE id=$2 RETURNING first_name`
	row := c.db.QueryRow(ctx, queryCustomer, cardId, card.CustomerId)

	if err := row.Scan(&name); err != nil {
		return "", err
	}

	query := `INSERT INTO card (customer_id,id,card_holder_name,cvv,expity_date,password) 
	                  VALUES ($1,$2,$3,$4,$5,$6)`
	_, err = c.db.Exec(ctx, query, card.CustomerId, cardId, name, cvv, expiryDate, hashedPassword)
	if err != nil {
		return "", err
	}

	return cardId, nil
}
