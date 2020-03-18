package apiserver

import "net/http"

type Error struct {
	RequestId string
	Message   string
}

func NewError(r *http.Request, message string) *Error {
	requestId := r.Context().Value("requestId")
	if requestId == nil {
		requestId = ""
	}

	return &Error{
		RequestId: requestId.(string),
		Message:   message,
	}
}
