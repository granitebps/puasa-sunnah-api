package service_test

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/stretchr/testify/assert"
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

		req := requests.FastingRequest{
			TypeID: "1",
			Day:    "1",
			Month:  "1",
			Year:   "2021",
		}
		fastings, err := fastingService.GetAll(context.Background(), &req)
		assert.NoError(t, err)
		assert.Len(t, fastings, 1)
		assert.Equal(t, fastings[0].ID, uint(1))
		assert.Equal(t, fastings[0].CategoryID, uint(1))
		assert.Equal(t, fastings[0].TypeID, uint(1))
		assert.Equal(t, fastings[0].Date, now.Format("2006-01-02"))
		assert.Equal(t, fastings[0].Year, uint(2021))
		assert.Equal(t, fastings[0].Month, uint(1))
		assert.Equal(t, fastings[0].Day, uint(1))
	})
	t.Run("should handle failed query", func(t *testing.T) {
		sqlmock.NewRows([]string{"id", "category_id", "type_id", "date", "year", "month", "day"})
		errMock := errors.New("failed query")
		mock.ExpectQuery(queryFasting).WillReturnError(errMock)

		req := requests.FastingRequest{}
		fastings, err := fastingService.GetAll(context.Background(), &req)
		assert.ErrorIs(t, err, errMock)
		assert.Len(t, fastings, 0)
	})
}
