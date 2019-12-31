package resolvers

import (
	"context"
	"github.com/deni1688/bookingql/graphql"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/repositories"
)

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

func (r *queryResolver) Bookings(ctx context.Context, params *models.BookingParams) ([]*models.Booking, error) {
	return repositories.BookingRepo{Ctx: ctx}.GetBookings(params)
}
