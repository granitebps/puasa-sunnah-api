package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

type TypesService struct {
	TypesRepo *repositories.TypesRepository
}

func NewTypesService(typesRepo *repositories.TypesRepository) *TypesService {
	return &TypesService{
		TypesRepo: typesRepo,
	}
}

func (s *TypesService) GetAll() ([]types.Type, error) {
	data, err := s.TypesRepo.ReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
