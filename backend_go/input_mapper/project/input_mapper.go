package project

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backedn_go/internal/shared/apperror"
	projectusecase "backedn_go/usecase/project"
)

func MapProjectID(c *gin.Context) (string, *apperror.AppError) {
	id := strings.TrimSpace(c.Param(Id))
	if id == "" {
		return "", apperror.New(http.StatusBadRequest, "project id is required")
	}
	return id, nil
}

func MapCreateProject(c *gin.Context) (projectusecase.CreateProjectInput, *apperror.AppError) {
	var req createProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return projectusecase.CreateProjectInput{}, apperror.New(http.StatusBadRequest, "invalid json payload")
	}

	if strings.TrimSpace(req.Name) == "" {
		return projectusecase.CreateProjectInput{}, apperror.WithFields(
			http.StatusBadRequest,
			"validation failed",
			map[string]string{"name": "is required"},
		)

	}
	return projectusecase.CreateProjectInput{
		Name:        strings.TrimSpace(req.Name),
		Description: strings.TrimSpace(req.Description),
	}, nil
}

func MapUpdateProject(c *gin.Context) (projectusecase.UpdateProjectInput, *apperror.AppError) {
	var req updateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return projectusecase.UpdateProjectInput{}, apperror.New(http.StatusBadRequest, "invalid json payload")
	}

	if req.Name != nil {
		trimmed := strings.TrimSpace(*req.Name)
		if trimmed == "" {
			return projectusecase.UpdateProjectInput{}, apperror.WithFields(
				http.StatusBadRequest,
				"validation failed",
				map[string]string{"name": "cannot be empty"},
			)
		}
		req.Name = &trimmed
	}
	if req.Description != nil {
		trimmed := strings.TrimSpace(*req.Description)
		req.Description = &trimmed
	}

	return projectusecase.UpdateProjectInput{
		Name:        req.Name,
		Description: req.Description,
	}, nil
}
