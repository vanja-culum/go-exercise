package helper

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var defaultValidator = validator.New(validator.WithRequiredStructEnabled())

func Bind[T any](r *http.Request) (T, error) {
	var body T
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return body, err
	}

	if err := defaultValidator.Struct(body); err != nil {
		return body, err
	}

	return body, nil
}
