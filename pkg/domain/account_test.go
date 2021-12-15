package domain_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hsequeda/money_processing_service_example/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAccount(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name            string
		shouldFail      bool
		accountID       string
		accountCurrency domain.Currency
		clientID        string
		errorMsg        string
	}{
		{
			name:            "Ok: MXN",
			shouldFail:      false,
			accountID:       uuid.New().String(),
			accountCurrency: domain.CurrencyMXN,
			clientID:        uuid.New().String(),
			errorMsg:        "",
		},
		{
			name:            "Ok: COP",
			shouldFail:      false,
			accountID:       uuid.New().String(),
			accountCurrency: domain.CurrencyCOP,
			clientID:        uuid.New().String(),
			errorMsg:        "",
		},
		{
			name:            "Ok: USD",
			shouldFail:      false,
			accountID:       uuid.New().String(),
			accountCurrency: domain.CurrencyUSD,
			clientID:        uuid.New().String(),
			errorMsg:        "",
		},
		{
			name:            "Fail: currency empty",
			shouldFail:      true,
			accountID:       uuid.New().String(),
			accountCurrency: domain.Currency{},
			clientID:        uuid.New().String(),
			errorMsg:        "account currency is empty",
		},
		{
			name:            "Fail: ID empty",
			shouldFail:      true,
			accountID:       "",
			accountCurrency: domain.CurrencyCOP,
			clientID:        uuid.New().String(),
			errorMsg:        "account ID is empty",
		},
		{
			name:            "Fail: ClientID empty",
			shouldFail:      true,
			accountID:       uuid.New().String(),
			accountCurrency: domain.CurrencyCOP,
			clientID:        "",
			errorMsg:        "client ID is empty",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			account, err := domain.NewAccount(tc.accountID, tc.accountCurrency, tc.clientID)
			if tc.shouldFail {
				require.EqualError(t, err, tc.errorMsg)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, 0.0, account.Balance())
			assert.Equal(t, tc.accountID, account.ID())
			assert.Equal(t, tc.accountCurrency, account.Currency())
			assert.Equal(t, tc.clientID, account.ClientID())
			assert.False(t, account.IsZero())
		})
	}
}
