package loaders

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
	"time"
)

const (
	UserLoaderKey     = "userLoader"
	VehicleLoaderKey  = "vehicleLoader"
	LocationLoaderKey = "locationLoader"
)

func (l *UserLoader) NewUserLoader(token string) *UserLoader {
	l.maxBatch = 100
	l.wait = 100 * time.Millisecond
	l.fetch = func(keys []string) ([]*models.User, []error) {
		api := services.FleetsterAPI{Token: token}
		query := api.BuildQuery(keys)

		results, err := api.Get("/users" + query)
		if err != nil {
			return nil, []error{errors.New("could not retrieve user with error " + err.Error())}
		}

		var users []*models.User
		err = json.Unmarshal(results, &users)
		if err != nil {
			return nil, []error{errors.New("could not parse user with error " + err.Error())}
		}

		userMap := map[string]*models.User{}
		for _, u := range users {
			userMap[u.ID] = u
		}

		for i, k := range keys {
			users[i] = userMap[k]
		}

		return users, nil
	}

	return l
}

func (l *LocationLoader) NewLocationLoader(token string) *LocationLoader {
	l.maxBatch = 100
	l.wait = 100 * time.Millisecond
	l.fetch = func(keys []string) ([]*models.Location, []error) {
		api := services.FleetsterAPI{Token: token}
		query := api.BuildQuery(keys)

		results, err := api.Get("/locations" + query)
		if err != nil {
			return nil, []error{errors.New("could not retrieve locations with error " + err.Error())}
		}

		var locations []*models.Location
		err = json.Unmarshal(results, &locations)
		if err != nil {
			return nil, []error{errors.New("could not parse locations with error " + err.Error())}
		}

		locationMap := map[string]*models.Location{}
		for _, lc := range locations {
			locationMap[lc.ID] = lc
		}

		for i, k := range keys {
			locations[i] = locationMap[k]
		}

		return locations, nil
	}

	return l
}

func (l *VehicleLoader) NewVehicleLoader(token string) *VehicleLoader {
	l.maxBatch = 100
	l.wait = 100 * time.Millisecond
	l.fetch = func(keys []string) ([]*models.Vehicle, []error) {
		api := services.FleetsterAPI{Token: token}
		query := api.BuildQuery(keys)

		results, err := api.Get("/vehicles" + query)
		if err != nil {
			return nil, []error{errors.New("could not retrieve vehicles with error " + err.Error())}
		}

		var vehicles []*models.Vehicle
		err = json.Unmarshal(results, &vehicles)
		if err != nil {
			return nil, []error{errors.New("could not parse vehicles with error " + err.Error())}
		}

		vehicleMap := map[string]*models.Vehicle{}
		for _, v := range vehicles {
			vehicleMap[v.ID] = v
		}

		for i, k := range keys {
			vehicles[i] = vehicleMap[k]
		}

		return vehicles, nil
	}

	return l
}

func Users(ctx context.Context) *UserLoader {
	return ctx.Value(UserLoaderKey).(*UserLoader)
}

func Vehicles(ctx context.Context) *VehicleLoader {
	return ctx.Value(VehicleLoaderKey).(*VehicleLoader)
}

func Locations(ctx context.Context) *LocationLoader {
	return ctx.Value(LocationLoaderKey).(*LocationLoader)
}
