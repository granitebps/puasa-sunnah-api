package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func FastingsGetAll() ([]types.Fasting, error) {
	data, err := repositories.FastingsReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
