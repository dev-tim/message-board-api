package apiserver

import (
	"net/http"
	"strconv"
)

func extractIntParam(r *http.Request, key string, defaultVal int) (*int, error) {
	query := r.URL.Query()
	val := query.Get(key)

	if len(val) == 0 {
		return &defaultVal, nil
	}

	if atoi, err := strconv.Atoi(val); err != nil {
		return nil, err
	} else {
		return &atoi, nil
	}
}
