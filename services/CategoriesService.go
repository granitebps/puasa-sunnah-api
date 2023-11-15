package services

import (
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/types"
)

type CategoryService struct {
	CategoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: categoryRepo,
	}
}

func (s *CategoryService) GetAll() ([]types.Category, error) {
	data, err := s.CategoryRepo.ReadFile()
	if err != nil {
		return data, err
	}

	return data, nil
}
