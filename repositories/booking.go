package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
	"net/url"
)

type BookingRepo struct {
	Ctx context.Context
}

func (r *BookingRepo) GetBookings(params *models.BookingParams) ([]*models.Booking, error) {
	api := services.FleetsterAPI{Token: r.Ctx.Value("token").(string)}

	query := buildBookingQuery(params)

	var bookings []*models.Booking
	err := api.Get("/bookings?"+query, &bookings)
	if err != nil {
		return nil, errors.New("could not retrieve bookings with error: " + err.Error())
	}

	return bookings, nil
}

func buildBookingQuery(params *models.BookingParams) string {
	q := map[string]string{
		"limit": params.Limit,
		"page":  params.Page,
	}

	p := url.Values{}
	for k, v := range q {
		p.Add(k, v)
	}

	return p.Encode() + hydrate([]string{"companyId", "userId", "vehicleId"})
}

func hydrate(entities []string) string {
	var hydrateQueryStr string

	for _, v := range entities {
		hydrateQueryStr += fmt.Sprintf("&hydrate[%s]=true", v)
	}

	return hydrateQueryStr
}
