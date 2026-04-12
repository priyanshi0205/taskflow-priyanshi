package task

import "backedn_go/model"

type Repository interface {
	ListByProject(projectID, status, assigneeID string) ([]model.Task, error)
	Create(task *model.Task) error
	FindByID(id string) (*model.Task, error)
	IsOwnedByUser(taskID, userID string) (bool, error)
	IsCreator(taskID, userID string) (bool, error)
	Update(task *model.Task) error
	Delete(id string) error
}
