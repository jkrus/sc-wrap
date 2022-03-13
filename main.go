package main

import (
	"log"

	"github.com/goava/di"

	"sc-wrap/app"
	"sc-wrap/config"
)

func main() {
	// create the application DI-container.
	c, err := di.New(
		di.Provide(app.NewApp),
		// provide the application's context.
		di.Provide(app.NewContext),
		// provide the application config.
		di.Provide(config.NewConfig),
	)

	if err != nil {
		log.Fatalln(err)
	}

	// invoke application starter.
	if err = c.Invoke(app.Start); err != nil {
		log.Fatal(err)
	}
}
