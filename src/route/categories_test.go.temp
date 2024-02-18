package routes

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TypesIndex(t *testing.T) {
	t.Setenv("TYPE_FILEPATH", "../data/types.json")
	app := fiber.New()

	InitRoutes(app)

	req, _ := http.NewRequest("GET", "/api/v1/types", nil)

	res, err := app.Test(req, -1)

	if err != nil {
		t.Errorf("Error when send request to types index : %s", err.Error())
	}

	if res.StatusCode != 200 {
		t.Errorf("Error when send request to types index : %s", err.Error())
	}
}
