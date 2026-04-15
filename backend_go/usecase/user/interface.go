package user

import (
	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
)

type Usecase interface {
	ListAllUser() ([]model.User, *apperror.AppError)
}
