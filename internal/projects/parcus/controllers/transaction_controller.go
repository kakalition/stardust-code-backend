package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TransactionController struct {
	DbPool *pgxpool.Pool
}

func (c *TransactionController) Get(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * 
		FROM transactions
		ORDER BY id ASC
	`

	output := []models.Transaction{}
	err := pgxscan.Select(context.Background(), c.DbPool, &output, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(output)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func (c *TransactionController) Post(w http.ResponseWriter, r *http.Request) {
	var payload models.Transaction

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO 
		categories (date, walletId, categoryId, amount, notes, type, is_recurring, created_date, updated_date)
		VALUES (@date, @walletId, @categoryId, @amount, @notes, @type, @isRecurring, @createdDate, @updatedDate)`

	args := pgx.NamedArgs{
		"date":        payload.Date,
		"walletId":    payload.WalletId,
		"categoryId":  payload.CategoryId,
		"amount":      payload.Amount,
		"notes":       payload.Notes,
		"type":        payload.Type,
		"isRecurring": payload.IsRecurring,
		"createdDate": time.Now().Format("YYYY-MM-DD HH:mm:ss"),
		"updatedDate": nil,
	}

	_, err = c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (c *TransactionController) Put(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	var payload models.Transaction

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE transactions
		SET date=@date, wallet_id=@walletId, category_id=@categoryId, amount=@amount, notes=@notes, type=@type, is_recurring=@isRecurring, updated_date=@updatedDate
		WHERE id=@id
	`

	args := pgx.NamedArgs{
		"id":          id,
		"date":        payload.Date,
		"walletId":    payload.WalletId,
		"categoryId":  payload.CategoryId,
		"amount":      payload.Amount,
		"notes":       payload.Notes,
		"type":        payload.Type,
		"isRecurring": payload.IsRecurring,
		"updatedDate": time.Now().Format("YYYY-MM-DD HH:mm:ss"),
	}

	_, err = c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func (c *TransactionController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	query := `DELETE FROM transactions 
		WHERE id=@id
	`

	args := pgx.NamedArgs{
		"id": id,
	}

	_, err := c.DbPool.Exec(context.Background(), query, args)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
