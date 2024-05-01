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

func TestAdminCreateSource(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `sources`")
	req := requests.SourceRequest{
		Url: "https://granitebps.com",
	}
	t.Run("should create source", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := adminService.CreateSource(context.Background(), &req)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, res.URL, "https://granitebps.com")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		res, err := adminService.CreateSource(context.Background(), &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
}

func TestAdminUpdateSource(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `sources`")
	req := requests.SourceRequest{
		Url: "https://granitebps.com",
	}
	t.Run("should update source", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "url"}).
			AddRow(1, "https://granitebps.com")
		mock.ExpectQuery(query).WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `sources`")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := adminService.UpdateSource(context.Background(), 1, &req)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, res.URL, "https://granitebps.com")
	})
	t.Run("should handle failed query get by ID", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		res, err := adminService.UpdateSource(context.Background(), 1, &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query update", func(t *testing.T) {
		errMock := errors.New("failed query")
		rows := sqlmock.NewRows([]string{"id", "url"}).
			AddRow(1, "https://granitebps.com")
		mock.ExpectQuery(query).WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `sources`")).WillReturnError(errMock)
		mock.ExpectRollback()

		res, err := adminService.UpdateSource(context.Background(), 1, &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
}

func TestAdminCreateType(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `types`")
	req := requests.TypeRequest{
		Name:        "Type Name",
		Description: "Type Description",
	}
	t.Run("should create type", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := adminService.CreateType(context.Background(), &req)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, res.Name, "Type Name")
		assert.Equal(t, res.Description, "Type Description")
	})
	t.Run("should handle failed query", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		res, err := adminService.CreateType(context.Background(), &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
}

func TestAdminUpdateType(t *testing.T) {
	query := regexp.QuoteMeta("SELECT * FROM `types`")
	req := requests.TypeRequest{
		Name:        "Type Name",
		Description: "Type Description",
	}
	t.Run("should update type", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "Type Name", "Type Description")
		mock.ExpectQuery(query).WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `types`")).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		res, err := adminService.UpdateType(context.Background(), 1, &req)
		assert.NoError(t, err)
		assert.Equal(t, res.Name, "Type Name")
		assert.Equal(t, res.Description, "Type Description")
	})
	t.Run("should handle failed query get by ID", func(t *testing.T) {
		errMock := errors.New("failed query")
		mock.ExpectQuery(query).WillReturnError(errMock)

		res, err := adminService.UpdateType(context.Background(), 1, &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query update", func(t *testing.T) {
		errMock := errors.New("failed query")
		rows := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "Type Name", "Type Description")
		mock.ExpectQuery(query).WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta("UPDATE `types`")).WillReturnError(errMock)
		mock.ExpectRollback()

		res, err := adminService.UpdateType(context.Background(), 1, &req)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, res.ID, uint(0))
	})
}
