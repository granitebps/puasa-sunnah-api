package service_test

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbMock          *gorm.DB
	mock            sqlmock.Sqlmock
	c               *core.Core
	typesService    *service.TypesService
	sourceService   *service.SourceService
	categoryService *service.CategoryService
	fastingService  *service.FastingService
)

func TestMain(m *testing.M) {
	newMockDB()
	c = &core.Core{
		Database: &core.Database{
			Db: dbMock,
		},
	}

	typesRepo := repository.NewTypesRepository(c)
	sourceRepo := repository.NewSourceRepository(c)
	categoryRepo := repository.NewCategoryRepository(c)
	fastingRepo := repository.NewFastingRepository(c)

	typesService = service.NewTypesService(typesRepo)
	sourceService = service.NewSourceService(sourceRepo)
	categoryService = service.NewCategoryService(categoryRepo)
	fastingService = service.NewFastingService(fastingRepo, categoryRepo, typesRepo)

	os.Exit(m.Run())
}

func newMockDB() {
	db, m, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}

	dbMock = gormDB
	mock = m
}
