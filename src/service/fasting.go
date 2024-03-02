package service

import (
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/repository"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/granitebps/puasa-sunnah-api/src/types"
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

func (s *FastingService) GetAll(q requests.FastingRequest) ([]types.Fasting, error) {
	data, err := s.FastingRepo.ReadFile()
	if err != nil {
		return data, err
	}

	// Append category
	// for i, v := range data {
	// 	category, _ := s.CategoryRepo.GetByID(v.CategoryID)
	// 	data[i].Category = category
	// }

	// Append type
	for i, v := range data {
		typeData, _ := s.TypesRepo.GetByID(v.TypeID)
		data[i].Type = typeData
	}

	filteredData := FastingsFilter(q, data)

	return filteredData, nil
}

func FastingsFilter(r requests.FastingRequest, d []types.Fasting) []types.Fasting {
	var result []types.Fasting

	for _, f := range d {
		checkTypeID := true
		checkDay := true
		checkMonth := true
		checkYear := true
		if r.TypeID != "" {
			checkTypeID = f.TypeID == utils.QueryToUint(r.TypeID)
		}
		if r.Day != "" {
			checkDay = f.Day == utils.QueryToUint(r.Day)
		}
		if r.Month != "" {
			checkMonth = f.Month == utils.QueryToUint(r.Month)
		}
		if r.Year != "" {
			checkYear = f.Year == utils.QueryToUint(r.Year)
		}

		if checkTypeID && checkDay && checkMonth && checkYear {
			result = append(result, f)
		}
	}

	return result
}
