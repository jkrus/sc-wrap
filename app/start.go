package app

import (
	"context"
	"log"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"sc-wrap/config"
)

func startCommand(ctx context.Context, cfg *config.Config, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + app.Usage,
		Before: func(context *cli.Context) error {
			// load data from config file.
			if err := cfg.Load(); err != nil {
				return errors.Wrap(err, "load config")
			}

			return nil
		},
		After: func(c *cli.Context) error {
			c.Done()

			ctx.Done()

			log.Println("Application shutdown complete.")

			return nil
		},
	})
}
