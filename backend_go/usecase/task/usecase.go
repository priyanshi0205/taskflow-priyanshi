package task

import (
	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
	taskrepo "backedn_go/repository/task"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type usecase struct {
	repo taskrepo.Repository
}

func New(repo taskrepo.Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) List(projectID, status, assigneeID string) ([]model.Task, *apperror.AppError) {
	tasks, err := u.repo.ListByProject(projectID, status, assigneeID)
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}
	return tasks, nil
}

func (u *usecase) Create(projectID string, input CreateTaskInput) (*model.Task, *apperror.AppError) {
	status := normalizeOrDefault(input.Status, "todo")
	if _, ok := validStatuses[status]; !ok {
		return nil, apperror.New(http.StatusBadRequest, "invalid status")
	}

	priority := normalizeOrDefault(input.Priority, "medium")
	if _, ok := validPriorities[priority]; !ok {
		return nil, apperror.New(http.StatusBadRequest, "invalid priority")
	}

	taskModel := &model.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      status,
		Priority:    priority,
		ProjectID:   projectID,
		CreatedByID: input.CreatedByID,
		AssigneeID:  normalizeOptionalString(input.AssigneeID),
		DueDate:     input.DueDate,
	}

	if err := u.repo.Create(taskModel); err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}
	return taskModel, nil
}

func (u *usecase) canManageTask(taskID, userID string) (bool, *apperror.AppError) {
	isOwner, err := u.repo.IsOwnedByUser(taskID, userID)
	if err != nil {
		return false, apperror.New(http.StatusInternalServerError, "server error")
	}
	if isOwner {
		return true, nil
	}

	isCreator, err := u.repo.IsCreator(taskID, userID)
	if err != nil {
		return false, apperror.New(http.StatusInternalServerError, "server error")
	}

	return isCreator, nil
}

func (u *usecase) Update(userID, taskID string, input UpdateTaskInput) (*model.Task, *apperror.AppError) {
	taskModel, err := u.repo.FindByID(taskID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.New(http.StatusNotFound, "not found")
		}
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	canManage, appErr := u.canManageTask(taskID, userID)
	if appErr != nil {
		return nil, appErr
	}
	if !canManage {
		return nil, apperror.New(http.StatusForbidden, "Forbidden")
	}

	if input.Title != nil {
		taskModel.Title = *input.Title
	}
	if input.Description != nil {
		taskModel.Description = *input.Description
	}
	if input.Status != nil {
		status := strings.TrimSpace(*input.Status)
		if _, ok := validStatuses[status]; !ok {
			return nil, apperror.New(http.StatusBadRequest, "invalid status")
		}
		taskModel.Status = status
	}
	if input.Priority != nil {
		priority := strings.TrimSpace(*input.Priority)
		if _, ok := validPriorities[priority]; !ok {
			return nil, apperror.New(http.StatusBadRequest, "invalid priority")
		}
		taskModel.Priority = priority
	}
	if input.AssigneeID != nil {
		assigneeID := strings.TrimSpace(*input.AssigneeID)
		taskModel.AssigneeID = &assigneeID
	}
	if input.DueDate != nil {
		taskModel.DueDate = input.DueDate
	}

	if err := u.repo.Update(taskModel); err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	return taskModel, nil
}

func (u *usecase) Delete(userID, taskID string) *apperror.AppError {
	_, err := u.repo.FindByID(taskID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.New(http.StatusNotFound, "not found")
		}
		return apperror.New(http.StatusInternalServerError, "server error")
	}

	canManage, appErr := u.canManageTask(taskID, userID)
	if appErr != nil {
		return appErr
	}
	if !canManage {
		return apperror.New(http.StatusForbidden, "Forbidden")
	}

	if err := u.repo.Delete(taskID); err != nil {
		return apperror.New(http.StatusInternalServerError, "server error")
	}
	return nil
}
