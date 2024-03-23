package service_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestTypesGetAll(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `types`")
	t.Run("should return all types", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")
		mock.ExpectQuery(query).WillReturnRows(rows)

		types, err := typesService.GetAll(context.Background())
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

		types, err := typesService.GetAll(context.Background())
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, types, 0)
	})
}
