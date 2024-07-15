package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionService struct {
	DbPool *pgxpool.Pool
}

func (c *TransactionService) Get(userId string) ([]models.Transaction, error) {
	query := `SELECT * 
		FROM transactions
		WHERE user_id = @userId
		ORDER BY id ASC
	;`

	args := pgx.NamedArgs{
		"userId": userId,
	}

	output := []models.Transaction{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query, args)
	if err != nil {
		return []models.Transaction{}, err
	}

	return output, nil
}

func (c *TransactionService) Post(userId string, payload models.Transaction) error {
	query := `INSERT INTO 
		transactions (id, local_id, user_id, date, accountId, categoryId, amount, notes, is_recurring, created_date, updated_date)
		VALUES (@id, @localId, @userId, @date, @accountId, @categoryId, @amount, @notes, @isRecurring, @createdDate, @updatedDate)
	;`

	args := pgx.NamedArgs{
		"id":          uuid.NewString(),
		"localId":     payload.LocalId,
		"userId":      userId,
		"date":        payload.Date,
		"accountId":   payload.AccountId,
		"categoryId":  payload.CategoryId,
		"amount":      payload.Amount,
		"notes":       payload.Notes,
		"isRecurring": payload.IsRecurring,
		"createdDate": payload.CreatedDate.Time,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *TransactionService) Put(id string, payload models.Transaction) error {
	query := `UPDATE transactions
		SET date=@date, wallet_id=@accountId, category_id=@categoryId, amount=@amount, 
			notes=@notes, is_recurring=@isRecurring, updated_date=@updatedDate
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":          id,
		"date":        payload.Date,
		"accountId":   payload.AccountId,
		"categoryId":  payload.CategoryId,
		"amount":      payload.Amount,
		"notes":       payload.Notes,
		"isRecurring": payload.IsRecurring,
		"updatedDate": payload.CreatedDate.Time,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *TransactionService) Delete(id string) error {
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
