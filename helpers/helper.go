package helpers

import (
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/spf13/viper"
)

type APIResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Success"`
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data"`
}

func SuccessAPIResponse(ctx *fiber.Ctx, message string, code int, data interface{}) error {
	responseData := APIResponse{
		Success: true,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return ctx.Status(code).JSON(responseData)
}

func FailedAPIResponse(ctx *fiber.Ctx, err error, data interface{}) error {
	msg := merry.UserMessage(err)
	code := merry.HTTPCode(err)

	if code == fiber.StatusInternalServerError && IsProduction() {
		// TODO: Log
		msg = errors.ErrInternalServerError.Error()
	} else {
		msg = err.Error()
	}

	responseData := APIResponse{
		Success: false,
		Message: msg,
		Code:    code,
		Data:    data,
	}

	return ctx.Status(code).JSON(responseData)
}

func ReadJsonFile(filename string) ([]byte, error) {
	var emptyData []byte

	jsonFile, err := os.Open(filename)
	if err != nil {
		return emptyData, err
	}
	defer jsonFile.Close()

	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		return emptyData, err
	}

	return jsonData, nil
}

func QueryToUint(query string) uint {
	queryString, _ := strconv.Atoi(query)
	return uint(queryString)
}

func IsProduction() bool {
	env := strings.ToLower(viper.GetString("APP_ENV"))
	return env == "production" || env == "prod"
}
