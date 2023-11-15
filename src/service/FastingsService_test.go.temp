package services

import (
	"testing"

	"github.com/granitebps/puasa-sunnah-api/requests"
)

func TestFastingsGetAll(t *testing.T) {
	t.Setenv("FASTING_FILEPATH", "../data/sunnah-fastings.json")
	request := requests.FastingRequest{}
	_, err := FastingsGetAll(request)
	if err != nil {
		t.Errorf("Error get categories data : %s", err.Error())
	}
}
