package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AccountService struct {
	DbPool *pgxpool.Pool
}

func (c *AccountService) Get(userId string) ([]models.Account, error) {
	query := `SELECT * 
		FROM accounts
		WHERE user_id = @userId
		ORDER BY id ASC
	;`

	args := pgx.NamedArgs{
		"userId": userId,
	}

	output := []models.Account{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query, args)
	if err != nil {
		return []models.Account{}, err
	}

	return output, nil
}

func (c *AccountService) Post(userId string, payload models.Account) error {
	query := `INSERT INTO 
		accounts (id, local_id, user_id, emoji, name, balance, created_date) 
		VALUES (@id, @localId, @userId, @emoji, @name, @balance, @createdDate)
	;`

	args := pgx.NamedArgs{
		"id":          uuid.NewString(),
		"userId":      userId,
		"localId":     payload.LocalId,
		"emoji":       payload.Emoji,
		"name":        payload.Name,
		"balance":     payload.Balance,
		"createdDate": payload.CreatedDate.Time,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *AccountService) Put(id string, payload models.Account) error {
	query := `UPDATE accounts 
		SET emoji=@emoji, name=@name, balance=@balance, updated_date=@updatedDate
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":          id,
		"name":        payload.Name,
		"emoji":       payload.Emoji,
		"balance":     payload.Balance,
		"updatedDate": payload.UpdatedDate.Time,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *AccountService) Delete(id string) error {
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
