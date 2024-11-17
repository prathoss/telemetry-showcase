package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/prathoss/telemetry_showcase/shared"
)

func UserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIdStr := r.Header.Get("X-User-ID")
		userID, err := uuid.Parse(userIdStr)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := shared.ContextWithUserId(r.Context(), userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
