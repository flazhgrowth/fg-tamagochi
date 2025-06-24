package request

import (
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ParamsValue string

func (params ParamsValue) String() string {
	return string(params)
}

func (params ParamsValue) Int64() (int64, error) {
	return strconv.ParseInt(string(params), 10, 64)
}

func (req *RequestImpl) URLParam(key string) ParamsValue {
	return ParamsValue(chi.URLParam(req.Request, key))
}
