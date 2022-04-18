package repositories

import (
	"encoding/json"
	"os"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func parseJSONFastingArray(jsonData []byte, data []types.Fasting) ([]types.Fasting, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func FastingsReadFile() ([]types.Fasting, error) {
	data := []types.Fasting{}

	filename := os.Getenv("FASTING_FILEPATH")
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	result, err := parseJSONFastingArray(jsonData, data)
	if err != nil {
		return data, err
	}

	return result, nil
}
