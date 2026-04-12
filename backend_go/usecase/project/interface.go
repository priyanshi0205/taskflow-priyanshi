package project

import (
	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
)

type Usecase interface {
	List(userID string) ([]model.Project, *apperror.AppError)
	Create(userID string, input CreateProjectInput) (*model.Project, *apperror.AppError)
	GetByID(id string) (*model.Project, *apperror.AppError)
	Stats(userID, id string) (*ProjectStatsOutput, *apperror.AppError)
	Update(userID, id string, input UpdateProjectInput) (*model.Project, *apperror.AppError)
	Delete(userID, id string) *apperror.AppError
}
