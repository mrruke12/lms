package postgres

import (
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mrruke12/lms/internal/application/apperr"
	"github.com/mrruke12/lms/internal/infrastructure/config"
)

func RunMigrations(cfg config.DBConfig) error {
	connString := ConnString(cfg)

	m, err := migrate.New(
		"file://internal/infrastructure/migrations",
		connString,
	)

	if err != nil {
		return apperr.MapDBErr(err)
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			slog.Info("migrations are up to date")
			return nil
		}

		return apperr.MapDBErr(err)
	}

	slog.Info("migrations applied successfuly")

	return nil
}
