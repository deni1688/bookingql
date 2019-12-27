package graphql

import (
	"context"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repositories"
)

type bookingResolver struct{ *Resolver }

func (r *bookingResolver) User(ctx context.Context, obj *models.Booking) (*models.User, error) {
	return r.Loaders.Users(ctx).Load(obj.UserID)
}

func (r *bookingResolver) Vehicle(ctx context.Context, obj *models.Booking) (*models.Vehicle, error) {
	return r.Loaders.Vehicles(ctx).Load(obj.VehicleID)
}

func (r *bookingResolver) Location(ctx context.Context, obj *models.Booking) (*models.Location, error) {
	return r.Loaders.Locations(ctx).Load(obj.LocationStartID)
}

func (r *bookingResolver) Company(ctx context.Context, obj *models.Booking) (*models.Company, error) {
	c := repositories.CompanyRepo{Ctx: ctx, CompanyCache: map[string]*models.Company{}}
	return c.GetCompanyByID(obj.CompanyID)
}

func (r *Resolver) Booking() BookingResolver {
	return &bookingResolver{r}
}
