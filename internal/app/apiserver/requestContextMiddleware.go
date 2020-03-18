package apiserver

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestContext := context.WithValue(context.Background(), "requestId", uuid.New().String())
		defer requestContext.Done()

		next.ServeHTTP(w, r.Clone(requestContext))
	})
}
