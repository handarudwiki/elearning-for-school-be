package helpers

import (
	"net/http"

	"github.com/handarudwiki/models/commons"
)

type response map[string]interface{}

type ErrValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

var ErrToResponseCode = map[error]int{
	commons.ErrCredentials:    http.StatusUnauthorized,
	commons.ErrNotFound:       http.StatusNotFound,
	commons.ErrConflict:       http.StatusConflict,
	commons.ErrInternalServer: http.StatusInternalServerError,
}

func ResponseSuccess(data interface{}) response {
	return response{
		"data": data,
	}
}

func ResponsePagination(data interface{}, pagination commons.Paginate) response {
	return response{
		"data":       data,
		"pagination": pagination,
	}
}

func ResponseError(message string) response {
	return response{
		"message": message,
	}
}

func ResponseErrorWithData(message string, data interface{}) response {
	return response{
		"message": message,
		"data":    data,
	}
}

func GetHttpStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	if code, ok := ErrToResponseCode[err]; ok {
		return code
	}

	return http.StatusInternalServerError
}
