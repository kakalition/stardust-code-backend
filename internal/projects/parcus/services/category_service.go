package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryService struct {
	DbPool *pgxpool.Pool
}

func (c *CategoryService) Get() ([]models.Category, error) {
	query := `SELECT * 
		FROM categories
		ORDER BY id ASC
	;`

	output := []models.Category{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		return []models.Category{}, err
	}

	return output, nil
}

func (c *CategoryService) Post(payload models.Category) error {
	query := `INSERT INTO 
		categories (user_id, emoji, name, category_type) 
		VALUES (@userId, @emoji, @name, @categoryType)
	;`

	args := pgx.NamedArgs{
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

func (c *CategoryService) Put(id int, payload models.Category) error {
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

func (c *CategoryService) Delete(id int) error {
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
