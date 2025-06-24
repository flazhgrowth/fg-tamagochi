package response

import "net/http"

type ResponseImpl struct {
	http.ResponseWriter
}

type Response interface {
	// Respond writes http response.
	//
	// By default, the status code for a response will be http.StatusOK (200).
	// If the error type is not nil and the type of error is errors.HTTPError, then the errors.HTTPError.StatusCode will be use as the status code,
	// or http.StatusInternalServerError (500) will be used if the error is not errors.HTTPError
	//
	// But this behavior can be overriden by passing status code after the argument error. (Eg: res.Respond(someData, nil, http.StatusCreated)).
	//
	// Note that if you inspect the method, it accepts ellipsis of int, meaning, you can pass many status code right after the first one.
	// It will only use the first passed status code to override the status code. IT WILL NOT OVERRIDE THE STATUS CODE IF THE ERROR IS NOT NIL
	Respond(data any, err error, statusCode ...int)
}

func New(w http.ResponseWriter) Response {
	return &ResponseImpl{
		ResponseWriter: w,
	}
}
