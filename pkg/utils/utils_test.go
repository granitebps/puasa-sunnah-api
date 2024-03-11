package utils_test

import (
	"errors"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

func TestCamelToSnake(t *testing.T) {
	t.Run("should convert camel case to snake case", func(t *testing.T) {
		camelCase := "camelCase"
		expected := "camel_case"
		actual := utils.CamelToSnake(camelCase)
		assert.Equal(t, expected, actual)
	})
}

func TestQueryToUint(t *testing.T) {
	t.Run("should convert string to uint", func(t *testing.T) {
		query := "1"
		expected := uint(1)
		actual := utils.QueryToUint(query)
		assert.Equal(t, expected, actual)
	})
}

func TestIsProduction(t *testing.T) {
	t.Run("should return true if environment is production", func(t *testing.T) {
		viper.Set(constants.APP_ENV, "production")
		expected := true
		actual := utils.IsProduction()
		assert.Equal(t, expected, actual)
	})
}

func TestStructToJSONString(t *testing.T) {
	t.Run("should convert struct to json string", func(t *testing.T) {
		type testStruct struct {
			Name string `json:"name"`
		}
		test := testStruct{Name: "test"}
		expected := "{\"name\":\"test\"}"
		actual := utils.StructToJSONString(test)
		assert.Equal(t, expected, actual)
	})
}

func TestReturnSuccessResponse(t *testing.T) {
	t.Run("shoud return no error for success response", func(t *testing.T) {
		app := fiber.New()
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

		message := "Success message"
		data := map[string]interface{}{"key": "value"}
		err := utils.ReturnSuccessResponse(ctx, 0, message, data)

		assert.NoError(t, err)
	})
}

func TestReturnErrorResponse(t *testing.T) {
	t.Run("shoud return no error for error response", func(t *testing.T) {
		app := fiber.New()
		ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

		errMsg := errors.New("Error message")
		err := utils.ReturnErrorResponse(ctx, errMsg, nil)

		assert.NoError(t, err)
	})
}
