package repositories

import (
	"encoding/json"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func parseJSONCategoryArray(jsonData []byte, data []types.Category) ([]types.Category, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func CategoriesReadFile() ([]types.Category, error) {
	data := []types.Category{}

	filename := "data/categories.json"
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	result, err := parseJSONCategoryArray(jsonData, data)
	if err != nil {
		return data, err
	}

	return result, nil
}
