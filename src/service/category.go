package service

import (
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/types"
)

type CategoryService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
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
