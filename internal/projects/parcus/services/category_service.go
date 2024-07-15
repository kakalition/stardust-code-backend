package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryService struct {
	DbPool *pgxpool.Pool
}

func (c *CategoryService) Get(userId string) ([]models.Category, error) {
	query := `SELECT * 
		FROM categories
		WHERE user_id = @userId
		ORDER BY id ASC
	;`

	args := pgx.NamedArgs{
		"userId": userId,
	}

	output := []models.Category{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query, args)
	if err != nil {
		return []models.Category{}, err
	}

	return output, nil
}

func (c *CategoryService) Post(userId string, payload models.Category) error {
	query := `INSERT INTO 
		categories (id, local_id, user_id, emoji, name, category_type) 
		VALUES (@id, @localId, @userId, @emoji, @name, @categoryType)
	;`

	args := pgx.NamedArgs{
		"id":           uuid.NewString(),
		"localId":      payload.LocalId,
		"name":         payload.Name,
		"emoji":        payload.Emoji,
		"categoryType": payload.CategoryType,
		"userId":       userId,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryService) Put(id string, payload models.Category) error {
	query := `UPDATE categories 
		SET emoji=@emoji, name=@name, category_type=@categoryType
		WHERE id=@id
	;`

	args := pgx.NamedArgs{
		"id":           id,
		"name":         payload.Name,
		"emoji":        payload.Emoji,
		"categoryType": payload.CategoryType,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryService) Delete(id string) error {
	query := `DELETE FROM categories 
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
