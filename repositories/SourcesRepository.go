package repositories

import (
	"encoding/json"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

type SourceRepository struct {
	Config *configs.Config
}

func NewSourceRepository(c *configs.Config) *SourceRepository {
	return &SourceRepository{
		Config: c,
	}
}

func (r *SourceRepository) ReadFile() ([]types.Source, error) {
	data := []types.Source{}

	filename := viper.GetString("SOURCE_FILEPATH")
	jsonData, err := helpers.ReadJsonFile(filename)
	if err != nil {
		return data, merry.Wrap(err)
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, merry.Wrap(err)
	}

	return data, nil
}
