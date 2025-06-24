package request

import (
	"encoding/json"
	"io"

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
