package database

import (
	"context"
	"digitalwallet-service/src/webapp/config"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var pgxPool *pgxpool.Pool

func Exec(query string, args ...any) (int64, error) {

	result, err := pgxPool.Exec(context.Background(), query, args...)
	if err != nil {
		return 0, err
	}

	rows := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rows == 0 {
		return 0, fmt.Errorf("Transactions was not inserted")
	}

	return rows, nil
}

func Query(query string, args ...any) (pgx.Rows, error) {
	rows, err := pgxPool.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetPgxPool() *pgxpool.Pool {
	if pgxPool == nil {
		return openPgxConnection()
	}

	if err := pgxPool.Ping(context.Background()); err != nil {
		pgxPool.Close()
		return openPgxConnection()
	}

	return pgxPool
}

func openPgxConnection() *pgxpool.Pool {

	dbConfig := config.DatabaseConfig
	var DATABASE_URL string = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbConfig.User, os.Getenv("DB_PASSWORD"), dbConfig.Host, dbConfig.Port, dbConfig.Database)

	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	pgxPool = dbpool

	return dbpool
}
