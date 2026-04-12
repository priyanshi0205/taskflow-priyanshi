package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backedn_go/internal/shared/apperror"
	authusecase "backedn_go/usecase/auth"
)

func MapRegister(c *gin.Context) (authusecase.RegisterInput, *apperror.AppError) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return authusecase.RegisterInput{}, apperror.New(http.StatusBadRequest, "invalid json payload")
	}
	fields := map[string]string{}
	if strings.TrimSpace(req.Name) == "" {
		fields["name"] = "is required"
	}
	if strings.TrimSpace(req.Email) == "" {
		fields["email"] = "is required"
	}
	if strings.TrimSpace(req.Password) == "" {
		fields["password"] = "is required"
	}
	if len(fields) > 0 {
		return authusecase.RegisterInput{}, apperror.WithFields(http.StatusBadRequest, "validation failed", fields)
	}
	return authusecase.RegisterInput{
		Name:     strings.TrimSpace(req.Name),
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
	}, nil
}

func MapLogin(c *gin.Context) (authusecase.LoginInput, *apperror.AppError) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return authusecase.LoginInput{}, apperror.New(http.StatusBadRequest, "invalid json payload")
	}

	fields := map[string]string{}
	if strings.TrimSpace(req.Email) == "" {
		fields["email"] = "is required"
	}
	if strings.TrimSpace(req.Password) == "" {
		fields["password"] = "is required"
	}
	if len(fields) > 0 {
		return authusecase.LoginInput{}, apperror.WithFields(http.StatusBadRequest, "validation failed", fields)
	}

	return authusecase.LoginInput{
		Email:    strings.TrimSpace(req.Email),
		Password: req.Password,
	}, nil
}
