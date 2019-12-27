package graphql

import (
	"context"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repositories"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Bookings(ctx context.Context, params *models.BookingParams) ([]*models.Booking, error) {
	b := repositories.BookingRepo{Ctx: ctx}
	return b.GetBookings(params)
}
