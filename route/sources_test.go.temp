package routes

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func SourcesIndex(t *testing.T) {
	t.Setenv("SOURCE_FILEPATH", "../data/sources.json")
	app := fiber.New()

	InitRoutes(app)

	req, _ := http.NewRequest("GET", "/api/v1/sources", nil)

	res, err := app.Test(req, -1)

	if err != nil {
		t.Errorf("Error when send request to sources index : %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("Error when send request to sources index : %s", err.Error())
	}
}
