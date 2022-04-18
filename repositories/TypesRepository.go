package repositories

import (
	"encoding/json"
	"path/filepath"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func parseJSONTypeArray(jsonData []byte, data []types.Type) ([]types.Type, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func TypesReadFile() ([]types.Type, error) {
	data := []types.Type{}

	filename, err := filepath.Abs("./data/sunnah-types.json")
	if err != nil {
		return data, err
	}
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	result, err := parseJSONTypeArray(jsonData, data)
	if err != nil {
		return data, err
	}

	return result, nil
}
