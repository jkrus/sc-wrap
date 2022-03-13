package app

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"sc-wrap/config"
)

// NewApp returns an application.
func NewApp() *cli.App {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"ver"},
		Usage:   "print the app version",
	}

	// construct cli application.
	app := &cli.App{
		Name:    config.AppName,
		Usage:   config.AppUsage,
		Version: version(),
		ExitErrHandler: func(context *cli.Context, err error) {
			if err != nil {
				log.Fatalln(err)
			}
		},
	}

	return app
}

// Start starts the application.
func Start(ctx context.Context, app *cli.App) error {
	return app.RunContext(ctx, os.Args)
}
