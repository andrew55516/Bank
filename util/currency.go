package util

const (
	RUB = "RUB"
	USD = "USD"
	EUR = "EUR"
)

// IsSupportedCurrency checks if currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case RUB, USD, EUR:
		return true
	}
	return false
}
