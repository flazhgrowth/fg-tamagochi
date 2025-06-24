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
