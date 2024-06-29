package auth

import (
	"context"

	"github.com/opoccomaxao/installable-bot-concept/pkg/clients/db"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth/repo"
	"github.com/samber/do"
)

func Provide(
	i *do.Injector,
) {
	do.ProvideNamed(i, "services/auth", func(i *do.Injector) (*Service, error) {
		db, err := db.InvokeLocal(i)
		if err != nil {
			return nil, err
		}

		return New(
			context.Background(),
			repo.New(db),
		)
	})
}

func Invoke(i *do.Injector) (*Service, error) {
	return do.InvokeNamed[*Service](i, "services/auth")
}
