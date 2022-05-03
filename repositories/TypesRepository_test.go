package repositories

import (
	"testing"
)

func TestTypesReadFile(t *testing.T) {
	t.Setenv("TYPE_FILEPATH", "../data/types.json")
	_, err := TypesReadFile()
	if err != nil {
		t.Errorf("Cannot read types json file : %s", err.Error())
	}
}

func TestTypesGetByID(t *testing.T) {
	t.Setenv("TYPE_FILEPATH", "../data/types.json")
	_, err := TypesGetByID(1)
	if err != nil {
		t.Errorf("Cannot get type by ID 1 : %s", err.Error())
	}
}
