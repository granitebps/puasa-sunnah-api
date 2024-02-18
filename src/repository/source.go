package repository

import (
	"encoding/json"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/types"
	"github.com/spf13/viper"
)

type SourceRepository struct{}

func NewSourceRepository() *SourceRepository {
	return &SourceRepository{}
}

func (r *SourceRepository) ReadFile() ([]types.Source, error) {
	data := []types.Source{}

	filename := viper.GetString("SOURCE_FILEPATH")
	jsonData, err := utils.ReadJsonFile(filename)
	if err != nil {
		return data, merry.Wrap(err)
	}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return data, merry.Wrap(err)
	}

	return data, nil
}
