package service_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/stretchr/testify/assert"
)

func TestAdminCreateCategory(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `categories`")
	req := requests.CategoryRequest{
		Name: "Category Name",
	}
	t.Run("should create category", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := adminService.CreateCategory(context.Background(), &req)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, res.Name, "Category Name")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		res, err := adminService.CreateCategory(context.Background(), &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
}

func TestAdminUpdateCategory(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `categories`")
	req := requests.CategoryRequest{
		Name: "Category Name",
	}
	t.Run("should update category", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		mock.ExpectQuery(query).WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `categories`")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := adminService.UpdateCategory(context.Background(), 1, &req)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, res.Name, "Category Name")
	})
	t.Run("should handle failed query get by ID", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		res, err := adminService.UpdateCategory(context.Background(), 1, &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query update", func(t *testing.T) {
		errMock := errors.New("failed query")
		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		mock.ExpectQuery(query).WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `categories`")).WillReturnError(errMock)
		mock.ExpectRollback()

		res, err := adminService.UpdateCategory(context.Background(), 1, &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
}
