package graphql

import (
	"context"
	"net/http"
)

func DataloaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
