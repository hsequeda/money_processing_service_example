package adapter_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hsequeda/money_processing_service_example/pkg/adapter"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Parallel()
	repo := newClientRepositoryPsql(t)
	for i := 0; i < 100; i++ {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()
			repo.NewClient(ctx, uuid.New().String())
		})
	}
}

func newClientRepositoryPsql(t *testing.T) *adapter.ClientRepositoryPsql {
	t.Helper()
	dbName := os.Getenv("PSQL_DB_NAME_TEST")
	dbUser := os.Getenv("PSQL_DB_USER_TEST")
	dbPass := os.Getenv("PSQL_DB_PASS_TEST")
	dbHost := os.Getenv("PSQL_DB_HOST_TEST")
	dbSSLMode := os.Getenv("PSQL_DB_SSLMODE_TEST")
	conn, err := adapter.NewPostgresConnPool(dbName, dbUser, dbPass, dbHost, dbSSLMode)
	require.NoError(t, err)
	return adapter.NewClientRepositoryPsql(conn)
}
