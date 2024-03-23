package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/granitebps/puasa-sunnah-api/src/model"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/stretchr/testify/assert"
	"gorm.io/datatypes"
)

func TestFastingGetAll(t *testing.T) {
	queryFasting := regexp.QuoteMeta("SELECT * FROM `fastings`")
	queryCat := regexp.QuoteMeta("SELECT * FROM `categories`")
	queryType := regexp.QuoteMeta("SELECT * FROM `types`")
	t.Run("should return all sources", func(t *testing.T) {
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

		repo := repository.NewFastingRepository(c)
		req := requests.FastingRequest{
			TypeID: "1",
			Day:    "1",
			Month:  "1",
			Year:   "2021",
		}
		fastings, err := repo.GetAll(context.Background(), &req)
		assert.NoError(t, err)
		assert.Len(t, fastings, 1)
		assert.Equal(t, fastings[0].ID, uint(1))
		assert.Equal(t, fastings[0].CategoryID, uint(1))
		assert.Equal(t, fastings[0].TypeID, uint(1))
		assert.Equal(t, time.Time(fastings[0].Date), now)
		assert.Equal(t, fastings[0].Year, uint32(2021))
		assert.Equal(t, fastings[0].Month, uint32(1))
		assert.Equal(t, fastings[0].Day, uint32(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(queryFasting).WillReturnError(errMock)

		repo := repository.NewFastingRepository(c)
		req := requests.FastingRequest{}
		fastings, err := repo.GetAll(context.Background(), &req)
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, fastings, 0)
	})
}

func TestFastingGetByID(t *testing.T) {
	queryFasting := regexp.QuoteMeta("SELECT * FROM `fastings`")
	queryCat := regexp.QuoteMeta("SELECT * FROM `categories`")
	queryType := regexp.QuoteMeta("SELECT * FROM `types`")
	t.Run("should return all sources", func(t *testing.T) {
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

		repo := repository.NewFastingRepository(c)
		fasting, err := repo.GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, fasting.ID, uint(1))
		assert.Equal(t, fasting.CategoryID, uint(1))
		assert.Equal(t, fasting.TypeID, uint(1))
		assert.Equal(t, time.Time(fasting.Date), now)
		assert.Equal(t, fasting.Year, uint32(2021))
		assert.Equal(t, fasting.Month, uint32(1))
		assert.Equal(t, fasting.Day, uint32(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(queryFasting).WillReturnError(errMock)

		repo := repository.NewFastingRepository(c)
		fasting, err := repo.GetByID(context.Background(), 1)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, fasting.ID, uint(0))
	})
}

func TestFastingGetByDateAndType(t *testing.T) {
	queryFasting := regexp.QuoteMeta("SELECT * FROM `fastings`")
	queryCat := regexp.QuoteMeta("SELECT * FROM `categories`")
	queryType := regexp.QuoteMeta("SELECT * FROM `types`")
	t.Run("should return fasting by date and type", func(t *testing.T) {
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

		repo := repository.NewFastingRepository(c)
		fasting, err := repo.GetByDateAndType(context.Background(), "2021-01-01", 1)
		assert.NoError(t, err)
		assert.Equal(t, fasting.ID, uint(1))
		assert.Equal(t, fasting.CategoryID, uint(1))
		assert.Equal(t, fasting.TypeID, uint(1))
		assert.Equal(t, time.Time(fasting.Date), now)
		assert.Equal(t, fasting.Year, uint32(2021))
		assert.Equal(t, fasting.Month, uint32(1))
		assert.Equal(t, fasting.Day, uint32(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(queryFasting).WillReturnError(errMock)

		repo := repository.NewFastingRepository(c)
		fasting, err := repo.GetByDateAndType(context.Background(), "2021-01-01", 1)
		assert.ErrorIs(t, err, errMock)
		assert.Equal(t, fasting.ID, uint(0))
	})
}

func TestFastingCreate(t *testing.T) {
	query := regexp.QuoteMeta("INSERT INTO `fastings`")
	fastingData := model.Fasting{
		CategoryID: 1,
		TypeID:     1,
		Date:       datatypes.Date(time.Now()),
		Year:       2021,
		Month:      1,
		Day:        1,
	}
	t.Run("should create fasting", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewFastingRepository(c)
		err := repo.Create(context.Background(), &fastingData)
		assert.NoError(t, err)
		assert.Equal(t, fastingData.ID, uint(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		mock.ExpectBegin()
		errMock := errors.New("failed query")
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewFastingRepository(c)
		err := repo.Create(context.Background(), &fastingData)
		assert.ErrorIs(t, err, errMock)
	})
}

func TestFastingUpdate(t *testing.T) {
	query := regexp.QuoteMeta("UPDATE `fastings`")
	fastingData := model.Fasting{
		ID:         1,
		CategoryID: 1,
		TypeID:     1,
		Date:       datatypes.Date(time.Now()),
		Year:       2021,
		Month:      1,
		Day:        1,
	}
	t.Run("should update fasting", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		repo := repository.NewFastingRepository(c)
		err := repo.Update(context.Background(), &fastingData)
		assert.NoError(t, err)
	})
	t.Run("should handle failed query", func(t *testing.T) {
		mock.ExpectBegin()
		errMock := errors.New("failed query")
		mock.ExpectExec(query).WillReturnError(errMock)
		mock.ExpectRollback()

		repo := repository.NewFastingRepository(c)
		err := repo.Update(context.Background(), &fastingData)
		assert.ErrorIs(t, err, errMock)
	})
}
