package controllers

import (
	"Tribal/app/models"
	"Tribal/app/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type creditController struct {
}

type CreditController interface {
	Validate(ctx echo.Context) error
}

func NewCreditController() CreditController {
	return &creditController{}
}

func (c *creditController) Validate(ctx echo.Context) (err error) {
	customer := &models.CreditLineRequest{}
	if err = ctx.Bind(customer); err != nil {
		return
	}

	customer.RequestedDateParsed, _ = time.Parse(time.RFC3339, customer.RequestedDate)
	creditLineService := services.NewCreditLineService(customer, ctx)
	err = creditLineService.CheckTimeForRequest()
	
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, customer)
}
