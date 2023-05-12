package delivery

import (
	"github.com/labstack/echo/v4"
	"self-payrol/helper"
	"self-payrol/model"
	"strconv"
)

type transactionDelivery struct {
	transactionUsecase model.TransactionUsecase
}

type TransactionDelivery interface {
	Mount(group *echo.Group)
}

func NewTransactionDelivery(transactionUsecase model.TransactionUsecase) TransactionDelivery {
	return &transactionDelivery{transactionUsecase: transactionUsecase}
}

func (p *transactionDelivery) Mount(group *echo.Group) {
	group.GET("", p.FetchTransactionHandler)
}

func (p *transactionDelivery) FetchTransactionHandler(c echo.Context) error {
	ctx := c.Request().Context()

	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")

	limitInt, _ := strconv.Atoi(limit)
	offsetInt, _ := strconv.Atoi(offset)

	transactions, i, err := p.transactionUsecase.Fetch(ctx, limitInt, offsetInt)
	if err != nil {
		return helper.ResponseErrorJson(c, i, err)
	}

	return helper.ResponseSuccessJson(c, "success", transactions)
}
