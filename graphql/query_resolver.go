package graphql

import (
	"context"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repo"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Bookings(ctx context.Context, filter *models.BookingFilter) ([]*models.Booking, error) {
	b := repo.BookingsRepo{Ctx: ctx}
	return b.GetBookings(filter)
}
