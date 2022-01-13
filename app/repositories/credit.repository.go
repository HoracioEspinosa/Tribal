package repositories

import (
	"Tribal/app/helpers"
	"Tribal/app/models"
	"time"
)

type creditLineRepository struct {
	creditLineRequest *models.CreditLineRequest
}

type CreditLineRepository interface {
	GetCreditLinesByLineAndValidStatus(creditLineRequested float32, valid bool) []models.CreditLine
	GetLinesFromCustomDate(date time.Time, checkFails bool) []models.CreditLine
}

func NewCreditLineRepository(creditLineRequest *models.CreditLineRequest) *creditLineRepository {
	return &creditLineRepository{
		creditLineRequest: creditLineRequest,
	}
}

func (c creditLineRepository) GetCreditLinesByLineAndValidStatus(creditLineRequested float32, valid bool) []models.CreditLine {
	var result []models.CreditLine
	query := helpers.DB.Debug().Table("credit_line")
	query.Where("requested_credit_line", creditLineRequested)
	query.Where("valid", valid)
	query.Find(&result)

	return result
}

func (c creditLineRepository) GetLinesFromCustomDate(date time.Time, checkFails bool) []models.CreditLine {
	var result []models.CreditLine
	query := helpers.DB.Debug().Table("credit_line")

	if checkFails {
		query.Where("fail_times >= 3")
	} else {
		query.Where("requested_date >= ?", date)
	}

	query.Find(&result)

	return result
}
