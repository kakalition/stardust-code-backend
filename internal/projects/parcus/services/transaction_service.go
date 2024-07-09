package services

import (
	"context"
	"time"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService struct {
	DbPool *pgxpool.Pool
}

func (c *TransactionService) Get() ([]models.Transaction, error) {
	query := `SELECT * 
		FROM transactions
		ORDER BY id ASC
	;`

	output := []models.Transaction{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		return []models.Transaction{}, err
	}

	return output, nil
}

func (c *TransactionService) Post(payload models.Transaction) error {
	query := `INSERT INTO 
		transactions (user_id, date, walletId, categoryId, amount, notes, is_recurring, created_date, updated_date)
		VALUES (@userId, @date, @walletId, @categoryId, @amount, @notes, @isRecurring, @createdDate, @updatedDate)
	;`

	args := pgx.NamedArgs{
		"userId":      payload.UserId,
		"date":        payload.Date,
		"walletId":    payload.WalletId,
		"categoryId":  payload.CategoryId,
		"amount":      payload.Amount,
		"notes":       payload.Notes,
		"isRecurring": payload.IsRecurring,
		"createdDate": time.Now(),
		"updatedDate": nil,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *TransactionService) Put(id int, payload models.Transaction) error {
	query := `UPDATE transactions
		SET date=@date, wallet_id=@walletId, category_id=@categoryId, amount=@amount, 
			notes=@notes, is_recurring=@isRecurring, updated_date=@updatedDate
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":          id,
		"date":        payload.Date,
		"walletId":    payload.WalletId,
		"categoryId":  payload.CategoryId,
		"amount":      payload.Amount,
		"notes":       payload.Notes,
		"isRecurring": payload.IsRecurring,
		"updatedDate": time.Now(),
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *TransactionService) Delete(id int) error {
	query := `DELETE FROM transactions 
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}
