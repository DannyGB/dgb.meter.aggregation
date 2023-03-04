//go:build wireinject
// +build wireinject

package main

import (
	"dgb/meter.aggregation/internal/application"
	"dgb/meter.aggregation/internal/configuration"

	"github.com/google/wire"
)

func CreateApi() *application.Api {

	panic(wire.Build(
		configuration.NewConfig,
		application.NewResponse,
		application.NewMiddleware,
		application.NewElecApi,
		application.NewOilApi,
		application.NewApi,
	))
}
