package services

import "testing"

func TestTypesGetAll(t *testing.T) {
	t.Setenv("TYPE_FILEPATH", "../data/types.json")
	_, err := TypesGetAll()
	if err != nil {
		t.Errorf("Error get types data : %s", err.Error())
	}
}
