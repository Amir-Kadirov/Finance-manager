package postgres

import (
	"context"
	"errors"
	"finance/api/models"
	"finance/pkg"
	"finance/pkg/hash"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type customerRepo struct {
	db *pgxpool.Pool
}

func NewCustomer(db *pgxpool.Pool) customerRepo {
	return customerRepo{
		db: db,
	}
}

func (c *customerRepo) CreateCustomer(ctx context.Context, customer models.Customer) (string, error) {
	id := uuid.New().String()

	hashingPassword, err := hash.HashPassword(customer.Passowrd)
	if err != nil {
		return "", errors.New("error while hashing password")
	}

	query := `INSERT INTO customer (id,first_name,last_name,gmail,password,phone)
	                      VALUES ($1,$2,$3,$4,$5,$6)`
	_, err = c.db.Query(ctx, query, id, customer.FirstName, customer.LastName, customer.Gmail, hashingPassword, customer.Phone)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *customerRepo) CustomerGetById(ctx context.Context, id string) (models.CustomerGet, error) {
	resp := models.CustomerGet{}

	query := `SELECT 
	                cus.last_name,
					cus.phone,
					cus.gmail,
					c.balance,
					c.card_holder_name,
					t.history,
					t.created_at 
					FROM customer cus 
					JOIN card c ON cus.id=c.customer_id 
					JOIN payment_info t on t.customer_id=cus.id 
					WHERE cus.id=$1`
	row,err := c.db.Query(ctx, query, id)
	if err!=nil {
		return resp,err
	}
	for row.Next(){
	card := models.CustomerCard{}
	info := models.TransactionInfo{}
	if err := row.Scan(
		&resp.FirstName,
		&resp.Phone,
		&resp.Gmail,
		&card.Balance,
		&card.CardHolderName,
		&info.History,
		&info.CreartedAt,
	); err != nil {
		return resp, err
	}
	resp.Info = append(resp.Info, info)
	resp.Card = append(resp.Card, card)
}

	return resp, nil
}


func (c *customerRepo) PaymentHistory(ctx context.Context, id string) (models.PaymentHistory,error) {
	PaymentHistory:=models.PaymentHistory{}

	query:=`SELECT card_id,history,created_at FROM payment_info WHERE customer_id=$1`
	row,err:=c.db.Query(ctx,query,id)
	if err!=nil {
		return PaymentHistory,err
	}

	for row.Next() {
		TransactionInfo:=models.TransactionInfo{}
		if err:=row.Scan(
			&TransactionInfo.CardId,
			&TransactionInfo.History,
			&TransactionInfo.CreartedAt,
			); err!=nil {
			return PaymentHistory,err
		}
		PaymentHistory.TransactionInfo=append(PaymentHistory.TransactionInfo, TransactionInfo)
	}

	return PaymentHistory,nil
}


func (c *customerRepo) ExpenseCalculator(ctx context.Context, id string) (models.ExpenseCalculator, error) {
	expense:=models.ExpenseCalculator{}

	query:=`SELECT total_expense,
				   restaurant,
				   supermarkets,
				   beauty_medecine,
				   entertaintment_sport FROM expense WHERE customer_id=$1`

	err:=c.db.QueryRow(ctx,query,id).Scan(&expense.TotalExpense,
		 								  &expense.Restaurant,
										  &expense.Supermarkets,
										  &expense.BeautyMedecine,
										  &expense.EnerprainmentSport)
	if err != nil {
		return expense, err
	}

	expense,err=pkg.PrecentageCalculator(expense)
	if err!=nil {
		return expense,err
	}

	return expense, nil
}