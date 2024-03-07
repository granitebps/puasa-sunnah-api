package service

import (
	"context"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/granitebps/puasa-sunnah-api/src/transformer"
)

type FastingService struct {
	FastingRepo  *repository.FastingRepository
	CategoryRepo *repository.CategoryRepository
	TypesRepo    *repository.TypesRepository
}

func NewFastingService(fastingRepo *repository.FastingRepository, categoryRepo *repository.CategoryRepository, typesRepo *repository.TypesRepository) *FastingService {
	return &FastingService{
		FastingRepo:  fastingRepo,
		CategoryRepo: categoryRepo,
		TypesRepo:    typesRepo,
	}
}

func (s *FastingService) GetAll(ctx context.Context, req *requests.FastingRequest) (res []transformer.FastingTransformer, err error) {
	data, err := s.FastingRepo.GetAll(ctx, req)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	for _, f := range data {
		dateTime := time.Time(f.Date)
		formattedDate := dateTime.Format("2006-01-02")
		res = append(res, transformer.FastingTransformer{
			ID:         f.ID,
			CategoryID: f.CategoryID,
			TypeID:     f.TypeID,
			Date:       formattedDate,
			Year:       uint(f.Year),
			Month:      uint(f.Month),
			Day:        uint(f.Day),
			Category: transformer.CategoryTransformer{
				ID:   f.Category.ID,
				Name: f.Category.Name,
			},
			Type: transformer.TypeTransformer{
				ID:   f.Type.ID,
				Name: f.Type.Name,
			},
		})
	}

	return
}
