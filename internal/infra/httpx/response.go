package httpx

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Message string `json:"message"`
}

func NewBadRequestResponse(e echo.Context, err error) error {
	return e.JSON(http.StatusBadRequest, Response{Message: err.Error()})
}

func NewInternalServerErrorResponse(e echo.Context, err error) error {
	return e.JSON(http.StatusInternalServerError, Response{Message: err.Error()})
}

func NewCreatedResponse(e echo.Context, data any) error {
	return e.JSON(http.StatusCreated, data)
}

func NewOKResponse(e echo.Context, data any) error {
	return e.JSON(http.StatusOK, data)
}
