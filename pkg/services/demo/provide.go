package demo

import (
	"github.com/opoccomaxao/installable-bot-concept/pkg/clients/db"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/demo/repo"
	"github.com/samber/do"
)

func Provide(
	i *do.Injector,
) {
	do.ProvideNamed(i, "services/demo", func(i *do.Injector) (*Service, error) {
		db, err := db.InvokeLocal(i)
		if err != nil {
			return nil, err
		}

		return New(
			repo.New(db),
		), nil
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.InvokeNamed[*Service](i, "services/demo")
}
