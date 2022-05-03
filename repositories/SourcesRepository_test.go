package repositories

import (
	"testing"
)

func TestSourcesReadFile(t *testing.T) {
	t.Setenv("SOURCE_FILEPATH", "../data/sources.json")
	_, err := SourcesReadFile()
	if err != nil {
		t.Errorf("Cannot read sources json file : %s", err.Error())
	}
}
