//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/controller"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

func SetupDependencies(c *core.Core) *controller.ControllerStruct {
	wire.Build(
		repository.NewCategoryRepository,
		repository.NewSourceRepository,
		repository.NewTypesRepository,
		repository.NewFastingRepository,

		service.NewCategoryService,
		service.NewSourceService,
		service.NewTypesService,
		service.NewFastingService,
		service.NewAdminService,

		controller.NewController,
	)

	return &controller.ControllerStruct{}
}
