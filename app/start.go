package app

import (
	"context"
	"log"

	"github.com/goava/di"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"sc-wrap/blockchain"
	"sc-wrap/config"
)

func startCommand(ctx context.Context, cfg *config.Config, app *cli.App, dic *di.Container) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + app.Usage,
		Before: func(context *cli.Context) error {
			// load data from config file.
			if err := cfg.Load(); err != nil {
				return errors.Wrap(err, "load config")
			}

			return provideServices(dic)
		},
		Action: func(ctx *cli.Context) error {
			return invokeServices(dic)
		},
		After: func(c *cli.Context) error {
			c.Done()

			ctx.Done()

			log.Println("Application shutdown complete.")

			return nil
		},
	})
}

// invokeServices tries to invoke required
// services from application DI container.
func invokeServices(dic *di.Container) error {
	// invoke quorum service starter.
	if err := dic.Invoke(blockchain.QuorumClientI.Start); err != nil {
		return errors.Wrap(err, "start quorum client")
	}

	return nil
}

// provideServices provides cli command specific
// services from application DI container.
func provideServices(dic *di.Container) error {
	// provide quorum service.
	if err := dic.Provide(blockchain.NewQuorumClient); err != nil {
		return errors.Wrap(err, "provide quorum client")
	}

	return nil
}
