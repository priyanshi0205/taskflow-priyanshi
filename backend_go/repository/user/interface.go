package user

import "backedn_go/model"

type Repository interface {
	ListAllUser() ([]model.User, error)
}
