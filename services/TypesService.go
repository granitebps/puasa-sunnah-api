package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func TypesGetAll() ([]types.Type, error) {
	data, err := repositories.TypesReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
