package repositories

import (
	"encoding/json"
	"errors"
	"os"

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

	filename := os.Getenv("CATEGORY_FILEPATH")
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

func CategoriesGetByID(ID uint) (types.Category, error) {
	category := types.Category{}
	categories, err := CategoriesReadFile()
	if err != nil {
		return category, err
	}

	for _, c := range categories {
		if c.ID == ID {
			category = c
		}
	}

	if category.ID == 0 {
		return category, errors.New("category not found")
	}

	return category, nil
}
