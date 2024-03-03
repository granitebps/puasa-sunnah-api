package repository

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/model"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
)

type FastingRepository struct {
	Core *core.Core
}

func NewFastingRepository(c *core.Core) *FastingRepository {
	return &FastingRepository{
		Core: c,
	}
}

func (r *FastingRepository) GetAll(ctx context.Context, req *requests.FastingRequest) (res []model.Fasting, err error) {
	query := r.Core.Database.Db.
		WithContext(ctx).
		Preload("Category").
		Preload("Type")

	if req.TypeID != "" {
		query.Where("type_id = ?", req.TypeID)
	}
	if req.Year != "" {
		query.Where("year = ?", req.Year)
	}
	if req.Month != "" {
		query.Where("month = ?", req.Month)
	}
	if req.Day != "" {
		query.Where("day = ?", req.Day)
	}

	err = query.
		Find(&res).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}
