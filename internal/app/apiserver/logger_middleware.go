package apiserver

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func CustomResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(logger *logrus.Logger) func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			elapsedTime := time.Now().Sub(startTime).Milliseconds()
			userAgent := r.Header.Get("User-Agent")
			requestId := r.Context().Value("requestId")

			rw := CustomResponseWriter(w)

			logger.Infof("[requestId] %s [Resource] %s %s - [statusCode] - %d [latency] %d - agent %s", requestId, r.Method, r.URL, rw.statusCode, elapsedTime, userAgent)
		})
	}
}
