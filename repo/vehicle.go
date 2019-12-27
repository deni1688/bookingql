package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type VehicleRepo struct {
	Ctx          context.Context
	VehicleCache map[string]*models.Vehicle
}

func (r *VehicleRepo) GetVehicleByID(vehicleID string) (*models.Vehicle, error) {
	if vehicle, ok := r.VehicleCache[vehicleID]; ok {
		return vehicle, nil
	}

	api := services.FleetsterAPI{Token: r.Ctx.Value("token").(string)}

	result, err := api.Get("/vehicles/" + vehicleID)
	if err != nil {
		return nil, errors.New("could not retrieve vehicle with error " + err.Error())
	}

	var vehicle *models.Vehicle
	err = json.Unmarshal(result, &vehicle)
	if err != nil {
		return nil, errors.New("could not parse vehicle with error: " + err.Error())
	}

	r.VehicleCache[vehicleID] = vehicle

	return vehicle, err
}
