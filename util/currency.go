package util

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUR"
	RUB = "RUB"
	CY  = "CY"
	JY  = "JY"
	KY  = "KY"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, EUR, RUB, CY, JY, KY:
		return true
	}
	return false
}
