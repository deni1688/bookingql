package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type UsersRepo struct {
	Ctx       context.Context
	UserCache map[string]*models.User
}

func (u *UsersRepo) GetUserByID(userID string) (*models.User, error) {
	if user, ok := u.UserCache[userID]; ok {
		return user, nil
	}

	api := services.FleetsterAPI{Token: u.Ctx.Value("token").(string)}

	result, err := api.Get("/users/" + userID)
	if err != nil {
		return nil, errors.New("could not retrieve user with error " + err.Error())
	}

	var user *models.User
	err = json.Unmarshal(result, &user)
	if err != nil {
		return nil, errors.New("could not parse user with error: " + err.Error())
	}

	u.UserCache[userID] = user

	return user, err
}
