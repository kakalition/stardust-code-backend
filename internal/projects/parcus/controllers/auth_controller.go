package controllers

import (
	"encoding/json"
	"net/http"

	"stardustcode/backend/internal/projects/parcus/models"
	"stardustcode/backend/internal/projects/parcus/services"
	"stardustcode/backend/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
)

type AuthController struct {
	Service   *services.AuthService
	Validator *validator.Validate
	Store     sessions.Store
}

func (c *AuthController) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, _ := ctx.Value(types.SessionUserKey).(*models.NetworkUser)

	output, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var payload *models.NetworkUser
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.Validator.Struct(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.Service.Login(payload.ToInternal())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := c.Store.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["id"] = user.Id
	session.Values["email"] = user.Email
	session.Values["displayName"] = user.DisplayName
	session.Values["lastSignedIn"] = user.LastSignedIn.Time

	if err = session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := c.Store.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Options.MaxAge = -1
	if err = session.Save(r, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
