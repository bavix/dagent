package library

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusTeapot, err.Error())
	}
	return nil
}

func NewValidator() *CustomValidator {
	return &CustomValidator{validator: validator.New()}
}
