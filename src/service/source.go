package service

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/transformer"
)

type SourceService struct {
	SourceRepo *repository.SourceRepository
}

func NewSourceService(sourceRepo *repository.SourceRepository) *SourceService {
	return &SourceService{
		SourceRepo: sourceRepo,
	}
}

func (s *SourceService) GetAll(ctx context.Context) (res []transformer.SourceTransformer, err error) {
	data, err := s.SourceRepo.GetAll(ctx)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	for _, s := range data {
		res = append(res, transformer.SourceTransformer{
			ID:  s.ID,
			URL: s.Url,
		})
	}

	return
}
