package apiserver

import (
	"net/http"
)

// Here we can add also error code for clients to handle different types of errors
// recoverable/unrecoverable for example
type Error struct {
	RequestId string
	Message   string
}

func NewError(w http.ResponseWriter, r *http.Request, statusCode int, message string) *Error {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	requestId := r.Context().Value("requestId")
	if requestId == nil {
		requestId = ""
	}

	return &Error{
		RequestId: requestId.(string),
		Message:   message,
	}
}
