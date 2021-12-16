package domain

import (
	"context"
)

type TxService interface {
	Deposit(ctx context.Context, to string, amount uint) error
	Withdraw(ctx context.Context, from string, amount uint) error
	Transfer(ctx context.Context, from, to string, amount uint) error
}
