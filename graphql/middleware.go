package graphql

import (
	"context"
	"net/http"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "token", r.Header.Get("Authorization"))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
