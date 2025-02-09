package controllers

import (
	"encoding/json"
	"net/http"

	"stardustcode/backend/internal/projects/parcus/models"
	"stardustcode/backend/internal/projects/parcus/services"
	"stardustcode/backend/internal/types"
	"stardustcode/backend/internal/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
)

type CategoryController struct {
	Service   *services.CategoryService
	Validator *validator.Validate
	Store     sessions.Store
}

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
	user := *r.Context().Value(types.SessionUserKey).(*models.NetworkUser)
	rawOutput, err := c.Service.Get(*user.Id)
	transformedRawOutput := utils.Map(rawOutput, func(local models.Category) models.NetworkCategory {
		return local.ToNetwork()
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	output, err := json.Marshal(transformedRawOutput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func (c *CategoryController) Post(w http.ResponseWriter, r *http.Request) {
	var payload models.NetworkCategory
	user := *r.Context().Value(types.SessionUserKey).(*models.NetworkUser)

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Validator.Struct(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.Post(*user.Id, payload.ToInternal()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *CategoryController) Put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var payload models.NetworkCategory

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Service.Put(id, payload.ToInternal()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.Service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
