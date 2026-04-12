package user

import "backedn_go/model"

type Repository interface {
	ListExceptUser(userID string) ([]model.User, error)
}
