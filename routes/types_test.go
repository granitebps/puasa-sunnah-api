package routes

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestCategoriesIndex(t *testing.T) {
	t.Setenv("CATEGORY_FILEPATH", "../data/categories.json")
	app := fiber.New()

	InitRoutes(app)

	req, _ := http.NewRequest("GET", "/api/v1/categories", nil)

	res, err := app.Test(req, -1)

	if err != nil {
		t.Errorf("Error when send request to categories index : %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("Error when send request to categories index : %s", err.Error())
	}
}
