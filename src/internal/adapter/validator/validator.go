package validator

import (
	v "github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *v.Validate
}

func (cv *CustomValidator) Validate(i any) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}

	return nil
}
