package service_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryGetAll(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `categories`")
	t.Run("should return all categories", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Test Category")
		mock.ExpectQuery(query).WillReturnRows(rows)

		categories, err := categoryService.GetAll(context.Background())
		assert.NoError(t, err)
		assert.Len(t, categories, 1)
		assert.Equal(t, categories[0].ID, uint(1))
		assert.Equal(t, categories[0].Name, "Test Category")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "name"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		categories, err := categoryService.GetAll(context.Background())
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, categories, 0)
	})
}
