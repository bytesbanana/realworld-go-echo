package errs

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Errors map[string][]string `json:"errors"`
}

var (
	ErrAlreadyBeenTaken = errors.New("has already been taken")
	ErrUnAuthorized     = errors.New("unauthorized")
)

const (
	required = "is required"
	email    = "is invalid"
	unknown  = "unknown error"
)

func getErrMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return required
	case "email":
		return email
	}

	return unknown
}

func ParseError(args ...error) ErrorResponse {
	e := ErrorResponse{
		Errors: make(map[string][]string),
	}

	for _, err := range args {
		if errors.Is(err, ErrAlreadyBeenTaken) {
			e.Errors["email"] = []string{"has already been taken"}
		}
	}

	for _, err := range args {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			for _, fe := range ve {
				key := strings.ToLower(fe.Field())
				if e.Errors[key] != nil {
					e.Errors[key] = []string{}
				}

				e.Errors[key] = append(e.Errors[key], getErrMsg(fe))
			}
		}
	}

	return e
}
