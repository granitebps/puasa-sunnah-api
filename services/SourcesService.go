package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

type SourceService struct {
	SourceRepo *repositories.SourceRepository
}

func NewSourceService(sourceRepo *repositories.SourceRepository) *SourceService {
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
