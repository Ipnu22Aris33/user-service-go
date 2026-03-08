package httpx

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func DecodeJSON(r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func DecodeAndValidate(r *http.Request, dst any) error {
	if err := DecodeJSON(r, dst); err != nil {
		return err
	}

	return validate.Struct(dst)
}
