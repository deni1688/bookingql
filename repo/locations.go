package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type LocationsRepo struct {
	Ctx          context.Context
	LocationCache map[string]*models.Location
}

func (l *LocationsRepo) GetLocationByID(locationID string) (*models.Location, error) {
	if vehicle, ok := l.LocationCache[locationID]; ok {
		return vehicle, nil
	}

	api := services.FleetsterAPI{Token: l.Ctx.Value("token").(string)}

	result, err := api.Get("/locations/" + locationID)
	if err != nil {
		return nil, errors.New("could not retrieve location with error " + err.Error())
	}

	var vehicle *models.Location
	err = json.Unmarshal(result, &vehicle)
	if err != nil {
		return nil, errors.New("could not parse location with error: " + err.Error())
	}

	l.LocationCache[locationID] = vehicle

	return vehicle, err
}
