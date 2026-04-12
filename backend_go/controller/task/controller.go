package task

import (
	"net/http"

	"github.com/gin-gonic/gin"

	taskInputMapper "backedn_go/input_mapper/task"
	"backedn_go/internal/shared/middleware"
	"backedn_go/internal/shared/response"
	taskOutputMapper "backedn_go/output_mapper/task"
	taskUsecase "backedn_go/usecase/task"
)

type controller struct {
	usecase taskUsecase.Usecase
}

func New(usecase taskUsecase.Usecase) Controller {
	return &controller{usecase: usecase}
}

func (ctl *controller) List(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	projectID, appErr := taskInputMapper.MapProjectID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map project id for task list", "user_id", userID, "project_id", projectID)
		return
	}

	status, assignee := taskInputMapper.MapListFilters(c)

	tasks, appErr := ctl.usecase.List(projectID, status, assignee)
	if appErr != nil {
		response.ErrorWithContext(
			c,
			appErr,
			"failed to list tasks",
			"user_id", userID,
			"project_id", projectID,
			"status_filter", status,
			"assignee_filter", assignee,
		)
		return
	}

	c.JSON(http.StatusOK, taskOutputMapper.MapList(tasks))
}

func (ctl *controller) Create(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	projectID, appErr := taskInputMapper.MapProjectID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map project id for task create", "user_id", userID, "project_id", projectID)
		return
	}

	input, appErr := taskInputMapper.MapCreateTask(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map create-task input", "user_id", userID, "project_id", projectID)
		return
	}

	input.CreatedByID = userID

	task, appErr := ctl.usecase.Create(projectID, input)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to create task", "user_id", userID, "project_id", projectID)
		return
	}

	c.JSON(http.StatusOK, taskOutputMapper.MapSingle(*task))
}

func (ctl *controller) Update(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	taskID, appErr := taskInputMapper.MapTaskID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map task id for update", "user_id", userID, "task_id", taskID)
		return
	}
	input, appErr := taskInputMapper.MapUpdateTask(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map update-task input", "user_id", userID, "task_id", taskID)
		return
	}

	task, appErr := ctl.usecase.Update(userID, taskID, input)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to update task", "user_id", userID, "task_id", taskID)
		return
	}

	c.JSON(http.StatusOK, taskOutputMapper.MapSingle(*task))
}

func (ctl *controller) Delete(c *gin.Context) {
	userID := middleware.UserIDFromContext(c)

	taskID, appErr := taskInputMapper.MapTaskID(c)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to map task id for delete", "user_id", userID, "task_id", taskID)
		return
	}

	appErr = ctl.usecase.Delete(userID, taskID)
	if appErr != nil {
		response.ErrorWithContext(c, appErr, "failed to delete task", "user_id", userID, "task_id", taskID)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
