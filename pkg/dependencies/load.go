package dependencies

import (
	"github.com/opoccomaxao/installable-bot-concept/pkg/clients/db"
	"github.com/opoccomaxao/installable-bot-concept/pkg/config"
	"github.com/opoccomaxao/installable-bot-concept/pkg/endpoints"
	"github.com/opoccomaxao/installable-bot-concept/pkg/server"
	"github.com/opoccomaxao/installable-bot-concept/pkg/services/auth"
	"github.com/samber/do"
)

func Load() (*do.Injector, error) {
	injector := do.New()

	err := LoadInto(injector)
	if err != nil {
		return nil, err
	}

	return injector, nil
}

func LoadInto(i *do.Injector) error {
	config, err := config.Load()
	if err != nil {
		return err
	}

	server.Provide(i, config.Server)
	db.Provide(i, config.DBLocal)
	endpoints.Provide(i)
	auth.Provide(i)

	return nil
}
