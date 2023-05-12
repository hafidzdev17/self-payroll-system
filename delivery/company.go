package delivery

import (
	"net/http"
	"self-payrol/helper"
	"self-payrol/model"
	"self-payrol/request"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type companyDelivery struct {
	companyUsecase model.CompanyUsecase
}

type CompanyDelivery interface {
	Mount(group *echo.Group)
}

func NewCompanyDelivery(companyUsecase model.CompanyUsecase) CompanyDelivery {
	return &companyDelivery{companyUsecase: companyUsecase}
}

func (comp *companyDelivery) Mount(group *echo.Group) {

	// TODO:
	// 1. Buatlah handler yang mengarah ke fungsi comp.GetDetailCompanyHandler
	// 2. Buatlah handler yang mengarah ke fungsi comp.UpdateOrCreateCompanyHandler

	group.POST("/topup", comp.TopupBalanceHandler)

}

func (comp *companyDelivery) GetDetailCompanyHandler(e echo.Context) error {
	ctx := e.Request().Context()

	info, i, err := comp.companyUsecase.GetCompanyInfo(ctx)
	if err != nil {
		return helper.ResponseErrorJson(e, i, err)
	}

	return helper.ResponseSuccessJson(e, "success", info)

}

func (comp *companyDelivery) UpdateOrCreateCompanyHandler(e echo.Context) error {
	ctx := e.Request().Context()

	var req request.CompanyRequest

	if err := e.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(e, "Error binding struct", err.Error())
	}

	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(e, "Error validation", errVal)
	}

	company, i, err := comp.companyUsecase.CreateOrUpdateCompany(ctx, req)
	if err != nil {
		return helper.ResponseErrorJson(e, i, err)
	}

	return helper.ResponseSuccessJson(e, "success", company)
}

func (comp *companyDelivery) TopupBalanceHandler(e echo.Context) error {
	ctx := e.Request().Context()

	var req request.TopupCompanyBalance
	if err := e.Bind(&req); err != nil {
		return helper.ResponseErrorJson(e, http.StatusBadRequest, err)
	}

	if err := e.Validate(&req); err != nil {
		return helper.ResponseErrorJson(e, http.StatusBadRequest, err)
	}

	company, _, err := comp.companyUsecase.TopupBalance(ctx, req)
	if err != nil {
		return helper.ResponseErrorJson(e, http.StatusInternalServerError, err)
	}

	return helper.ResponseSuccessJson(e, "success", company)
}
