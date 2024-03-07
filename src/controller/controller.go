package controller

import (
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

type ControllerStruct struct {
	Core               *core.Core
	SourceController   *SourceController
	TypesController    *TypesController
	CategoryController *CategoryController
	FastingController  *FastingController
	AdminController    *AdminController
}

func NewController(
	c *core.Core,
	sourceService *service.SourceService,
	typesService *service.TypesService,
	categoryService *service.CategoryService,
	fastingService *service.FastingService,
	adminService *service.AdminService,
) *ControllerStruct {
	return &ControllerStruct{
		SourceController:   newSourceController(sourceService),
		TypesController:    newTypesController(typesService),
		CategoryController: newCategoryController(categoryService),
		FastingController:  newFastingController(fastingService),
		AdminController:    newAdminController(c, adminService),
	}
}
