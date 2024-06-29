package db

import (
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DSN string `env:"DSN,required"`
}

func OpenSQLite(config Config) (*gorm.DB, error) {
	res, err := gorm.Open(sqlite.Open(config.DSN), &gorm.Config{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return res, nil
}
