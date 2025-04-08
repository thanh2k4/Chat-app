package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PostgresConfig) (*pgxpool.Pool, error) {
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = cfg.Host
	}
	DATABASE_URL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", cfg.User, cfg.Password, dbHost, cfg.Port, cfg.DBName, cfg.SSLMode)

	pool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
