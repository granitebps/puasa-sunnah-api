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

func TestSourceGetAll(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `sources`")
	t.Run("should return all sources", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "url"}).
			AddRow(1, "https://example.com")
		mock.ExpectQuery(query).WillReturnRows(rows)

		repo := repository.NewSourceRepository(c)
		sources, err := repo.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Len(t, sources, 1)
		assert.Equal(t, sources[0].ID, uint(1))
		assert.Equal(t, sources[0].Url, "https://example.com")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "name", "description"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		repo := repository.NewSourceRepository(c)
		sources, err := repo.GetAll(context.Background())
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, sources, 0)
	})
}

func TestSourceGetByID(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `sources`")
	t.Run("should return source", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "url"}).
			AddRow(1, "https://example.com")
		mock.ExpectQuery(query).WillReturnRows(rows)

		repo := repository.NewSourceRepository(c)
		source, err := repo.GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, source.ID, uint(1))
		assert.Equal(t, source.Url, "https://example.com")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "url"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(`SELECT`).WillReturnError(errMock)

		repo := repository.NewSourceRepository(c)
		source, err := repo.GetByID(context.Background(), 1)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, source.ID, uint(0))
	})
}

func TestSourceCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `sources`")
	sourceData := model.Source{
		Url: "https://example.com",
	}
	t.Run("should create source", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewSourceRepository(c)
		err := repo.Create(context.Background(), &sourceData)
		assert.NoError(t, err)
		assert.Equal(t, sourceData.ID, uint(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewSourceRepository(c)
		err := repo.Create(context.Background(), &sourceData)
		assert.ErrorIs(t, err, errMock)
	})
}

func TestSourceUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE `sources`")
	sourceData := model.Source{
		ID:  1,
		Url: "https://example.com",
	}
	t.Run("should update source", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewSourceRepository(c)
		err := repo.Update(context.Background(), &sourceData)
		assert.NoError(t, err)
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewSourceRepository(c)
		err := repo.Update(context.Background(), &sourceData)
		assert.ErrorIs(t, err, errMock)
	})
}
