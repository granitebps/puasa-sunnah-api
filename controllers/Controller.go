package controllers

import "github.com/granitebps/puasa-sunnah-api/services"

type Controller struct {
	SourceController *SourceController
	TypesController  *TypesController
}

func NewController(
	sourceService *services.SourceService,
	typesService *services.TypesService,
) *Controller {
	return &Controller{
		SourceController: newSourceController(sourceService),
		TypesController:  newTypesController(typesService),
	}
}
