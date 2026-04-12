package task

import (
	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
)

type Usecase interface {
	List(projectID, status, assigneeID string) ([]model.Task, *apperror.AppError)
	Create(projectID string, input CreateTaskInput) (*model.Task, *apperror.AppError)
	Update(userID, taskID string, input UpdateTaskInput) (*model.Task, *apperror.AppError)
	Delete(userID, taskID string) *apperror.AppError
}
