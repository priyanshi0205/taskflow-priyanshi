package response

import (
	"log/slog"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backedn_go/internal/shared/apperror"
	"backedn_go/internal/shared/middleware"
)

func Error(c *gin.Context, appErr *apperror.AppError) {
	ErrorWithContext(c, appErr, "request failed")
}

func ErrorWithContext(c *gin.Context, appErr *apperror.AppError, logMessage string, attrs ...any) {
	if strings.TrimSpace(logMessage) == "" {
		logMessage = "request failed"
	}

	status := http.StatusInternalServerError
	errorMessage := "server error"
	if appErr != nil {
		status = appErr.Status
		errorMessage = appErr.Error
	}

	logAttrs := []any{
		"http_method", c.Request.Method,
		"http_path", c.Request.URL.Path,
		"http_route", c.FullPath(),
		"status", status,
		"error", errorMessage,
	}
	if claims, ok := middleware.ClaimsFromContext(c); ok && claims != nil && strings.TrimSpace(claims.UserID) != "" {
		logAttrs = append(logAttrs, "user_id", claims.UserID)
	}

	projectID, taskID := requestResourceIDs(c)
	if projectID != "" {
		logAttrs = append(logAttrs, "project_id", projectID)
	}
	if taskID != "" {
		logAttrs = append(logAttrs, "task_id", taskID)
	}
	if appErr != nil && appErr.Fields != nil {
		logAttrs = append(logAttrs, "error_fields", appErr.Fields)
	}
	logAttrs = append(logAttrs, attrs...)
	slog.Error(logMessage, logAttrs...)

	if appErr == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	if appErr.Fields != nil {
		c.JSON(appErr.Status, gin.H{"error": appErr.Error, "fields": appErr.Fields})
		return
	}
	c.JSON(appErr.Status, gin.H{"error": appErr.Error})
}

func requestResourceIDs(c *gin.Context) (string, string) {
	id := strings.TrimSpace(c.Param("id"))
	route := c.FullPath()

	projectID := strings.TrimSpace(c.Param("project_id"))
	taskID := strings.TrimSpace(c.Param("task_id"))

	if projectID == "" {
		projectID = strings.TrimSpace(c.Query("project_id"))
	}
	if taskID == "" {
		taskID = strings.TrimSpace(c.Query("task_id"))
	}

	switch {
	case strings.HasPrefix(route, "/projects/:id/tasks"):
		if projectID == "" {
			projectID = id
		}
	case strings.HasPrefix(route, "/projects/:id"):
		if projectID == "" {
			projectID = id
		}
	case strings.HasPrefix(route, "/tasks/:id"):
		if taskID == "" {
			taskID = id
		}
	}

	return projectID, taskID
}
