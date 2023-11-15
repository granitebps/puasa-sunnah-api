package controllers

import "github.com/granitebps/puasa-sunnah-api/services"

type Controller struct {
	SourceController *SourceController
}

func NewController(
	sourceService *services.SourceService,
) *Controller {
	return &Controller{
		SourceController: newSourceController(sourceService),
	}
}
