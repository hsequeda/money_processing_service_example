package command_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hsequeda/money_processing_service_example/pkg/adapter/mocks"
	"github.com/hsequeda/money_processing_service_example/pkg/app/command"
	"github.com/hsequeda/money_processing_service_example/pkg/domain"
	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestAddAccountToClient(t *testing.T) {
	t.Parallel()
	t.Run("Ok case", func(t *testing.T) {
		t.Parallel()
		cmd := command.AddAccountToClient{
			AccountID: uuid.New().String(),
			Currency:  domain.CurrencyUSD,
			ClientID:  uuid.New().String(),
		}
		account, err := domain.NewAccount(cmd.AccountID, cmd.Currency, cmd.ClientID)
		require.NoError(t, err)
		ctrl := gomock.NewController(t)
		accountRepository := mocks.NewMockAccountRepository(ctrl)
		accountRepository.EXPECT().Save(gomock.Any(), gomock.Eq(account)).Return(nil)
		handler := command.NewAddAccountToClientHandler(accountRepository)
		err = handler.Handle(context.Background(), cmd)
		require.NoError(t, err)
	})

	t.Run("Fail: ClientID not found", func(t *testing.T) {
		t.Parallel()
		cmd := command.AddAccountToClient{
			AccountID: uuid.New().String(),
			Currency:  domain.CurrencyUSD,
			ClientID:  uuid.New().String(),
		}
		ctrl := gomock.NewController(t)
		accountRepository := mocks.NewMockAccountRepository(ctrl)
		accountRepository.EXPECT().Save(gomock.Any(), gomock.Any()).Return(domain.ErrClientIDNotFound)
		handler := command.NewAddAccountToClientHandler(accountRepository)
		err := handler.Handle(context.Background(), cmd)
		require.EqualError(t, err, domain.ErrClientIDNotFound.Error())
	})
}
