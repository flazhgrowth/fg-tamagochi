package request

import (
	"encoding/json"
	"io"
	"reflect"

	"github.com/flazhgrowth/fg-tamagochi/pkg/http/apierrors"
	"github.com/gorilla/schema"
)

func (req *RequestImpl) DecodeBody(dest any) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return apierrors.ErrorBadRequest("invalid request body")
	}
	defer req.Body.Close()

	if err = json.Unmarshal(b, dest); err != nil {
		return apierrors.ErrorBadRequest("invalid request body")
	}

	return nil
}

func (req *RequestImpl) DecodeQueryParam(dest any) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(dest, req.URL.Query()); err != nil {
		return apierrors.ErrorBadRequest("invalid query params")
	}

	return nil
}

func (req *RequestImpl) DecodeURLParam(dest any) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return apierrors.ErrorInternalServerError("tags: dest must be a pointer to struct")
	}

	typ := reflect.TypeOf(dest).Elem()
	val := reflect.ValueOf(dest).Elem()
	for i := range typ.NumField() {
		field := typ.Field(i)
		fieldVal := val.Field(i)
		if !fieldVal.CanSet() {
			continue
		}

		tag := field.Tag.Get("path")
		typetag := field.Tag.Get("pathtype")
		paramVal := req.URLParam(tag)
		switch typetag {
		case "number":
			actualValue, err := paramVal.Int64()
			if err != nil {
				continue
			}
			fieldVal.SetInt(actualValue)
		case "string":
			fieldVal.SetString(paramVal.String())
		case "bool":
			actualValue, err := paramVal.Bool()
			if err != nil {
				continue
			}
			fieldVal.SetBool(actualValue)
		default:
			continue
		}
	}

	return nil
}
