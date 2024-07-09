package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountService struct {
	DbPool *pgxpool.Pool
}

func (c *AccountService) Get() ([]models.Account, error) {
	query := `SELECT * 
		FROM accounts
		ORDER BY id ASC
	;`

	output := []models.Account{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		return []models.Account{}, err
	}

	return output, nil
}

func (c *AccountService) Post(payload models.Account) error {
	query := `INSERT INTO 
		accounts (user_id, emoji, name, balance, created_date) 
		VALUES (@userId, @emoji, @name, @balance, @createdDate)
	;`

	args := pgx.NamedArgs{
		"userId":      payload.UserId,
		"emoji":       payload.Emoji,
		"name":        payload.Name,
		"balance":     payload.Balance,
		"createdDate": payload.CreatedDate,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *AccountService) Put(id int, payload models.Account) error {
	query := `UPDATE accounts 
		SET user_id=@user_id, emoji=@emoji, name=@name, balance=@balance, updated_date=@updatedDate
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":          id,
		"userId":      payload.UserId,
		"name":        payload.Name,
		"emoji":       payload.Emoji,
		"balance":     payload.Balance,
		"updatedDate": payload.UpdatedDate,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *AccountService) Delete(id int) error {
	query := `DELETE FROM accounts 
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
