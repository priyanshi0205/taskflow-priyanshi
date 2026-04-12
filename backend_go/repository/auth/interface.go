package auth

import "backedn_go/model"

type Repository interface {
	CreateUser(name, email, hashedPassword string) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}
