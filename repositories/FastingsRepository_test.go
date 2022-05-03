package repositories

import (
	"testing"
)

func TestFastingsReadFile(t *testing.T) {
	t.Setenv("FASTING_FILEPATH", "../data/sunnah-fastings.json")
	_, err := FastingsReadFile()
	if err != nil {
		t.Errorf("Cannot read fastings json file : %s", err.Error())
	}
}
