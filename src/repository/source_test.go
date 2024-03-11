package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
