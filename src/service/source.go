package service

import (
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/types"
)

type SourceService struct {
	SourceRepo *repository.SourceRepository
}

func NewSourceService(sourceRepo *repository.SourceRepository) *SourceService {
	return &SourceService{
		SourceRepo: sourceRepo,
	}
}

func (s *SourceService) GetAll() ([]types.Source, error) {
	data, err := s.SourceRepo.ReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
