package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type BookingsRepo struct {
	Ctx context.Context
}

func (b *BookingsRepo) GetBookings(params *models.BookingParams) ([]*models.Booking, error) {
	api := services.FleetsterAPI{Token: b.Ctx.Value("token").(string)}
	query := fmt.Sprintf("/bookings?limit=%s&page=%s&sort[%s]=%d", params.Limit, params.Page, params.Sort, -1)
	result, err := api.Get(query)
	if err != nil {
		return nil, errors.New("could not retrieve bookings with error " + err.Error())
	}

	var bookings []*models.Booking
	err = json.Unmarshal(result, &bookings)
	if err != nil {
		return nil, errors.New("could not parse bookings with error: " + err.Error())
	}

	return bookings, nil
}
