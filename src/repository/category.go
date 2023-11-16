package repository

import (
	"encoding/json"
	"errors"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/types"
	"github.com/spf13/viper"
)

type CategoryRepository struct{}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (r *CategoryRepository) ReadFile() ([]types.Category, error) {
	data := []types.Category{}

	filename := viper.GetString("CATEGORY_FILEPATH")
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

func (r *CategoryRepository) GetByID(ID uint) (types.Category, error) {
	category := types.Category{}
	categories, err := r.ReadFile()
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
