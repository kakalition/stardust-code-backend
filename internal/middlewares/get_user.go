package middlewares

import (
	"context"
	"net/http"
	"stardustcode/backend/internal/types"

	"github.com/gorilla/sessions"
)

func GetUser(store sessions.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, err := store.Get(r, "auth")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			ctx := r.Context()
			id, _ := session.Values["id"].(string)

			if id == "" {
				ctx = context.WithValue(ctx, types.SessionUserKey, nil)
			} else {
				session, err := store.Get(r, "auth")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				genericMap := types.GenericMap(session.Values)

				user := genericMap.GetUser()
				ctx = context.WithValue(ctx, types.SessionUserKey, user)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
