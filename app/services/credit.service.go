package services

import (
	"Tribal/app/models"
	"Tribal/app/repositories"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

type creditLineService struct {
	creditLineRequest *models.CreditLineRequest
	context           echo.Context
	repository        repositories.CreditLineRepository
}

type CreditLineService interface {
	CheckTimeForRequest() (err error)
	CheckCreditLineForValidation() (err error)
	CheckCreditStatus() []models.CreditLine
}

func NewCreditLineService(creditLineRequest *models.CreditLineRequest, ctx echo.Context) *creditLineService {
	repository := repositories.NewCreditLineRepository(creditLineRequest)

	return &creditLineService{
		creditLineRequest: creditLineRequest,
		context:           ctx,
		repository:        repository,
	}
}

func (c creditLineService) CheckTimeForRequest() (err error) {
	now := time.Now()
	timestamp := now.Unix()
	var resultsOn30Seconds = c.repository.GetLinesFromCustomDate(getDateFromTimestamp(timestamp-30), false)
	var resultsOn2Minutes = c.repository.GetLinesFromCustomDate(getDateFromTimestamp(timestamp-120), false)
	var has3FailedRequests = c.repository.GetLinesFromCustomDate(getDateFromTimestamp(timestamp), true)

	if len(has3FailedRequests) > 0 {
		// TODO: Send message "A sales agent will contact you"
	}

	if len(resultsOn2Minutes) >= 3 {
		// TODO: Send error with code 429
	}

	if len(resultsOn30Seconds) > 0 {
		// TODO: Send error with code 429
	}

	fmt.Printf("Lineas de credito \n%v\n%v\n%v\n", has3FailedRequests, resultsOn2Minutes, resultsOn30Seconds)

	return nil
}

func (c creditLineService) CheckCreditLineForValidation() (err error) {
	monthlyValue := c.creditLineRequest.MonthlyRevenue / 5
	cashValue := c.creditLineRequest.CashBalance
	isValidCreditLine := false

	switch c.creditLineRequest.FoundingType {
	case "STARTUP":
		cashValue = cashValue / 3
		isValidCreditLine = (monthlyValue > 0) && (cashValue > 0)
		break
	case "SME":
		isValidCreditLine = monthlyValue > 0
		break
	default:
		isValidCreditLine = false
		// TODO: Print message "FoundingType not found" > 400
		break
	}

	if isValidCreditLine {
		// TODO: Print APPROVED message with 200 OK status code
	} else {
		// TODO: Print REJECTED message with 406 NOT_ACCEPTABLE status code
	}

	c.CheckCreditStatus(c.creditLineRequest.RequestedCreditLine)

	return nil
}

func (c creditLineService) CheckCreditStatus(requestedCreditLine float32) {
	creditLine := c.repository.GetCreditLinesByLineAndValidStatus(c.creditLineRequest.RequestedCreditLine, true)

	if len(creditLine) <= 0 {
		creditLineList := c.repository.GetCreditLinesByLineAndValidStatus(c.creditLineRequest.RequestedCreditLine, false)
		if len(creditLineList) >= 3 {
			// TODO: Generate an exception
		} else {
			// TODO: Create a new registry in credit_line table
		}
	}
}

func getDateFromTimestamp(timestamp int64) time.Time {
	i, _ := strconv.ParseInt(strconv.FormatInt(timestamp, 10), 10, 64)
	return time.Unix(i, 0)
}
