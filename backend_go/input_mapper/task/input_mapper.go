package task

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"backedn_go/internal/shared/apperror"
	taskusecase "backedn_go/usecase/task"
)

func MapProjectID(c *gin.Context) (string, *apperror.AppError) {
	projectID := strings.TrimSpace(c.Param(Id))
	if projectID == "" {
		return "", apperror.New(http.StatusBadRequest, "project id is required")
	}
	return projectID, nil
}

func MapTaskID(c *gin.Context) (string, *apperror.AppError) {
	taskID := strings.TrimSpace(c.Param(Id))
	if taskID == "" {
		return "", apperror.New(http.StatusBadRequest, "task id is required")
	}
	return taskID, nil
}

func MapListFilters(c *gin.Context) (string, string) {
	status := strings.TrimSpace(c.Query(Status))
	assignee := strings.TrimSpace(c.Query(Assignee))
	return status, assignee
}

func MapCreateTask(c *gin.Context) (taskusecase.CreateTaskInput, *apperror.AppError) {
	var req createTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return taskusecase.CreateTaskInput{}, apperror.New(http.StatusBadRequest, "invalid json payload")
	}
	if strings.TrimSpace(req.Title) == "" {
		return taskusecase.CreateTaskInput{}, apperror.WithFields(http.StatusBadRequest, "validation failed", map[string]string{
			"title": "is required",
		})
	}

	dueDate, appErr := parseDueDate(req.DueDate)
	if appErr != nil {
		return taskusecase.CreateTaskInput{}, appErr
	}

	return taskusecase.CreateTaskInput{
		Title:       strings.TrimSpace(req.Title),
		Description: strings.TrimSpace(req.Description),
		Status:      strings.TrimSpace(req.Status),
		Priority:    strings.TrimSpace(req.Priority),
		AssigneeID:  req.AssigneeID,
		DueDate:     dueDate,
	}, nil
}

func MapUpdateTask(c *gin.Context) (taskusecase.UpdateTaskInput, *apperror.AppError) {
	var req updateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return taskusecase.UpdateTaskInput{}, apperror.New(http.StatusBadRequest, "invalid json payload")
	}

	dueDate, appErr := parseDueDate(req.DueDate)
	if appErr != nil {
		return taskusecase.UpdateTaskInput{}, appErr
	}

	trimPointer(req.Title)
	trimPointer(req.Description)
	trimPointer(req.Status)
	trimPointer(req.Priority)
	trimPointer(req.AssigneeID)

	return taskusecase.UpdateTaskInput{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		Priority:    req.Priority,
		AssigneeID:  req.AssigneeID,
		DueDate:     dueDate,
	}, nil
}

func parseDueDate(raw *string) (*time.Time, *apperror.AppError) {
	if raw == nil {
		return nil, nil
	}

	dateText := strings.TrimSpace(*raw)
	if dateText == "" {
		return nil, nil
	}

	if parsed, err := time.Parse("2006-01-02", dateText); err == nil {
		return &parsed, nil
	}
	if parsed, err := time.Parse(time.RFC3339, dateText); err == nil {
		return &parsed, nil
	}

	return nil, apperror.New(http.StatusBadRequest, "due_date must be YYYY-MM-DD or RFC3339")
}

func trimPointer(value *string) {
	if value == nil {
		return
	}
	trimmed := strings.TrimSpace(*value)
	*value = trimmed
}
