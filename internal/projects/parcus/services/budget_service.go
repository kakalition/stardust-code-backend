package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BudgetService struct {
	DbPool *pgxpool.Pool
}

func (c *BudgetService) Get(userId string) ([]models.Budget, error) {
	query := `SELECT * 
		FROM budgets
		WHERE user_id = @userId
		ORDER BY id ASC
	;`

	args := pgx.NamedArgs{
		"userId": userId,
	}

	output := []models.Budget{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query, args)
	if err != nil {
		return []models.Budget{}, err
	}

	return output, nil
}

func (c *BudgetService) Post(userId string, payload models.Budget) error {
	query := `INSERT INTO 
		budgets (id, local_id, user_id, category_id, amount, balance, period) 
		VALUES (@id, @localId, @userId, @categoryId, @amount, @period)
	;`

	args := pgx.NamedArgs{
		"id":         uuid.NewString(),
		"userId":     userId,
		"localId":    payload.LocalId,
		"categoryId": payload.CategoryId,
		"amount":     payload.Amount,
		"period":     payload.Period,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *BudgetService) Put(id string, payload models.Budget) error {
	query := `UPDATE budgets 
		SET user_id=@user_id, category_id=@categoryId, amount=@amount, period=@period
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":         id,
		"categoryId": payload.CategoryId,
		"amount":     payload.Amount,
		"period":     payload.Period,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *BudgetService) Delete(id string) error {
	query := `DELETE FROM budgets 
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
