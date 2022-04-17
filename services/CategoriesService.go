package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func CategoriesGetAll() ([]types.Category, error) {
	data, err := repositories.CategoriesReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
