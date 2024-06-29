package db

import (
	"github.com/opoccomaxao/installable-bot-concept/pkg/migrations"
	"github.com/pkg/errors"
	"github.com/samber/do"
	"gorm.io/gorm"
)

func Provide(
	i *do.Injector,
	local Config,
) {
	do.ProvideNamed(i, "clients/db/local", func(i *do.Injector) (*gorm.DB, error) {
		res, err := OpenSQLite(local)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		err = migrations.Init(res)
		if err != nil {
			return nil, err
		}

		return res, nil
	})
}

func InvokeLocal(i *do.Injector) (*gorm.DB, error) {
	return do.InvokeNamed[*gorm.DB](i, "clients/db/local")
}
