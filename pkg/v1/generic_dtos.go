package v1

import (
	"github.com/labstack/echo"
	"net/http"
)

type PostRequestBody struct {
	Id                 interface{}     `json:"id"`
	Type               string     `json:"type"`
	Attributes         interface{}     `json:"attributes"`
}


type ResponseDTO struct {
	Data    interface{} `json:"data"`
	Status  string      `json:"status"`
	Message interface{}      `json:"message"`
}

func GenerateSuccessResponse(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, ResponseDTO{
		Status:  "success",
		Message: message,
		Data:    data,

	})
}

func GenerateErrorResponse(c echo.Context, data interface{}, message interface{}) error {
	return c.JSON(http.StatusBadRequest, ResponseDTO{
		Status:  "error",
		Message: message,
		Data:    data,
	})
}
//func GenerateErrorResponse(c echo.Context, data interface{}, message []string) error {
//	return c.JSON(http.StatusBadRequest, ResponseDTO{
//		Status:  "error",
//		Message: message,
//		Data:    data,
//	})
//}