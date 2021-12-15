package command

import (
	"context"

	"github.com/hsequeda/money_processing_service_example/pkg/domain"
)

type AddAccountToClient struct {
	AccountID string
	Currency  domain.Currency
	ClientID  string
}

type AddAccountToClientHandler struct {
	accountRepository domain.AccountRepository
}

func NewAddAccountToClientHandler(accountRepository domain.AccountRepository) AddAccountToClientHandler {
	return AddAccountToClientHandler{accountRepository}
}

func (h AddAccountToClientHandler) Handle(ctx context.Context, cmd AddAccountToClient) error {
	account, err := domain.NewAccount(cmd.AccountID, cmd.Currency, cmd.ClientID)
	if err != nil {
		return err
	}

	return h.accountRepository.Save(ctx, account)
}
