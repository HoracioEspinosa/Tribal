package models

import (
	"time"
)

type (
	CreditLineRequest struct {
		FoundingType        string  `json:"foundingType"`
		CashBalance         float32 `json:"cashBalance"`
		MonthlyRevenue      float32 `json:"monthlyRevenue"`
		RequestedCreditLine float32 `json:"requestedCreditLine"`
		RequestedDate       string  `json:"requestedDate"`
	}

	CreditLine struct {
		Id                  int32     `json:"id"`
		CashBalance         float32   `json:"cash_balance"`
		FailTimes           int64     `json:"fail_times"`
		FoundingTypes       string    `json:"founding_types"`
		MonthlyRevenue      float32   `json:"monthly_revenue"`
		RequestedCreditLine float32   `json:"requested_credit_line"`
		RequestedDate       time.Time `json:"requested_date"`
		Valid               bool      `json:"valid"`
	}
)
