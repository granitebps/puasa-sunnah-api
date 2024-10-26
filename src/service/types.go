package service

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/transformer"
)

type TypesService struct {
	TypesRepo *repository.TypesRepository
}

func NewTypesService(typesRepo *repository.TypesRepository) *TypesService {
	return &TypesService{
		TypesRepo: typesRepo,
	}
}

func (s *TypesService) GetAll(ctx context.Context) (res []transformer.TypeTransformer, err error) {
	data, err := s.TypesRepo.GetAll(ctx)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	for _, t := range data {
		res = append(res, transformer.TypeTransformer{
			ID:              t.ID,
			Name:            t.Name,
			Description:     t.Description,
			BackgroundColor: t.BackgroundColor,
			TextColor:       t.TextColor,
		})
	}

	return
}
