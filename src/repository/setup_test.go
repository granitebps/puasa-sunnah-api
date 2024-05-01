package repository_test

import (
	"log"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbMock      *gorm.DB
	mock        sqlmock.Sqlmock
	c           *core.Core
	typesRepo   *repository.TypesRepository
	sourceRepo  *repository.SourceRepository
	catRepo     *repository.CategoryRepository
	fastingRepo *repository.FastingRepository
)

func TestMain(m *testing.M) {
	newMockDB()
	c = &core.Core{
		Database: &core.Database{
			Db: dbMock,
		},
	}
	typesRepo = repository.NewTypesRepository(c)
	sourceRepo = repository.NewSourceRepository(c)
	catRepo = repository.NewCategoryRepository(c)
	fastingRepo = repository.NewFastingRepository(c)

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
