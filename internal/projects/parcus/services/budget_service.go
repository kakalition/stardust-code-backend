package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BudgetService struct {
	DbPool *pgxpool.Pool
}

func (c *BudgetService) Get() ([]models.Budget, error) {
	query := `SELECT * 
		FROM budgets
		ORDER BY id ASC
	;`

	output := []models.Budget{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		return []models.Budget{}, err
	}

	return output, nil
}

func (c *BudgetService) Post(payload models.Budget) error {
	query := `INSERT INTO 
		budgets (user_id, category_id, amount, balance, period) 
		VALUES (@userId, @categoryId, @amount, @period)
	;`

	args := pgx.NamedArgs{
		"userId":     payload.UserId,
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

func (c *BudgetService) Put(id int, payload models.Budget) error {
	query := `UPDATE budgets 
		SET user_id=@user_id, category_id=@categoryId, amount=@amount, period=@period
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":         id,
		"userId":     payload.UserId,
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

func (c *BudgetService) Delete(id int) error {
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
