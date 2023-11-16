package repository

import (
	"encoding/json"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

type FastingRepository struct{}

func NewFastingRepository() *FastingRepository {
	return &FastingRepository{}
}

func (r *FastingRepository) ReadFile() ([]types.Fasting, error) {
	data := []types.Fasting{}

	filename := viper.GetString("FASTING_FILEPATH")
	jsonData, err := utils.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, merry.Wrap(err)
	}

	return data, nil
}
