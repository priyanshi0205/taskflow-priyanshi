package middleware

import (
	"log/slog"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		attrs := []any{
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"route", c.FullPath(),
			"status", c.Writer.Status(),
			"latency_ms", time.Since(start).Milliseconds(),
			"client_ip", c.ClientIP(),
		}
		if claims, ok := ClaimsFromContext(c); ok && claims != nil && strings.TrimSpace(claims.UserID) != "" {
			attrs = append(attrs, "user_id", claims.UserID)
		}

		switch status := c.Writer.Status(); {
		case status >= 500:
			slog.Error("http_request_server_error", attrs...)
		case status >= 400:
			slog.Warn("http_request_client_error", attrs...)
		default:
			slog.Info("http_request", attrs...)
		}
	}
}
