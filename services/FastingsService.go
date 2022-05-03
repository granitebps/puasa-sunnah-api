package services

import (
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/requests"
	"github.com/granitebps/puasa-sunnah-api/types"
)

func FastingsGetAll(q requests.FastingRequest) ([]types.Fasting, error) {
	data, err := repositories.FastingsReadFile()
	if err != nil {
		return data, err
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
			checkTypeID = f.TypeID == helpers.QueryToUint(r.TypeID)
		}
		if r.Day != "" {
			checkDay = f.Day == helpers.QueryToUint(r.Day)
		}
		if r.Month != "" {
			checkMonth = f.Month == helpers.QueryToUint(r.Month)
		}
		if r.Year != "" {
			checkYear = f.Year == helpers.QueryToUint(r.Year)
		}

		if checkTypeID && checkDay && checkMonth && checkYear {
			result = append(result, f)
		}
	}

	return result
}
