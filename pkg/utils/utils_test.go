package utils_test

import (
	"testing"

	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
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
