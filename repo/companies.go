package repo

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type CompaniesRepo struct {
	Ctx       context.Context
	CompanyCache map[string]*models.Company
}

func (c *CompaniesRepo) GetCompanyByID(companyID string) (*models.Company, error) {
	if company, ok := c.CompanyCache[companyID]; ok {
		return company, nil
	}

	api := services.FleetsterAPI{Token: c.Ctx.Value("token").(string)}

	result, err := api.Get("/companies/" + companyID)
	if err != nil {
		return nil, errors.New("could not retrieve company with error " + err.Error())
	}

	var company *models.Company
	err = json.Unmarshal(result, &company)
	if err != nil {
		return nil, errors.New("could not parse company with error: " + err.Error())
	}

	c.CompanyCache[companyID] = company

	return company, err
}
