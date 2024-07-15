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
)

type RecurringTransactionController struct {
	Service   *services.RecurringTransactionService
	Validator *validator.Validate
}

func (c *RecurringTransactionController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value(types.SessionUserKey).(*models.NetworkUser)

	rawOutput, err := c.Service.Get(*user.Id)
	transformedRawOutput := utils.Map(rawOutput, func(local models.RecurringTransaction) models.NetworkRecurringTransaction {
		return local.ToNetwork()
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(transformedRawOutput)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data))
}

func (c *RecurringTransactionController) Post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user := ctx.Value(types.SessionUserKey).(*models.NetworkUser)

	var payload models.NetworkRecurringTransaction
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Service.Post(*user.Id, payload.ToInternal())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

func (c *RecurringTransactionController) Put(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var payload models.NetworkRecurringTransaction
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.Service.Put(id, payload.ToInternal())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *RecurringTransactionController) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	err := c.Service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
