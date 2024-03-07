package service

import (
	"context"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/src/model"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/granitebps/puasa-sunnah-api/src/transformer"
	"gorm.io/datatypes"
)

type AdminService struct {
	CategoryRepo *repository.CategoryRepository
	SourceRepo   *repository.SourceRepository
	TypeRepo     *repository.TypesRepository
	FastingRepo  *repository.FastingRepository
}

func NewAdminService(
	categoryRepo *repository.CategoryRepository,
	sourceRepo *repository.SourceRepository,
	typeRepo *repository.TypesRepository,
	fastingRepo *repository.FastingRepository,
) *AdminService {
	return &AdminService{
		CategoryRepo: categoryRepo,
		SourceRepo:   sourceRepo,
		TypeRepo:     typeRepo,
		FastingRepo:  fastingRepo,
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

func (s *AdminService) CreateType(ctx context.Context, req *requests.TypeRequest) (trans transformer.TypeTransformer, err error) {
	types := model.Type{
		Name:        req.Name,
		Description: req.Description,
	}

	err = s.TypeRepo.Create(ctx, &types)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = types.ID
	trans.Name = types.Name
	trans.Description = types.Description

	return
}

func (s *AdminService) UpdateType(ctx context.Context, id uint, req *requests.TypeRequest) (trans transformer.TypeTransformer, err error) {
	types, err := s.TypeRepo.GetByID(ctx, id)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	types.Name = req.Name
	types.Description = req.Description

	err = s.TypeRepo.Update(ctx, &types)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = types.ID
	trans.Name = types.Name
	trans.Description = types.Description

	return
}

func (s *AdminService) CreateFasting(ctx context.Context, req *requests.FastingCreateUpdateRequest) (trans transformer.FastingTransformer, err error) {
	// Check if fasting already exists
	_, err = s.FastingRepo.GetByDateAndType(ctx, req.Date, req.TypeID)
	if err == nil {
		err = merry.New("Fasting already exists", merry.WithHTTPCode(fiber.StatusBadRequest))
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		err = merry.Wrap(err, merry.WithUserMessage("Invalid date format"), merry.WithHTTPCode(fiber.StatusBadRequest))
		return
	}

	fasting := model.Fasting{
		CategoryID: req.CategoryID,
		TypeID:     req.TypeID,
		Date:       datatypes.Date(date),
		Year:       uint32(req.Year),
		Month:      uint32(req.Month),
		Day:        uint32(req.Day),
	}

	err = s.FastingRepo.Create(ctx, &fasting)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	fasting, err = s.FastingRepo.GetByID(ctx, fasting.ID)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = fasting.ID
	trans.CategoryID = fasting.CategoryID
	trans.TypeID = fasting.TypeID
	trans.Year = uint(fasting.Year)
	trans.Month = uint(fasting.Month)
	trans.Day = uint(fasting.Day)

	dateTime := time.Time(fasting.Date)
	formattedDate := dateTime.Format("2006-01-02")
	trans.Date = formattedDate

	trans.Category = transformer.CategoryTransformer{
		ID:   fasting.Category.ID,
		Name: fasting.Category.Name,
	}
	trans.Type = transformer.TypeTransformer{
		ID:          fasting.Type.ID,
		Name:        fasting.Type.Name,
		Description: fasting.Type.Description,
	}

	return
}

func (s *AdminService) UpdateFasting(ctx context.Context, req *requests.FastingCreateUpdateRequest, id uint) (trans transformer.FastingTransformer, err error) {
	fasting, err := s.FastingRepo.GetByID(ctx, id)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	fasting.CategoryID = req.CategoryID
	fasting.TypeID = req.TypeID
	fasting.Year = uint32(req.Year)
	fasting.Month = uint32(req.Month)
	fasting.Day = uint32(req.Day)
	fasting.Category = model.Category{}
	fasting.Type = model.Type{}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		err = merry.Wrap(err, merry.WithUserMessage("Invalid date format"), merry.WithHTTPCode(fiber.StatusBadRequest))
		return
	}
	fasting.Date = datatypes.Date(date)

	// Validate fasting date
	existsFasting, err := s.FastingRepo.GetByDateAndType(ctx, req.Date, req.TypeID)
	if err == nil {
		if existsFasting.ID != id {
			err = merry.New("Fasting already exists", merry.WithHTTPCode(fiber.StatusBadRequest))
			return
		}
	}

	err = s.FastingRepo.Update(ctx, &fasting)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	newFasting, err := s.FastingRepo.GetByID(ctx, fasting.ID)
	if err != nil {
		err = merry.Wrap(err)
		return
	}

	trans.ID = newFasting.ID
	trans.CategoryID = newFasting.CategoryID
	trans.TypeID = newFasting.TypeID
	trans.Year = uint(newFasting.Year)
	trans.Month = uint(newFasting.Month)
	trans.Day = uint(newFasting.Day)

	dateTime := time.Time(newFasting.Date)
	formattedDate := dateTime.Format("2006-01-02")
	trans.Date = formattedDate

	trans.Category = transformer.CategoryTransformer{
		ID:   newFasting.Category.ID,
		Name: newFasting.Category.Name,
	}
	trans.Type = transformer.TypeTransformer{
		ID:          newFasting.Type.ID,
		Name:        newFasting.Type.Name,
		Description: newFasting.Type.Description,
	}

	return
}
