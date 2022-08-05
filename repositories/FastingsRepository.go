package repositories

import (
	"encoding/json"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

func parseJSONFastingArray(jsonData []byte, data []types.Fasting) ([]types.Fasting, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func FastingsReadFile() ([]types.Fasting, error) {
	data := []types.Fasting{}

	filename := viper.GetString("FASTING_FILEPATH")
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	result, err := parseJSONFastingArray(jsonData, data)
	if err != nil {
		return data, err
	}

	// Append category
	for i, v := range result {
		category, _ := CategoriesGetByID(v.CategoryID)
		result[i].Category = category
	}

	// Append type
	for i, v := range result {
		typeData, _ := TypesGetByID(v.TypeID)
		result[i].Type = typeData
	}

	return result, nil
}
