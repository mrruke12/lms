package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mrruke12/lms/internal/infrastructure/config"
)

func Connect(ctx context.Context, cfg *config.DBConfig) *pgxpool.Pool {
	_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.SSL,
	)

	connConfig, err := pgxpool.ParseConfig(connString)

	if err != nil {
		panic(err)
	}

	pool, err := pgxpool.NewWithConfig(
		_ctx,
		connConfig,
	)

	if err != nil {
		panic(err)
	}

	return pool
}

func Health(ctx context.Context, pool *pgxpool.Pool) error {
	_ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	return pool.Ping(_ctx)
}
