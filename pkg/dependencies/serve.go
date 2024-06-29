package dependencies

import (
	"context"

	"github.com/opoccomaxao/installable-bot-concept/pkg/endpoints"
	"github.com/opoccomaxao/installable-bot-concept/pkg/server"
	"github.com/samber/do"
)

// Serve is non blocking function that starts all servables and returns immediately.
func Serve(
	_ context.Context,
	cancelCause func(error),
	i *do.Injector,
) error {
	err := endpoints.Invoke(i)
	if err != nil {
		return err
	}

	server, err := server.Invoke(i)
	if err != nil {
		return err
	}

	go func() {
		err := server.Serve()
		if err != nil {
			cancelCause(err)
		}
	}()

	return nil
}
