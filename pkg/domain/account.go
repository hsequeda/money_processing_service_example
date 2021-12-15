package domain

import "errors"

// Account represent one of the bank account of the client.
type Account struct {
	id       string
	balance  float64
	currency Currency
	clientID string
}

// NewAccount create a new account with 0 balance.
func NewAccount(id string, currency Currency, clientID string) (*Account, error) {
	if currency.IsZero() {
		return nil, errors.New("account currency is empty")
	}

	if id == "" {
		return nil, errors.New("account ID is empty")
	}

	if clientID == "" {
		return nil, errors.New("client ID is empty")
	}

	return &Account{id: id, balance: 0, currency: currency, clientID: clientID}, nil
}

func (a Account) ID() string {
	return a.id
}

func (a Account) Balance() float64 {
	return a.balance
}

func (a Account) Currency() Currency {
	return a.currency
}

func (a Account) ClientID() string {
	return a.clientID
}

func (a Account) IsZero() bool {
	return a == Account{}
}
