package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/utils"
	"github.com/rs/zerolog/log"
)

type ResWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResWriter(w http.ResponseWriter) *ResWriter {
	return &ResWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}

func (rw *ResWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := NewResWriter(w)
		next.ServeHTTP(rw, r)

		duration := fmt.Sprintf("%dns", time.Since(start).Nanoseconds())
		request := fmt.Sprintf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
		message := utils.GetLogMessage(r.Context())

		status := response.Success
		logEvent := log.Info()
		if rw.statusCode >= 400 {
			status = response.Error
			logEvent = log.Error()
		}

		logEvent.
			Str("request", request).
			Str("status", string(status)).
			Int("status_code", rw.statusCode).
			Str("status_description", strings.ToLower(http.StatusText(rw.statusCode))).
			Interface("message", message).
			Str("process_time", duration).
			Send()
	})
}
