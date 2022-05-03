package services

import "testing"

func TestSourcesGetAll(t *testing.T) {
	t.Setenv("SOURCE_FILEPATH", "../data/sources.json")
	_, err := SourcesGetAll()
	if err != nil {
		t.Errorf("Error get sources data : %s", err.Error())
	}
}
