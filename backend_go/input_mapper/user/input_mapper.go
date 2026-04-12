package user

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"backedn_go/internal/shared/apperror"
	"backedn_go/internal/shared/middleware"
)

func MapRequesterUserID(c *gin.Context) (string, *apperror.AppError) {
	requesterID := strings.TrimSpace(middleware.UserIDFromContext(c))
	if requesterID == "" {
		return "", apperror.New(http.StatusUnauthorized, "Unauthorized")
	}
	return requesterID, nil
}
