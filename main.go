package main

import (
	"log"

	"github.com/goava/di"

	"sc-wrap/app"
)

func main() {
	// create the application DI-container.
	c, err := di.New(
		di.Provide(app.NewApp),
		// provide the application's context.
		di.Provide(app.NewContext),
	)

	if err != nil {
		log.Fatalln(err)
	}

	// invoke application starter.
	if err = c.Invoke(app.Start); err != nil {
		log.Fatal(err)
	}
}
