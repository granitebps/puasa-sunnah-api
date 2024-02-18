package repository

import (
	"encoding/json"
	"errors"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/types"
	"github.com/spf13/viper"
)

type TypesRepository struct{}

func NewTypesRepository() *TypesRepository {
	return &TypesRepository{}
}

func (r *TypesRepository) ReadFile() ([]types.Type, error) {
	data := []types.Type{}

	filename := viper.GetString("TYPE_FILEPATH")
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
