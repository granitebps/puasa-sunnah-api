package controllers

import "github.com/granitebps/puasa-sunnah-api/services"

type Controller struct {
	SourceController   *SourceController
	TypesController    *TypesController
	CategoryController *CategoryController
	FastingController  *FastingController
}

func NewController(
	sourceService *services.SourceService,
	typesService *services.TypesService,
	categoryService *services.CategoryService,
	fastingService *services.FastingService,
) *Controller {
	return &Controller{
		SourceController:   newSourceController(sourceService),
		TypesController:    newTypesController(typesService),
		CategoryController: newCategoryController(categoryService),
		FastingController:  newFastingController(fastingService),
	}
}
