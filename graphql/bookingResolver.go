package graphql

import (
	"context"
	"github.com/deni1688/bookingql/dataloaders"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repo"
)

type bookingResolver struct{ *Resolver }

func (r *bookingResolver) User(ctx context.Context, obj *models.Booking) (*models.User, error) {
	return ctx.Value("userLoader").(*dataloaders.UserLoader).Load(obj.UserID)
}

func (r *bookingResolver) Vehicle(ctx context.Context, obj *models.Booking) (*models.Vehicle, error) {
	v := repo.VehicleRepo{Ctx: ctx, VehicleCache: map[string]*models.Vehicle{}}
	return v.GetVehicleByID(obj.VehicleID)
}

func (r *bookingResolver) Company(ctx context.Context, obj *models.Booking) (*models.Company, error) {
	c := repo.CompanyRepo{Ctx: ctx, CompanyCache: map[string]*models.Company{}}
	return c.GetCompanyByID(obj.CompanyID)
}

func (r *bookingResolver) Location(ctx context.Context, obj *models.Booking) (*models.Location, error) {
	c := repo.LocationRepo{Ctx: ctx, LocationCache: map[string]*models.Location{}}
	return c.GetLocationByID(obj.LocationStartID)
}

func (r *Resolver) Booking() BookingResolver {
	return &bookingResolver{r}
}
