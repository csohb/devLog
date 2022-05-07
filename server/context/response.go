package context

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type CommonError struct {
	Message string `json:"message"`
}

type CommonResponse struct {
	Code  int          `json:"code"`
	Error *CommonError `json:"error,omitempty"`
	Data  interface{}  `json:"data"`
}

func (resp CommonResponse) Send(c echo.Context) error {
	return c.JSON(http.StatusOK, &resp)
}

func BadJSON(msg string) *CommonResponse {
	return &CommonResponse{
		Code:  http.StatusBadRequest,
		Error: &CommonError{Message: msg},
		Data:  nil,
	}
}

func SuccessJSON(data interface{}) *CommonResponse {
	return &CommonResponse{
		Code: http.StatusOK,
		Data: data,
	}
}

func EmptyJSON() *CommonResponse {
	return &CommonResponse{
		Code: http.StatusOK,
		Data: nil,
	}

}

func FailureJSON(code int, msg string) *CommonResponse {
	return &CommonResponse{
		Code:  code,
		Error: &CommonError{Message: msg},
		Data:  nil,
	}
}

func NotFoundJSON(msg string) *CommonResponse {
	return &CommonResponse{
		Code:  http.StatusNotFound,
		Error: &CommonError{Message: msg},
		Data:  nil,
	}
}

func InternalServerErrorJSON(msg string) *CommonResponse {
	return &CommonResponse{
		Code:  http.StatusInternalServerError,
		Error: &CommonError{Message: msg},
		Data:  nil,
	}
}
