package repositories

import (
	"context"
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

	var vehicle *models.Vehicle
	err := api.Get("/vehicles/"+vehicleID, &vehicle)
	if err != nil {
		return nil, errors.New("could not parse vehicle with error: " + err.Error())
	}

	r.VehicleCache[vehicleID] = vehicle

	return vehicle, err
}
