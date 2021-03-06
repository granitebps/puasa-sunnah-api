package repositories

import (
	"encoding/json"
	"os"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func parseJSONSourceArray(jsonData []byte, data []types.Source) ([]types.Source, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func SourcesReadFile() ([]types.Source, error) {
	data := []types.Source{}

	filename := os.Getenv("SOURCE_FILEPATH")
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	result, err := parseJSONSourceArray(jsonData, data)
	if err != nil {
		return data, err
	}

	return result, nil
}
