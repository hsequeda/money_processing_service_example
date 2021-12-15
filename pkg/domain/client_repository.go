package domain

import "context"

type ClientRepository interface {
	NewClient(ctx context.Context, clientId string) error
}
