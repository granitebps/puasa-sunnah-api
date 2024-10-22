package repository

import (
	"context"
	"fmt"
	"time"

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

	year := req.Year
	month := req.Month
	if year == "" {
		year = fmt.Sprintf("%d", time.Now().Year())
	}
	if month == "" {
		month = fmt.Sprintf("%d", time.Now().Month())
	}

	if req.TypeID != "" {
		query.Where("type_id = ?", req.TypeID)
	}
	query.Where("year = ?", year)
	query.Where("month = ?", month)
	if req.Day != "" {
		query.Where("day = ?", req.Day)
	}

	err = query.
		Order("date ASC").
		Find(&res).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}

func (r *FastingRepository) GetByID(ctx context.Context, id uint) (res model.Fasting, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Preload("Category").
		Preload("Type").
		First(&res, id).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}

func (r *FastingRepository) GetByDateAndType(ctx context.Context, date string, typeId uint) (res model.Fasting, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Preload("Category").
		Preload("Type").
		Where("date = ?", date).
		Where("type_id = ?", typeId).
		First(&res).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}

func (r *FastingRepository) Create(ctx context.Context, fasting *model.Fasting) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Preload("Category").
		Preload("Type").
		Create(&fasting).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}

func (r *FastingRepository) Update(ctx context.Context, fasting *model.Fasting) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Updates(&fasting).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}

func (r *FastingRepository) Delete(ctx context.Context, fasting *model.Fasting) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Delete(&fasting).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}
