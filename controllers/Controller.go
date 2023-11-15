package controllers

import "github.com/granitebps/puasa-sunnah-api/services"

type Controller struct {
	SourceController   *SourceController
	TypesController    *TypesController
	CategoryController *CategoryController
}

func NewController(
	sourceService *services.SourceService,
	typesService *services.TypesService,
	categoryService *services.CategoryService,
) *Controller {
	return &Controller{
		SourceController:   newSourceController(sourceService),
		TypesController:    newTypesController(typesService),
		CategoryController: newCategoryController(categoryService),
	}
}
