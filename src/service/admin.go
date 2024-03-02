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
	SourceRepo   *repository.SourceRepository
}

func NewAdminService(
	categoryRepo *repository.CategoryRepository,
	sourceRepo *repository.SourceRepository,
) *AdminService {
	return &AdminService{
		CategoryRepo: categoryRepo,
		SourceRepo:   sourceRepo,
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

func (s *AdminService) CreateSource(ctx context.Context, req *requests.SourceRequest) (trans transformer.SourceTransformer, err error) {
	source := model.Source{
		Url: req.Url,
	}

	err = s.SourceRepo.Create(ctx, &source)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = source.ID
	trans.URL = source.Url

	return
}

func (s *AdminService) UpdateSource(ctx context.Context, id uint, req *requests.SourceRequest) (trans transformer.SourceTransformer, err error) {
	source, err := s.SourceRepo.GetByID(ctx, id)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	source.Url = req.Url

	err = s.SourceRepo.Update(ctx, &source)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = source.ID
	trans.URL = source.Url

	return
}
