package domain

import "errors"

// Currency represent an account currency (COP, USD, MXN)
type Currency struct {
	s string
}

var (
	CurrencyUSD = Currency{"usd"}
	CurrencyCOP = Currency{"cop"}
	CurrencyMXN = Currency{"mxn"}
)

func NewCurrencyFromStr(s string) (Currency, error) {
	switch s {
	case "usd":
		return CurrencyUSD, nil
	case "cop":
		return CurrencyCOP, nil
	case "mxn":
		return CurrencyMXN, nil
	default:
		return Currency{}, errors.New("unexpected currency str")
	}
}

func (c Currency) Value() string {
	return c.s
}
func (c Currency) IsZero() bool {
	return c.s == ""
}
