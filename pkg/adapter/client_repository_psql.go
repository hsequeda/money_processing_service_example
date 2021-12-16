package adapter

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type ClientRepositoryPsql struct {
	conn *sql.DB
}

func NewClientRepositoryPsql(conn *sql.DB) *ClientRepositoryPsql {
	return &ClientRepositoryPsql{
		conn: conn,
	}
}

func (c *ClientRepositoryPsql) NewClient(ctx context.Context, clientId string) error {
	insertSmtp, err := c.conn.Prepare("insert into client(id) values($1)")
	if err != nil {
		return errors.Wrap(err, "unhandled error")
	}

	defer insertSmtp.Close()
	if _, err := insertSmtp.ExecContext(ctx, clientId); err != nil {
		return errors.Wrap(err, "unhandled error")
	}

	return nil
}
