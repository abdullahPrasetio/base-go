package helpers

import "github.com/go-playground/validator/v10"

type Response struct {
	Status       string      `json:"status"`
	ResponseCode int         `json:"response_code"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
}

func APIResponseSuccess(message string, data interface{}) Response {
	jsonResponse := Response{
		Status:       "success",
		Message:      message,
		Data:         data,
		ResponseCode: 200,
	}

	return jsonResponse
}

func APIResponseError(message string, code int, data interface{}) Response {
	jsonResponse := Response{
		Status:       "error",
		Message:      message,
		Data:         data,
		ResponseCode: code,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {

	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
