package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecurringTransactionService struct {
	DbPool *pgxpool.Pool
}

func (c *RecurringTransactionService) Get(userId string) ([]models.RecurringTransaction, error) {
	query := `SELECT * 
		FROM recurring_transactions
		WHERE user_id = @userId
		ORDER BY id ASC
	`

	args := pgx.NamedArgs{
		"userId": userId,
	}

	output := []models.RecurringTransaction{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query, args)
	if err != nil {
		return []models.RecurringTransaction{}, nil
	}

	return output, nil
}

func (c *RecurringTransactionService) Post(userId string, payload models.RecurringTransaction) error {
	query := `INSERT INTO 
		categories (id, local_id, user_id, name, frequency, start_date, next_due_date, category_id, deduct_from_account_id, amount, notes, created_date)
		VALUES (@id, @localId, @userId, @name, @frequency, @startDate, @nextDueDate, @categoryId, @deductFromAccountId, @amount, @notes, @createdDate)
	;`

	args := pgx.NamedArgs{
		"id":                  uuid.NewString(),
		"userId":              userId,
		"localId":             payload.LocalId,
		"name":                payload.Name,
		"frequency":           payload.Frequency,
		"startDate":           payload.StartDate,
		"nextDueDate":         payload.NextDueDate,
		"categoryId":          payload.CategoryId,
		"deductFromAccountId": payload.DeductFromAccountId,
		"amount":              payload.Amount,
		"notes":               payload.Notes,
		"createdDate":         payload.CreatedDate.Time,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *RecurringTransactionService) Put(id string, payload models.RecurringTransaction) error {
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
		"updatedDate":         payload.UpdatedDate.Time,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *RecurringTransactionService) Delete(id string) error {
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
