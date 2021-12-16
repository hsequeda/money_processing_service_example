package adapter

import (
	"database/sql"
	"fmt"
)

func NewPostgresConnPool(dbName, dbUser, dbPass, dbHost, dbSSLMode string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("dbname=%s user=%s password=%s host=%s sslmode=%s", dbName, dbUser, dbPass, dbHost, dbSSLMode))
	if err != nil {
		return nil, fmt.Errorf("cannot connect to PostgreSql: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
