package apierrors

import (
	"net/http"
	"strings"
)

func ErrorInternalServerError(msgs ...string) HTTPError {
	if len(msgs) == 0 {
		msgs = []string{"internal server error"}
	}

	return HTTPError{
		StatusCode: http.StatusInternalServerError,
		Code:       "internal_server_error",
		Message:    strings.Join(msgs, ","),
	}
}
