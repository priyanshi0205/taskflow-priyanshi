package user

import (
	"net/http"

	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
	userrepo "backedn_go/repository/user"
)

type usecase struct {
	repo userrepo.Repository
}

func New(repo userrepo.Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) ListAllUser() ([]model.User, *apperror.AppError) {
	users, err := u.repo.ListAllUser()
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "server error")
	}
	return users, nil
}
