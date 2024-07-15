package middlewares

import (
	"net/http"
	"stardustcode/backend/internal/projects/parcus/models"
	"stardustcode/backend/internal/types"
)

func Authenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user, _ := ctx.Value(types.SessionUserKey).(*models.NetworkUser)

		if user == nil {
			w.WriteHeader(401)
		}

		next.ServeHTTP(w, r)
	})
}
