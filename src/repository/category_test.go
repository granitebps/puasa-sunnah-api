package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/granitebps/puasa-sunnah-api/src/model"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/stretchr/testify/assert"
)

func TestCategoryGetAll(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `categories`")
	t.Run("should return all categories", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		mock.ExpectQuery(query).WillReturnRows(rows)

		repo := repository.NewCategoryRepository(c)
		sources, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Len(t, sources, 1)
		assert.Equal(t, sources[0].ID, uint(1))
		assert.Equal(t, sources[0].Name, "Category Name")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "name"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		repo := repository.NewCategoryRepository(c)
		sources, err := repo.GetAll(context.Background())
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, sources, 0)
	})
}

func TestCategoryGetByID(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `categories`")
	t.Run("should return category", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		mock.ExpectQuery(query).WillReturnRows(rows)

		repo := repository.NewCategoryRepository(c)
		source, err := repo.GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, source.ID, uint(1))
		assert.Equal(t, source.Name, "Category Name")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "name"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(`SELECT`).WillReturnError(errMock)

		repo := repository.NewCategoryRepository(c)
		cat, err := repo.GetByID(context.Background(), 1)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, cat.ID, uint(0))
	})
}

func TestCategoryCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `categories`")
	catData := model.Category{Name: "Category Name"}
	t.Run("should create category", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewCategoryRepository(c)
		err := repo.Create(context.Background(), &catData)
		assert.NoError(t, err)
		assert.Equal(t, catData.ID, uint(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewCategoryRepository(c)
		err := repo.Create(context.Background(), &catData)
		assert.ErrorIs(t, err, errMock)
	})
}

func TestCategoryUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE `categories`")
	catData := model.Category{
		ID:   1,
		Name: "Category Name",
	}
	t.Run("should update category", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewCategoryRepository(c)
		err := repo.Update(context.Background(), &catData)
		assert.NoError(t, err)
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewCategoryRepository(c)
		err := repo.Update(context.Background(), &catData)
		assert.ErrorIs(t, err, errMock)
	})
}
