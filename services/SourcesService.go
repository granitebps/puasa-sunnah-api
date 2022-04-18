package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func SourcesGetAll() ([]types.Source, error) {
	data, err := repositories.SourcesReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
