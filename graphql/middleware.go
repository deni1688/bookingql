package graphql

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deni1688/bookingql/dataloaders"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
	"net/http"
)

func TokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", token)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

const (
	userLoaderKey     = "userLoader"
	companyLoaderKey  = "companyLoader"
	vehicleLoaderKey  = "vehicleLoader"
	locationLoaderKey = "locationLoader"
)

func DataloaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		ctx := context.WithValue(r.Context(), "token", token)

		userLoader := dataloaders.UserLoader{
			Fetch: func(keys []string) ([]*models.User, []error) {
				api := services.FleetsterAPI{Token: token}
				query := buildQuery(keys)

				results, err := api.Get("/users" + query)
				if err != nil {
					return nil, []error{errors.New("could not retrieve user with error " + err.Error())}
				}

				var users []*models.User
				err = json.Unmarshal(results, &users)
				if err != nil {
					return nil, []error{errors.New("could not parse user with error " + err.Error())}
				}

				users = orderByKeys(users, keys)

				return users, nil
			},
		}

		ctx = context.WithValue(ctx, userLoaderKey, &userLoader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func orderByKeys(users []*models.User, keys []string) []*models.User {
	userMap := map[string]*models.User{}
	for _, u := range users {
		userMap[u.ID] = u
	}

	for i, k := range keys {
		users[i] = userMap[k]
	}
	return users
}

func buildQuery(keys []string) string {
	query := "?"

	for i, k := range keys {
		if i <= len(keys)-1 {
			query += "&"
		}
		query += fmt.Sprintf("_id[$in][%d]=%s", i, k)
	}
	return query
}
