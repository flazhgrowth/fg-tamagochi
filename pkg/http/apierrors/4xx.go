package apierrors

import (
	"net/http"
	"strings"
)

func ErrorBadRequest(msgs ...string) HTTPError {
	if len(msgs) == 0 {
		msgs = []string{"bad request"}
	}

	return HTTPError{
		StatusCode: http.StatusBadRequest,
		Code:       "bad_request",
		Message:    strings.Join(msgs, ","),
	}
}

func ErrorUnauthorized(msgs ...string) HTTPError {
	if len(msgs) == 0 {
		msgs = []string{"unauthorized"}
	}

	return HTTPError{
		StatusCode: http.StatusUnauthorized,
		Code:       "unauthorized",
		Message:    strings.Join(msgs, ","),
	}
}

func ErrorForbidden(msgs ...string) HTTPError {
	if len(msgs) == 0 {
		msgs = []string{"forbidden"}
	}

	return HTTPError{
		StatusCode: http.StatusForbidden,
		Code:       "forbidden",
		Message:    strings.Join(msgs, ","),
	}
}

func ErrorUnprocessableEntity(msgs ...string) HTTPError {
	if len(msgs) == 0 {
		msgs = []string{"unprocessable entity"}
	}

	return HTTPError{
		StatusCode: http.StatusUnprocessableEntity,
		Code:       "unprocessable_entity",
		Message:    strings.Join(msgs, ","),
	}
}

func ErrorDataNotFound(msgs ...string) HTTPError {
	if len(msgs) == 0 {
		msgs = []string{"data not found"}
	}

	return HTTPError{
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
		Message:    strings.Join(msgs, ","),
	}
}
