package models

import (
	"time"
)

type (
	CreditLineRequest struct {
		FoundingType        string    `json:"foundingType" bson:"foundingType"`
		CashBalance         float32   `json:"cashBalance" bson:"cashBalance"`
		MonthlyRevenue      float32   `json:"monthlyRevenue" bson:"monthlyRevenue"`
		RequestedCreditLine float32   `json:"requestedCreditLine" bson:"requestedCreditLine"`
		RequestedDate       string    `json:"requestedDate" bson:"requestedDate"`
		RequestedDateParsed time.Time `json:"-" bson:"-"`
	}

	CreditLine struct {
		Id                  int32     `json:"id" bson:"id"`
		CashBalance         float32   `json:"cash_balance" bson:"cash_balance"`
		FailTimes           int64     `json:"fail_times" bson:"fail_times"`
		FoundingTypes       string    `json:"founding_types" bson:"founding_types"`
		MonthlyRevenue      float32   `json:"monthly_revenue" bson:"monthly_revenue"`
		RequestedCreditLine float32   `json:"requested_credit_line" bson:"requested_credit_line"`
		RequestedDate       time.Time `json:"requested_date" bson:"requested_date"`
		Valid               bool      `json:"valid" bson:"valid"`
	}
)
