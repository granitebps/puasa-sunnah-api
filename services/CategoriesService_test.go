package services

import "testing"

func TestCategoriesGetAll(t *testing.T) {
	t.Setenv("CATEGORY_FILEPATH", "../data/categories.json")
	_, err := CategoriesGetAll()
	if err != nil {
		t.Errorf("Error get categories data : %s", err.Error())
	}
}
