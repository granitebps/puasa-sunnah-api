package service

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/transformer"
)

type CategoryService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryService(categoryRepo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepo: categoryRepo,
	}
}

func (s *CategoryService) GetAll(ctx context.Context) (res []transformer.CategoryTransformer, err error) {
	data, err := s.CategoryRepo.GetAll(ctx)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	for _, c := range data {
		res = append(res, transformer.CategoryTransformer{
			ID:   c.ID,
			Name: c.Name,
		})
	}

	return
}
