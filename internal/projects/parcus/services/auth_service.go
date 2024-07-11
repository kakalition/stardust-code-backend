package services

import (
	"context"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct {
	DbPool *pgxpool.Pool
	Store  *sessions.CookieStore
}

func (c *AuthService) Post(payload models.User) error {
	currentUser := c.GetUserWithEmail(*payload.Email)

	if currentUser == nil {
		query := `INSERT INTO 
		users (email, display_name) 
		VALUES (@email, @displayName)
	;`

		args := pgx.NamedArgs{
			"email":       payload.Email,
			"displayName": payload.DisplayName,
		}

		_, err := c.DbPool.Exec(context.Background(), query, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *AuthService) GetUserWithEmail(email string) *models.User {
	query := `SELECT * 
		FROM users
		WHERE email = ?
		LIMIT 1
	;`

	output := models.User{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		return nil
	}

	return &output
}
