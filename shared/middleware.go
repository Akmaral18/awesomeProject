package shared

import (
	"context"
	"github.com/rs/cors"
	"net/http"
)

func JsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func CorsPolicy(next http.Handler) http.Handler {
	corsConfig := cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions, http.MethodDelete},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:         3600,
	}

	return cors.New(corsConfig).Handler(next)
}

func AuthorizationRequired(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, ok := GetUserIdFromAuthHeader(r.Header.Get("Authorization"))

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		requestContext := context.WithValue(r.Context(), TokenContextKey{}, TokenContextKey{UserId: userId})
		r = r.WithContext(requestContext)
		next.ServeHTTP(w, r)
	})
}
