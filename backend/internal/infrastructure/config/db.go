package config

import (
	"fmt"
	"strconv"

	"github.com/mrruke12/lms/internal/apperr"
	"github.com/mrruke12/lms/pkg/env"
)

// Basic Database Config
type DBConfig struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	SSL      string
}

func GetDBConfig() (*DBConfig, error) {
	values, err := env.KeysMustExist([]string{
		"DB_HOST",
		"DB_PORT",
		"DB_DATABASE",
		"DB_USER",
		"DB_PASSWORD",
		"DB_SSL",
	})

	if err != nil {
		return nil, apperr.EnvironmentVariableError(err.Error())
	}

	port, err := strconv.Atoi(values[1])

	if err != nil {
		return nil, apperr.EnvironmentVariableError(
			fmt.Sprintf("could not convert \"%s\" to int", values[1]),
		)
	}

	return &DBConfig{
		Host:     values[0],
		Port:     port,
		Database: values[2],
		User:     values[3],
		Password: values[4],
		SSL:      values[5],
	}, nil
}
