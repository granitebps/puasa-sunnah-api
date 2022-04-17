package repositories

import (
	"encoding/json"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func parseJSONArray(jsonData []byte, data []types.Source) ([]types.Source, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func SourcesReadFile() ([]types.Source, error) {
	data := []types.Source{}

	filename := "data/sources.json"
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	result, err := parseJSONArray(jsonData, data)
	if err != nil {
		return data, err
	}

	return result, nil
}
