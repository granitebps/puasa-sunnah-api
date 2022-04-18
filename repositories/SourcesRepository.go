package repositories

import (
	"encoding/json"
	"path/filepath"

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

	filename, err := filepath.Abs("./data/sources.json")
	if err != nil {
		return data, err
	}
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
