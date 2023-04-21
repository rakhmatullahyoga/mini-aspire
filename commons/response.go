package commons

type SuccessResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func BuildErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "failed",
		Message: message,
	}
}

func BuildResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Status: "success",
		Data:   data,
	}
}
