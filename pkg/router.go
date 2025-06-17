package pkg

import (
	"encoding/json"
	"net/http"
)

func BindJSON[T any](r *http.Request) (*T, error) {
	var val T
	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
		return nil, err
	}

	return &val, nil
}
