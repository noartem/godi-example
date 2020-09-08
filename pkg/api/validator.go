package api

import (
	"github.com/go-playground/validator"
	"github.com/noartem/godi-example/pkg/util"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func NewValidator() (*CustomValidator, error) {
	validate := validator.New()
	err := validate.RegisterValidation("password", func(field validator.FieldLevel) bool {
		return util.ValidatePassword(field.Field().String())
	})
	if err != nil {
		return nil, err
	}

	return &CustomValidator{validator: validate}, nil
}
