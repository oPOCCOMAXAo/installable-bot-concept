package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/opoccomaxao/installable-bot-concept/pkg/dependencies"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func run() error {
	appCtx, cancelCause := context.WithCancelCause(context.Background())
	defer cancelCause(nil)

	appCtx, _ = signal.NotifyContext(appCtx, os.Interrupt)

	deps, err := dependencies.Load()
	if err != nil {
		return err
	}

	defer func() {
		err := deps.Shutdown()
		if err != nil {
			log.Printf("failed to close dependencies: %+v", err)
		}
	}()

	err = dependencies.Serve(appCtx, cancelCause, deps)
	if err != nil {
		return err
	}

	<-appCtx.Done()

	return nil
}
