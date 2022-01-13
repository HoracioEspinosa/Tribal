package services

import (
	"Tribal/app/exceptions"
	"Tribal/app/models"
	"Tribal/app/repositories"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type creditLineService struct {
	creditLine *models.CreditLine
	context    echo.Context
	repository repositories.CreditLineRepository
}

type CreditLineService interface {
	CheckTimeForRequest() (err error)
	CheckCreditLineForValidation() (err error)
	CheckCreditStatus() []models.CreditLine
}

func NewCreditLineService(creditLine *models.CreditLine, ctx echo.Context) *creditLineService {
	repository := repositories.NewCreditLineRepository(creditLine)

	return &creditLineService{
		creditLine: creditLine,
		context:    ctx,
		repository: repository,
	}
}

func (c creditLineService) CheckTimeForRequest() (err error, code int) {
	now := time.Now()
	timestamp := now.Unix()
	code = http.StatusOK
	err = nil

	var resultsOn30Seconds = c.repository.GetLinesFromCustomDate(getDateFromTimestamp(timestamp-30), false)
	var resultsOn2Minutes = c.repository.GetLinesFromCustomDate(getDateFromTimestamp(timestamp-120), false)
	var has3FailedRequests = c.repository.GetLinesFromCustomDate(getDateFromTimestamp(timestamp), true)

	if len(has3FailedRequests) > 0 {
		err = &exceptions.CreditLineFailException{}
		code = http.StatusBadRequest
	}

	if (len(resultsOn2Minutes) >= 3) || (len(resultsOn30Seconds) > 0) {
		err = &exceptions.RequestLimitException{}
		code = http.StatusTooManyRequests
	}

	return err, code
}

func (c creditLineService) CheckCreditLineForValidation() (err error, code int) {
	monthlyValue := c.creditLine.MonthlyRevenue / 5
	cashValue := c.creditLine.CashBalance
	isValidCreditLine := false
	code = http.StatusOK

	switch c.creditLine.FoundingTypes {
	case "STARTUP":
		cashValue = cashValue / 3
		isValidCreditLine = (monthlyValue > 0) && (cashValue > 0)
		break
	case "SME":
		isValidCreditLine = monthlyValue > 0
		break
	default:
		isValidCreditLine = false
		err = &exceptions.FoundingTypeException{}
		code = http.StatusBadRequest
	}

	err, isValidCreditLine = c.CheckCreditStatus(c.creditLine, isValidCreditLine)

	if isValidCreditLine {
		code = http.StatusOK
		err = nil
	}

	if !isValidCreditLine || err != nil {
		err = &exceptions.RejectedException{}
		code = http.StatusNotAcceptable
	}

	return err, code
}

func (c creditLineService) CheckCreditStatus(creditLineRequest *models.CreditLine, validCreditLine bool) (err error, line bool) {
	creditLine := c.repository.GetCreditLinesByLineAndValidStatus(c.creditLine.RequestedCreditLine, true)

	if len(creditLine) <= 0 {
		creditLineList := c.repository.GetCreditLinesByLineAndValidStatus(c.creditLine.RequestedCreditLine, false)

		if len(creditLineList) >= 3 {
			err = &exceptions.CreditLineFailException{}
			validCreditLine = false
		} else {
			if len(creditLineList) > 0 {
				creditLineRequest.FailTimes = creditLineList[len(creditLineList)-1].FailTimes + 1
				creditLineRequest.Id = creditLineList[len(creditLineList)-1].Id
			} else {
				creditLineRequest.FailTimes = 1
			}

			creditLineRequest.Valid = validCreditLine

			if creditLineRequest.FailTimes < 4 {
				err = c.repository.InsertCreditLineItem(creditLineRequest)
			}
		}
	}

	return err, validCreditLine
}

func getDateFromTimestamp(timestamp int64) time.Time {
	i, _ := strconv.ParseInt(strconv.FormatInt(timestamp, 10), 10, 64)
	return time.Unix(i, 0)
}
