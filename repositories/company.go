package repositories

import (
	"context"
	"errors"
	"github.com/deni1688/bookingql/models"
	"github.com/deni1688/bookingql/services"
)

type CompanyRepo struct {
	Ctx          context.Context
	CompanyCache map[string]*models.Company
}

func (r *CompanyRepo) GetCompanyByID(companyID string) (*models.Company, error) {
	if company, ok := r.CompanyCache[companyID]; ok {
		return company, nil
	}

	api := services.FleetsterAPI{Token: r.Ctx.Value("token").(string)}

	var company *models.Company
	err := api.Get("/companies/"+companyID, &company)
	if err != nil {
		return nil, errors.New("could not retrieve company with error: " + err.Error())
	}

	r.CompanyCache[companyID] = company

	return company, err
}
