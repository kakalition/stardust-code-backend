package services

import (
	"context"
	"time"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecurringTransactionService struct {
	DbPool *pgxpool.Pool
}

func (c *RecurringTransactionService) Get() ([]models.RecurringTransaction, error) {
	query := `SELECT * 
		FROM recurring_transactions
		ORDER BY id ASC
	`

	output := []models.RecurringTransaction{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		return []models.RecurringTransaction{}, nil
	}

	return output, nil
}

func (c *RecurringTransactionService) Post(payload models.RecurringTransaction) error {
	query := `INSERT INTO 
		categories (user_id, name, frequency, start_date, next_due_date, category_id, deduct_from_account_id, amount, notes, created_date)
		VALUES (@userId, @name, @frequency, @startDate, @nextDueDate, @categoryId, @deductFromAccountId, @amount, @notes, @createdDate)
	;`

	args := pgx.NamedArgs{
		"userId":              payload.UserId,
		"name":                payload.Name,
		"frequency":           payload.Frequency,
		"startDate":           payload.StartDate,
		"nextDueDate":         payload.NextDueDate,
		"categoryId":          payload.CategoryId,
		"deductFromAccountId": payload.DeductFromAccountId,
		"amount":              payload.Amount,
		"notes":               payload.Notes,
		"createdDate":         time.Now(),
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *RecurringTransactionService) Put(id int, payload models.RecurringTransaction) error {
	query := `UPDATE recurring_transactions
		SET name=@name, frequency=@frequency, start_date=@startDate, next_due_date=@nextDueDate, 
			category_id=@categoryId, deduct_from_account_id=@deductFromAccountId, amount=@amount, 
			notes=@notes, updated_date=@updatedDate
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"name":                payload.Name,
		"frequency":           payload.Frequency,
		"start_date":          payload.StartDate,
		"next_due_date":       payload.NextDueDate,
		"categoryId":          payload.CategoryId,
		"deductFromAccountId": payload.DeductFromAccountId,
		"amount":              payload.Amount,
		"notes":               payload.Notes,
		"updatedDate":         time.Now(),
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *RecurringTransactionService) Delete(id int) error {
	query := `DELETE FROM recurring_transactions 
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
