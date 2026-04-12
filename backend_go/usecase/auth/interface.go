package auth

import "backedn_go/internal/shared/apperror"

type Usecase interface {
	Register(input RegisterInput) (*RegisterOutput, *apperror.AppError)
	Login(input LoginInput) (*LoginOutput, *apperror.AppError)
}
