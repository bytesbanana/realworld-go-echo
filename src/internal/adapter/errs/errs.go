package errs

import "errors"

type ErrorResponse struct {
	Errors map[string]interface{} `json:"errors"`
}

var (
	ErrAlreadyBeenTaken = errors.New("has already been taken")
)

func ParseError(args ...error) ErrorResponse {
	e := ErrorResponse{
		Errors: make(map[string]interface{}),
	}

	for _, err := range args {
		if errors.Is(err, ErrAlreadyBeenTaken) {
			e.Errors["email"] = []string{"has already been taken"}
		}
	}

	return e
}
