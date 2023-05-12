package usecase

import (
	"context"
	"net/http"
	"self-payrol/model"
	"self-payrol/request"
)

type companyUsecase struct {
	companyRepo model.CompanyRepository
}

func NewCompanyUsecase(repo model.CompanyRepository) model.CompanyUsecase {
	return &companyUsecase{companyRepo: repo}
}

func (c *companyUsecase) GetCompanyInfo(ctx context.Context) (*model.Company, int, error) {
	company, err := c.companyRepo.Get(ctx)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	return company, http.StatusOK, err
}

func (c *companyUsecase) CreateOrUpdateCompany(ctx context.Context, req request.CompanyRequest) (*model.Company, int, error) {
	company, err := c.companyRepo.CreateOrUpdate(ctx, &model.Company{
		Name:    req.Name,
		Address: req.Address,
		Balance: req.Balance,
	})

	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	return company, http.StatusOK, nil

}

func (c *companyUsecase) TopupBalance(ctx context.Context, req request.TopupCompanyBalance) (*model.Company, int, error) {
	company, err := c.companyRepo.AddBalance(ctx, req.Balance)
	if err != nil {
		return nil, http.StatusUnprocessableEntity, err
	}

	return company, http.StatusOK, nil
}
