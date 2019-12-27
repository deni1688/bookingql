package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type BookingRepo struct {
	Ctx context.Context
}

func (r *BookingRepo) GetBookings(params *models.BookingParams) ([]*models.Booking, error) {
	api := services.FleetsterAPI{Token: r.Ctx.Value("token").(string)}
	query := fmt.Sprintf("/bookings?limit=%s&page=%s&sort[%s]=%d", params.Limit, params.Page, params.Sort, -1)

	var bookings []*models.Booking
	err := api.Get(query, &bookings)
	if err != nil {
		return nil, errors.New("could not parse bookings with error: " + err.Error())
	}

	return bookings, nil
}
