package domain

import (
	"context"
	"errors"
)

var ErrClientIDNotFound = errors.New("account client ID doesn't exist")

type AccountRepository interface {
	Save(ctx context.Context, account *Account) error
	GetOne(ctx context.Context, accountID string) (*Account, error)
}
