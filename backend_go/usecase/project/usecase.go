package project

import (
	"errors"
	"net/http"

	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
	projectrepo "backedn_go/repository/project"

	"gorm.io/gorm"
)

type usecase struct {
	repo projectrepo.Repository
}

func New(repo projectrepo.Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) List(userID string) ([]model.Project, *apperror.AppError) {
	projects, err := u.repo.ListByUser(userID)
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}
	return projects, nil
}

func (u *usecase) Create(userID string, input CreateProjectInput) (*model.Project, *apperror.AppError) {
	project, err := u.repo.Create(input.Name, input.Description, userID)
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}
	return project, nil
}

func (u *usecase) GetByID(id string) (*model.Project, *apperror.AppError) {
	project, err := u.repo.FindByIDWithTasks(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.New(http.StatusNotFound, "not found")
		}
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}
	return project, nil
}

func (u *usecase) Stats(userID, id string) (*ProjectStatsOutput, *apperror.AppError) {
	project, err := u.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.New(http.StatusNotFound, "not found")
		}
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	if project.OwnerID != userID {
		return nil, apperror.New(http.StatusForbidden, "Forbidden")
	}

	byStatus, err := u.repo.CountByStatus(id)
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	byAssignee, err := u.repo.CountByAssignee(id)
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	return &ProjectStatsOutput{
		ByStatus:   byStatus,
		ByAssignee: byAssignee,
	}, nil
}

func (u *usecase) Update(userID, id string, input UpdateProjectInput) (*model.Project, *apperror.AppError) {
	project, err := u.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, apperror.New(http.StatusNotFound, "not found")
		}
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	if project.OwnerID != userID {
		return nil, apperror.New(http.StatusForbidden, "Forbidden")
	}

	if input.Name != nil {
		project.Name = *input.Name
	}
	if input.Description != nil {
		project.Description = *input.Description
	}

	if err := u.repo.Update(project); err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}

	return project, nil
}

func (u *usecase) Delete(userID, id string) *apperror.AppError {
	project, err := u.repo.FindByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return apperror.New(http.StatusNotFound, "not found")
		}
		return apperror.New(http.StatusInternalServerError, "server error")
	}

	if project.OwnerID != userID {
		return apperror.New(http.StatusForbidden, "Forbidden")
	}

	if err := u.repo.Delete(id); err != nil {
		return apperror.New(http.StatusInternalServerError, "server error")
	}

	return nil
}
