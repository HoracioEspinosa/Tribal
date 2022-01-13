package exceptions

type RequestLimitException struct{}
type CreditLineFailException struct{}
type FoundingTypeException struct{}
type RejectedException struct{}

func (m *RequestLimitException) Error() string {
	return "Number of successful attempts"
}

func (m *CreditLineFailException) Error() string {
	return "A sales agent will contact you"
}

func (m *FoundingTypeException) Error() string {
	return "FoundingType not found"
}

func (m *RejectedException) Error() string {
	return "REJECTED"
}
