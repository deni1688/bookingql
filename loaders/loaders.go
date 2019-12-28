package loaders

import (
	"context"
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

type Loaders struct{}

func (l *Loaders) Users(ctx context.Context) *UserLoader {
	return ctx.Value(UserLoaderKey).(*UserLoader)
}

func (l *Loaders) Vehicles(ctx context.Context) *VehicleLoader {
	return ctx.Value(VehicleLoaderKey).(*VehicleLoader)
}

func (l *Loaders) Locations(ctx context.Context) *LocationLoader {
	return ctx.Value(LocationLoaderKey).(*LocationLoader)
}

func (l *UserLoader) NewUserLoader(token string) *UserLoader {
	l.maxBatch = 100
	l.wait = 1 * time.Millisecond
	l.fetch = func(keys []string) ([]*models.User, []error) {
		api := services.FleetsterAPI{Token: token}

		var users []*models.User
		err := api.GetKeys("/users", keys, &users)
		if err != nil {
			return nil, []error{errors.New("could not retrieve users with error " + err.Error())}
		}

		collectionMap := make(map[string]*models.User)
		for _, u := range users {
			collectionMap[u.ID] = u
		}

		for i, k := range keys {
			users[i] = collectionMap[k]
		}

		return users, nil
	}

	return l
}

func (l *LocationLoader) NewLocationLoader(token string) *LocationLoader {
	l.maxBatch = 100
	l.wait = 1 * time.Millisecond
	l.fetch = func(keys []string) ([]*models.Location, []error) {
		api := services.FleetsterAPI{Token: token}

		var locations []*models.Location
		err := api.GetKeys("/locations", keys, &locations)
		if err != nil {
			return nil, []error{errors.New("could not retrieve locations with error " + err.Error())}
		}

		collectionMap := make(map[string]*models.Location)
		for _, lc := range locations {
			collectionMap[lc.ID] = lc
		}

		for i, k := range keys {
			locations[i] = collectionMap[k]
		}

		return locations, nil
	}

	return l
}

func (l *VehicleLoader) NewVehicleLoader(token string) *VehicleLoader {
	l.maxBatch = 100
	l.wait = 1 * time.Millisecond
	l.fetch = func(keys []string) ([]*models.Vehicle, []error) {
		api := services.FleetsterAPI{Token: token}

		var vehicles []*models.Vehicle
		err := api.GetKeys("/vehicles", keys, &vehicles)
		if err != nil {
			return nil, []error{errors.New("could not retrieve vehicles with error " + err.Error())}
		}

		collectionMap := make(map[string]*models.Vehicle)
		for _, v := range vehicles {
			collectionMap[v.ID] = v
		}

		for i, k := range keys {
			vehicles[i] = collectionMap[k]
		}

		return vehicles, nil
	}

	return l
}
