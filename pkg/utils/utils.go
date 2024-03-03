package utils

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/spf13/viper"
)

func CamelToSnake(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func QueryToUint(query string) uint {
	queryString, _ := strconv.Atoi(query)
	return uint(queryString)
}

func IsProduction() bool {
	env := strings.ToLower(viper.GetString(constants.APP_ENV))
	return env == "production" || env == "prod"
}

func ReadJsonFile(filename string) ([]byte, error) {
	var emptyData []byte

	jsonFile, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return emptyData, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return emptyData, err
	}

	return jsonData, nil
}

// Convert struct to json string.
// Error ignored
func StructToJSONString(s any) string {
	jsonString, _ := json.Marshal(s)
	return string(jsonString)
}
