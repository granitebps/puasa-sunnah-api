package helpers

type SuccessResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SuccessAPIResponse(
	message string,
	code int,
	data interface{},
) SuccessResponse {
	responseData := SuccessResponse{
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
		Message: message,
		Code:    code,
	}

	return responseData
}
