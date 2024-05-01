package service_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
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

func TestAdminCreateFasting(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `fastings`")
	queryFasting := regexp.QuoteMeta("SELECT * FROM `fastings`")
	queryCat := regexp.QuoteMeta("SELECT * FROM `categories`")
	queryType := regexp.QuoteMeta("SELECT * FROM `types`")
	req := requests.FastingCreateUpdateRequest{
		CategoryID: 1,
		TypeID:     1,
		Date:       "2021-01-01",
		Year:       2021,
		Month:      1,
		Day:        1,
	}
	t.Run("should create fasting", func(t *testing.T) {
		now := time.Now()
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)

		res, err := adminService.CreateFasting(context.Background(), &req)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(1))
	})
	t.Run("should handle failed query: create", func(t *testing.T) {
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errors.New("failed query"))
		mock.ExpectRollback()

		res, err := adminService.CreateFasting(context.Background(), &req)
		assert.Equal(t, res.ID, uint(0))
		assert.Equal(t, err.Error(), "failed query")
	})
	t.Run("should handle failed query: get by ID", func(t *testing.T) {
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))

		res, err := adminService.CreateFasting(context.Background(), &req)
		assert.Equal(t, err.Error(), "not found")
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query: fasting exists", func(t *testing.T) {
		now := time.Now()
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)

		res, err := adminService.CreateFasting(context.Background(), &req)
		assert.Equal(t, err.Error(), "Fasting already exists")
		assert.Equal(t, merry.HTTPCode(err), fiber.StatusBadRequest)
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query: date format not valid", func(t *testing.T) {
		req.Date = "123456789"
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))

		res, err := adminService.CreateFasting(context.Background(), &req)
		assert.Equal(t, merry.UserMessage(err), "Invalid date format")
		assert.Equal(t, merry.HTTPCode(err), fiber.StatusBadRequest)
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query: date not valid", func(t *testing.T) {
		req.Date = "2021-01-01"
		req.Year = 1234
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))

		res, err := adminService.CreateFasting(context.Background(), &req)
		assert.Equal(t, err.Error(), "Invalid date")
		assert.Equal(t, merry.HTTPCode(err), fiber.StatusBadRequest)
		assert.Equal(t, res.ID, uint(0))
	})
}

func TestAdminUpdateFasting(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE `fastings`")
	queryFasting := regexp.QuoteMeta("SELECT * FROM `fastings`")
	queryCat := regexp.QuoteMeta("SELECT * FROM `categories`")
	queryType := regexp.QuoteMeta("SELECT * FROM `types`")
	req := requests.FastingCreateUpdateRequest{
		CategoryID: 1,
		TypeID:     1,
		Date:       "2021-01-01",
		Year:       2021,
		Month:      1,
		Day:        1,
	}
	t.Run("should update fasting", func(t *testing.T) {
		now := time.Now()
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01").
			AddRow(2, 1, 1, now, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)

		res, err := adminService.UpdateFasting(context.Background(), &req, 2)
		assert.NoError(t, err)
		assert.Equal(t, res.ID, uint(2))
	})
	t.Run("should handle failed query: update", func(t *testing.T) {
		now := time.Now()
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnError(errors.New("failed query"))
		mock.ExpectRollback()

		res, err := adminService.UpdateFasting(context.Background(), &req, 1)
		assert.Equal(t, res.ID, uint(0))
		assert.Equal(t, err.Error(), "failed query")
	})
	t.Run("should handle failed query: first get by ID", func(t *testing.T) {
		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))

		res, err := adminService.UpdateFasting(context.Background(), &req, 1)
		assert.Equal(t, err.Error(), "not found")
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query: second get by ID", func(t *testing.T) {
		now := time.Now()
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()
		mock.ExpectQuery(queryFasting).WillReturnError(errors.New("not found"))

		res, err := adminService.UpdateFasting(context.Background(), &req, 1)
		assert.Equal(t, err.Error(), "not found")
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query: fasting exists", func(t *testing.T) {
		now := time.Now()
		tom := time.Now().Add(time.Hour * 24)
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01").
			AddRow(2, 1, 1, tom, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)

		req.Date = "2021-01-02"
		res, err := adminService.UpdateFasting(context.Background(), &req, 1)
		assert.Equal(t, err.Error(), "Fasting already exists")
		assert.Equal(t, merry.HTTPCode(err), fiber.StatusBadRequest)
		assert.Equal(t, res.ID, uint(0))
	})
	t.Run("should handle failed query: date format not valid", func(t *testing.T) {
		req.Date = "123456789"
		now := time.Now()
		rowsFasting := sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"}).
			AddRow(1, 1, 1, now, "2021", "01", "01")
		rowsCat := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Category Name")
		rowsType := sqlmock.NewRows([]string{"id", "name", "description"}).
			AddRow(1, "John Doe", "Test")

		mock.MatchExpectationsInOrder(true)
		mock.ExpectQuery(queryFasting).WillReturnRows(rowsFasting)
		mock.ExpectQuery(queryCat).WillReturnRows(rowsCat)
		mock.ExpectQuery(queryType).WillReturnRows(rowsType)

		res, err := adminService.UpdateFasting(context.Background(), &req, 1)
		assert.Equal(t, merry.UserMessage(err), "Invalid date format")
		assert.Equal(t, merry.HTTPCode(err), fiber.StatusBadRequest)
		assert.Equal(t, res.ID, uint(0))
	})
}
