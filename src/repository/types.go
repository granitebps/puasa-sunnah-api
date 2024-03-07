package repository

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/model"
)

type TypesRepository struct {
	Core *core.Core
}

func NewTypesRepository(c *core.Core) *TypesRepository {
	return &TypesRepository{
		Core: c,
	}
}

func (r *TypesRepository) GetAll(ctx context.Context) (res []model.Type, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Find(&res).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}

func (r *TypesRepository) GetByID(ctx context.Context, id uint) (res model.Type, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		First(&res, id).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}

func (r *TypesRepository) Create(ctx context.Context, types *model.Type) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Create(&types).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}

func (r *TypesRepository) Update(ctx context.Context, types *model.Type) (err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Save(&types).Error
	if err != nil {
		err = merry.Wrap(err)
		return
	}
	return
}
