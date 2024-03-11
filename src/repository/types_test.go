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

func TestTypesGetAll(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `types`")
	t.Run("should return all types", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")
		mock.ExpectQuery(query).WillReturnRows(rows)

		repo := repository.NewTypesRepository(c)
		types, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Len(t, types, 1)
		assert.Equal(t, types[0].ID, uint(1))
		assert.Equal(t, types[0].Name, "John Doe")
		assert.Equal(t, types[0].Description, "Test")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "name", "description"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		repo := repository.NewTypesRepository(c)
		types, err := repo.GetAll(context.Background())
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, types, 0)
	})
}

func TestTypesGetByID(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `types`")
	t.Run("should return type", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")
		mock.ExpectQuery(query).WillReturnRows(rows)

		repo := repository.NewTypesRepository(c)
		typeData, err := repo.GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, typeData.ID, uint(1))
		assert.Equal(t, typeData.Name, "John Doe")
		assert.Equal(t, typeData.Description, "Test")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "name", "description"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(`SELECT`).WillReturnError(errMock)

		repo := repository.NewTypesRepository(c)
		typeData, err := repo.GetByID(context.Background(), 1)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, typeData.ID, uint(0))
	})
}

func TestTypesCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `types`")
	typeData := model.Type{
		Name:        "John Doe",
		Description: "Test",
	}
	t.Run("should create type", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewTypesRepository(c)
		err := repo.Create(context.Background(), &typeData)
		assert.NoError(t, err)
		assert.Equal(t, typeData.ID, uint(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewTypesRepository(c)
		err := repo.Create(context.Background(), &typeData)
		assert.ErrorIs(t, err, errMock)
	})
}

func TestTypesUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE `types`")
	typeData := model.Type{
		ID:          1,
		Name:        "John Doe",
		Description: "Test",
	}
	t.Run("should update type", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewTypesRepository(c)
		err := repo.Update(context.Background(), &typeData)
		assert.NoError(t, err)
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewTypesRepository(c)
		err := repo.Update(context.Background(), &typeData)
		assert.ErrorIs(t, err, errMock)
	})
}
