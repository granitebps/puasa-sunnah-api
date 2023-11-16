package routes

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func FastingsIndex(t *testing.T) {
	t.Setenv("FASTING_FILEPATH", "../data/fastings.json")
	app := fiber.New()

	InitRoutes(app)

	req, _ := http.NewRequest("GET", "/api/v1/fastings", nil)

	res, err := app.Test(req, -1)

	if err != nil {
		t.Errorf("Error when send request to fastings index : %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("Error when send request to fastings index : %s", err.Error())
	}
}
