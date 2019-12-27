package graphql

import (
	"context"
	"github.com/deni1688/bookingql/loaders"
	"net/http"
)

func DataloaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", token)

		userLoader := new(loaders.UserLoader)
		vehicleLoader := new(loaders.VehicleLoader)
		locationLoader := new(loaders.LocationLoader)

		ctx = context.WithValue(ctx, loaders.UserLoaderKey, userLoader.NewUserLoader(token))
		ctx = context.WithValue(ctx, loaders.VehicleLoaderKey, vehicleLoader.NewVehicleLoader(token))
		ctx = context.WithValue(ctx, loaders.LocationLoaderKey, locationLoader.NewLocationLoader(token))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
