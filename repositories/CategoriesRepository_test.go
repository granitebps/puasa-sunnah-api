package repositories

import (
	"testing"
)

func TestCategoriesReadFile(t *testing.T) {
	t.Setenv("CATEGORY_FILEPATH", "../data/categories.json")
	_, err := CategoriesReadFile()
	if err != nil {
		t.Errorf("Cannot read categories json file : %s", err.Error())
	}
}

func TestCategoriesGetByID(t *testing.T) {
	t.Setenv("CATEGORY_FILEPATH", "../data/categories.json")
	_, err := CategoriesGetByID(1)
	if err != nil {
		t.Errorf("Cannot get category by ID : %s", err.Error())
	}
}
