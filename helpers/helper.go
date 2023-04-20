package helpers

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
)

type SuccessResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Success"`
	Code    int         `json:"code" example:"200"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message" example:"Failed"`
	Code    int    `json:"code" example:"400"`
}

func SuccessAPIResponse(
	message string,
	code int,
	data interface{},
) SuccessResponse {
	responseData := SuccessResponse{
		Success: true,
		Data:    data,
		Message: message,
		Code:    code,
	}

	return responseData
}

func FailedAPIResponse(
	message string,
	code int,
) FailedResponse {
	responseData := FailedResponse{
		Success: false,
		Message: message,
		Code:    code,
	}

	return responseData
}

func ReadJsonFile(filename string) ([]byte, error) {
	var emptyData []byte

	jsonFile, err := os.Open(filename)
	if err != nil {
		return emptyData, err
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
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
