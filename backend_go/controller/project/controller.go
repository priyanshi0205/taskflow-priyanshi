package project

import (
	"net/http"

	"github.com/gin-gonic/gin"

	projectInputMapper "backedn_go/input_mapper/project"
	"backedn_go/internal/shared/middleware"
	"backedn_go/internal/shared/response"
	projectOutputMapper "backedn_go/output_mapper/project"
	projectUsecase "backedn_go/usecase/project"
)

type controller struct {
	usecase projectUsecase.Usecase
}

func New(usecase projectUsecase.Usecase) Controller {
	return &controller{usecase: usecase}
}

func (ctl *controller) List(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	projects, appErr := ctl.usecase.List(userID)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to list projects", "user_id", userID)
		return
	}

	c.JSON(http.StatusOK, projectOutputMapper.MapList(projects))
}

func (ctl *controller) Create(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	input, appErr := projectInputMapper.MapCreateProject(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map create-project input", "user_id", userID)
		return
	}

	project, appErr := ctl.usecase.Create(userID, input)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to create project", "user_id", userID)
		return
	}

	c.JSON(http.StatusOK, projectOutputMapper.MapSingle(*project))
}

func (ctl *controller) GetByID(c *gin.Context) {
	projectID, appErr := projectInputMapper.MapProjectID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map project id", "project_id", projectID)
		return
	}

	project, appErr := ctl.usecase.GetByID(projectID)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to fetch project", "project_id", projectID)
		return
	}

	c.JSON(http.StatusOK, projectOutputMapper.MapWithTasks(*project))
}

func (ctl *controller) Stats(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	projectID, appErr := projectInputMapper.MapProjectID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map project id for stats", "user_id", userID, "project_id", projectID)
		return
	}

	stats, appErr := ctl.usecase.Stats(userID, projectID)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to fetch project stats", "user_id", userID, "project_id", projectID)
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (ctl *controller) Update(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	projectID, appErr := projectInputMapper.MapProjectID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map project id for update", "user_id", userID, "project_id", projectID)
		return
	}

	input, appErr := projectInputMapper.MapUpdateProject(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map update-project input", "user_id", userID, "project_id", projectID)
		return
	}

	project, appErr := ctl.usecase.Update(userID, projectID, input)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to update project", "user_id", userID, "project_id", projectID)
		return
	}

	c.JSON(http.StatusOK, projectOutputMapper.MapSingle(*project))
}

func (ctl *controller) Delete(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	projectID, appErr := projectInputMapper.MapProjectID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map project id for delete", "user_id", userID, "project_id", projectID)
		return
	}

	appErr = ctl.usecase.Delete(userID, projectID)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to delete project", "user_id", userID, "project_id", projectID)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Project deleted"})
}
