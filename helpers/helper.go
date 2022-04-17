package helpers

import (
	"io/ioutil"
	"os"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Code    int    `json:"code"`
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
