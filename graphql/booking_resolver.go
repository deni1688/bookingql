package graphql

import (
	"context"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repo"
)

type bookingResolver struct{ *Resolver }

func (r *bookingResolver) User(ctx context.Context, obj *models.Booking) (*models.User, error) {
	u := repo.UsersRepo{Ctx: ctx, UserCache: map[string]*models.User{}}
	return u.GetUserByID(obj.UserID)
}

func (r *bookingResolver) Vehicle(ctx context.Context, obj *models.Booking) (*models.Vehicle, error) {
	v := repo.VehiclesRepo{Ctx: ctx, VehicleCache: map[string]*models.Vehicle{}}
	return v.GetVehicleByID(obj.VehicleID)
}

func (r *Resolver) Booking() BookingResolver {
	return &bookingResolver{r}
}
