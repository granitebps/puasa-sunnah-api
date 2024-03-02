package repository

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/src/model"
)

type SourceRepository struct {
	Core *core.Core
}

func NewSourceRepository(c *core.Core) *SourceRepository {
	return &SourceRepository{
		Core: c,
	}
}

func (r *SourceRepository) GetAll(ctx context.Context) (res []model.Source, err error) {
	err = r.Core.Database.Db.
		WithContext(ctx).
		Find(&res).Error

	if err != nil {
		err = merry.Wrap(err)
		return
	}

	return
}
