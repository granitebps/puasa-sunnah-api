package controller

import "github.com/granitebps/puasa-sunnah-api/src/service"

type ControllerStruct struct {
	SourceController   *SourceController
	TypesController    *TypesController
	CategoryController *CategoryController
	FastingController  *FastingController
}

func NewController(
	sourceService *service.SourceService,
	typesService *service.TypesService,
	categoryService *service.CategoryService,
	fastingService *service.FastingService,
) *ControllerStruct {
	return &ControllerStruct{
		SourceController:   newSourceController(sourceService),
		TypesController:    newTypesController(typesService),
		CategoryController: newCategoryController(categoryService),
		FastingController:  newFastingController(fastingService),
	}
}
