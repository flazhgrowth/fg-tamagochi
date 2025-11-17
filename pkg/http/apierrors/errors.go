package apierrors

import (
	"fmt"
)

type HTTPError struct {
	StatusCode int
	Code       string
	Message    string
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%s-%s", e.Code, e.Message)
}

func (e HTTPError) WithCode(strCode string) HTTPError {
	e.Code = strCode

	return e
}
