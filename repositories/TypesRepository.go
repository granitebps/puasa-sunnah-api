package repositories

import (
	"encoding/json"
	"errors"

	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

func parseJSONTypeArray(jsonData []byte, data []types.Type) ([]types.Type, error) {
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return data, err
	}

	return data, nil
}

func TypesReadFile() ([]types.Type, error) {
	data := []types.Type{}

	filename := viper.GetString("TYPE_FILEPATH")
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

func TypesGetByID(ID uint) (types.Type, error) {
	typeData := types.Type{}
	typesData, err := TypesReadFile()
	if err != nil {
		return typeData, err
	}

	for _, t := range typesData {
		if t.ID == ID {
			typeData = t
		}
	}

	if typeData.ID == 0 {
		return typeData, errors.New("type not found")
	}

	return typeData, nil
}
