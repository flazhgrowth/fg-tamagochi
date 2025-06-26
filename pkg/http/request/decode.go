package request

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"

	"github.com/gorilla/schema"
)

func (req *RequestImpl) DecodeBody(dest any) error {
	b, err := io.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, dest); err != nil {
		return err
	}

	return nil
}

func (req *RequestImpl) DecodeQueryParam(dest any) error {
	decoder := schema.NewDecoder()
	decoder.IgnoreUnknownKeys(true)

	if err := decoder.Decode(dest, req.URL.Query()); err != nil {
		return err
	}

	return nil
}

func (req *RequestImpl) URLParamDecode(dest any) error {
	v := reflect.ValueOf(dest)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("tags: dest must be a pointer to struct")
	}

	typ := reflect.TypeOf(dest)
	res := map[string]any{}
	for i := range typ.NumField() {
		field := typ.Field(i)

		tag := field.Tag.Get("urlparam")
		typetag := field.Tag.Get("urlparamtype")
		paramVal := req.URLParam(tag)
		switch typetag {
		case "number":
			actualValue, err := paramVal.Int64()
			if err != nil {
				continue
			}
			res[tag] = actualValue
		case "string":
			res[tag] = paramVal.String()
		case "bool":
			actualValue, err := paramVal.Bool()
			if err != nil {
				continue
			}
			res[tag] = actualValue
		default:
			continue
		}
	}
	resBytes, err := json.Marshal(res)
	if err != nil {
		return err
	}

	return json.Unmarshal(resBytes, dest)
}
