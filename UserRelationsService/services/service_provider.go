package services

import (
	"log"

	"github.com/sarulabs/di"
)

var Provider = buildServiceContainer()

const ()

var serviceContainer = []di.Def{}

func buildServiceContainer() di.Container {
	builder, err := di.NewBuilder()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = builder.Add(serviceContainer...)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return builder.Build()
}
