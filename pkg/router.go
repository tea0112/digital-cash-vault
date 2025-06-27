package pkg

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func BindJSON[T any](r *http.Request) (*T, error) {
	var val T
	if err := json.NewDecoder(r.Body).Decode(&val); err != nil {
		return nil, err
	}

	return &val, nil
}

func ParseInt(s string, defaultVal int) int {
	if s == "" {
		return defaultVal
	}
	i, err := strconv.Atoi(s)
	if err != nil || i < 1 {
		return defaultVal
	}
	return i
}
