package project

import "backedn_go/model"

type Repository interface {
	ListByUser(userID string) ([]model.Project, error)
	Create(name, description, ownerID string) (*model.Project, error)
	FindByID(id string) (*model.Project, error)
	FindByIDWithTasks(id string) (*model.Project, error)
	CountByStatus(projectID string) (map[string]int64, error)
	CountByAssignee(projectID string) (map[string]int64, error)
	Update(project *model.Project) error
	Delete(id string) error
}
