package service

import (
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/types"
)

type TypesService struct {
	TypesRepo *repository.TypesRepository
}

func NewTypesService(typesRepo *repository.TypesRepository) *TypesService {
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
