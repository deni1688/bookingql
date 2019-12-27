package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type LocationRepo struct {
	Ctx          context.Context
	LocationCache map[string]*models.Location
}

func (r *LocationRepo) GetLocationByID(locationID string) (*models.Location, error) {
	if vehicle, ok := r.LocationCache[locationID]; ok {
		return vehicle, nil
	}

	api := services.FleetsterAPI{Token: r.Ctx.Value("token").(string)}

	result, err := api.Get("/locations/" + locationID)
	if err != nil {
		return nil, errors.New("could not retrieve location with error " + err.Error())
	}

	var vehicle *models.Location
	err = json.Unmarshal(result, &vehicle)
	if err != nil {
		return nil, errors.New("could not parse location with error: " + err.Error())
	}

	r.LocationCache[locationID] = vehicle

	return vehicle, err
}
