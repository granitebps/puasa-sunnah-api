package service

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/src/model"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/granitebps/puasa-sunnah-api/src/transformer"
)

type AdminService struct {
	CategoryRepo *repository.CategoryRepository
}

func NewAdminService(categoryRepo *repository.CategoryRepository) *AdminService {
	return &AdminService{
		CategoryRepo: categoryRepo,
	}
}

func (s *AdminService) CreateCategory(ctx context.Context, req *requests.CategoryRequest) (trans transformer.CategoryTransformer, err error) {
	category := model.Category{
		Name: req.Name,
	}

	err = s.CategoryRepo.Create(ctx, &category)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = category.ID
	trans.Name = category.Name

	return
}

func (s *AdminService) UpdateCategory(ctx context.Context, id uint, req *requests.CategoryRequest) (trans transformer.CategoryTransformer, err error) {
	cat, err := s.CategoryRepo.GetByID(ctx, id)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	cat.Name = req.Name

	err = s.CategoryRepo.Update(ctx, &cat)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = cat.ID
	trans.Name = cat.Name

	return
}
