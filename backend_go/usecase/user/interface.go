package user

import (
	"backedn_go/internal/shared/apperror"
	"backedn_go/model"
)

type Usecase interface {
	ListExceptUser(userID string) ([]model.User, *apperror.AppError)
}
