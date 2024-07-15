package services

import (
	"context"
	"time"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct {
	DbPool *pgxpool.Pool
}

func (c *AuthService) Login(payload models.User) (*models.User, error) {
	currentUser := c.GetUserWithEmail(*payload.Email)

	if currentUser == nil {
		query := `INSERT INTO 
			users (id, email, display_name) 
			VALUES (@id, @email, @displayName)
		;`

		args := pgx.NamedArgs{
			"id":          uuid.NewString(),
			"email":       payload.Email,
			"displayName": payload.DisplayName,
		}

		_, err := c.DbPool.Exec(context.Background(), query, args)
		if err != nil {
			return nil, err
		}

		currentUser = c.GetUserWithEmail(*payload.Email)
	}

	query := `UPDATE users
			SET last_signed_in = @lastSignedIn	
			WHERE id = @id
		;`

	currentTime := time.Now()

	args := pgx.NamedArgs{
		"id":           currentUser.Id,
		"lastSignedIn": currentTime,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		return nil, err
	}

	currentUser.LastSignedIn = pgtype.Timestamp{
		Time: currentTime,
	}

	return currentUser, nil
}

func (c *AuthService) GetUserWithEmail(email string) *models.User {
	query := `SELECT * 
		FROM users
		WHERE email = @email
		LIMIT 1
	;`

	args := pgx.NamedArgs{
		"email": email,
	}

	output := models.User{}
	rows, err := c.DbPool.Query(context.Background(), query, args)
	if err != nil {
		println("error: " + err.Error())
		return nil
	}

	if err := pgxscan.ScanOne(&output, rows); err != nil {
		println("error: " + err.Error())
		return nil
	}

	return &output
}
