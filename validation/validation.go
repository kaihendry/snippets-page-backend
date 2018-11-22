package validation

import "github.com/go-playground/validator"

type RequestValidator struct {
	validator *validator.Validate
}

func New() *RequestValidator {
	return &RequestValidator{validator: validator.New()}
}

func (v RequestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
