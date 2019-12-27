package graphql

import (
	"context"
	"github.com/deni1688/bookingql/loaders"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repo"
)

type bookingResolver struct{ *Resolver }

func (r *bookingResolver) Company(ctx context.Context, obj *models.Booking) (*models.Company, error) {
	c := repo.CompanyRepo{Ctx: ctx, CompanyCache: map[string]*models.Company{}}
	return c.GetCompanyByID(obj.CompanyID)
}

func (r *bookingResolver) User(ctx context.Context, obj *models.Booking) (*models.User, error) {
	return loaders.Users(ctx).Load(obj.UserID)
}

func (r *bookingResolver) Vehicle(ctx context.Context, obj *models.Booking) (*models.Vehicle, error) {
	return loaders.Vehicles(ctx).Load(obj.VehicleID)
}

func (r *bookingResolver) Location(ctx context.Context, obj *models.Booking) (*models.Location, error) {
	return loaders.Locations(ctx).Load(obj.LocationStartID)
}

func (r *Resolver) Booking() BookingResolver {
	return &bookingResolver{r}
}
