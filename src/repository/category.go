package repository

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/model"
)

type CategoryRepository struct {
	Core *core.Core
}

func NewCategoryRepository(c *core.Core) *CategoryRepository {
	return &CategoryRepository{
		Core: c,
	}
}

func (r *CategoryRepository) GetAll(ctx context.Context) (res []model.Category, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Find(&res).Error
	return
}

func (r *CategoryRepository) GetByID(ctx context.Context, id uint) (res model.Category, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		First(&res, id).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}

func (r *CategoryRepository) Create(ctx context.Context, category *model.Category) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Create(&category).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}

func (r *CategoryRepository) Update(ctx context.Context, category *model.Category) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Save(&category).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}

func (r *CategoryRepository) Delete(ctx context.Context, category *model.Category) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Delete(&category).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}
