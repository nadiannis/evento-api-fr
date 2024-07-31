package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nadiannis/evento-api-fr/internal/domain/response"
	"github.com/nadiannis/evento-api-fr/internal/utils"
	"github.com/rs/zerolog/log"
)

func requestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := fmt.Sprintf("%dns", time.Since(start).Nanoseconds())
		request := fmt.Sprintf("%s %s %s", c.Request.Proto, c.Request.Method, c.Request.RequestURI)
		message := utils.GetLogMessage(c)

		status := response.Success
		logEvent := log.Info()
		statusCode := c.Writer.Status()
		if statusCode >= 400 {
			status = response.Error
			logEvent = log.Error()
		}

		logEvent.
			Str("request", request).
			Str("status", string(status)).
			Int("status_code", statusCode).
			Str("status_description", strings.ToLower(http.StatusText(statusCode))).
			Interface("message", message).
			Str("process_time", duration).
			Send()
	}
}
