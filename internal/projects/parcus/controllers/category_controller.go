package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"stardustcode/backend/internal/projects/parcus/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryController struct {
	DbPool *pgxpool.Pool
}

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	query := `SELECT * 
		FROM categories
		ORDER BY id ASC
	`

	output := []models.Category{}
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

func (c *CategoryController) Post(w http.ResponseWriter, r *http.Request) {
	var payload models.Category

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO categories (emoji, name, category_type) VALUES (@emoji, @name, @category_type);"
	args := pgx.NamedArgs{
		"name":          payload.Name,
		"user_id":       payload.UserId,
		"emoji":         payload.Emoji,
		"category_type": payload.CategoryType,
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

func (c *CategoryController) Put(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var payload models.Category

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE categories 
		SET user_id=@user_id, emoji=@emoji, name=@name, category_type=@category_type
		WHERE id=@id
	`

	args := pgx.NamedArgs{
		"id":            id,
		"user_id":       payload.UserId,
		"name":          payload.Name,
		"emoji":         payload.Emoji,
		"category_type": payload.CategoryType,
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

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	query := `DELETE FROM categories 
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
