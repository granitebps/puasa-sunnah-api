package repositories

import (
	"encoding/json"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

type FastingRepository struct {
	Config *configs.Config
}

func NewFastingRepository(c *configs.Config) *FastingRepository {
	return &FastingRepository{
		Config: c,
	}
}

func (r *FastingRepository) ReadFile() ([]types.Fasting, error) {
	data := []types.Fasting{}

	filename := viper.GetString("FASTING_FILEPATH")
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, merry.Wrap(err)
	}

	return data, nil
}
