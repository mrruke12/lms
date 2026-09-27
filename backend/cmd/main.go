package main

import (
	"log/slog"

	"github.com/mrruke12/lms/internal/infrastructure/config"
	"github.com/mrruke12/lms/internal/infrastructure/postgres"
)

func main() {
	config.LoadEnvironment()

	cfg, _ := config.GetDBConfig()

	err := postgres.RunMigrations(*cfg)

	if err != nil {
		slog.Error(err.Error())
	}
}
