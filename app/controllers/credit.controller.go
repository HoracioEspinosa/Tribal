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
	request := &models.CreditLineRequest{}

	if err = ctx.Bind(request); err != nil {
		return
	}

	parsedDate, _ := time.Parse(time.RFC3339, request.RequestedDate)

	customer := &models.CreditLine{
		CashBalance:         request.CashBalance,
		FailTimes:           0,
		FoundingTypes:       request.FoundingType,
		MonthlyRevenue:      request.MonthlyRevenue,
		RequestedCreditLine: request.RequestedCreditLine,
		RequestedDate:       parsedDate,
		Valid:               false,
	}

	creditLineService := services.NewCreditLineService(customer, ctx)
	err, code := creditLineService.CheckTimeForRequest()

	if err != nil {
		return ctx.JSON(code, map[string]interface{}{
			"status":  "ERROR",
			"message": err.Error(),
		})
	} else {
		err, code = creditLineService.CheckCreditLineForValidation()
		if err != nil {
			return ctx.JSON(code, map[string]interface{}{
				"status":  "ERROR",
				"message": err.Error(),
			})
		} else {
			err = nil
			code = http.StatusOK
		}
	}

	if err == nil {
		return ctx.JSON(code, map[string]interface{}{
			"status":  "OK",
			"message": "APPROVED",
		})
	}

	return err
}
