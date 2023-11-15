package repositories

import (
	"encoding/json"
	"errors"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

type TypesRepository struct {
	Config *configs.Config
}

func NewTypesRepository(c *configs.Config) *TypesRepository {
	return &TypesRepository{
		Config: c,
	}
}

func (r *TypesRepository) ReadFile() ([]types.Type, error) {
	data := []types.Type{}

	filename := viper.GetString("TYPE_FILEPATH")
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

func (r *TypesRepository) GetByID(ID uint) (types.Type, error) {
	typeData := types.Type{}
	typesData, err := r.ReadFile()
	if err != nil {
		return typeData, merry.Wrap(err)
	}

	for _, t := range typesData {
		if t.ID == ID {
			typeData = t
		}
	}

	if typeData.ID == 0 {
		return typeData, merry.Wrap(errors.New("type not found"))
	}

	return typeData, nil
}
