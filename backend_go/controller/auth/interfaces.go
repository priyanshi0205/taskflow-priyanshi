package auth

import "github.com/gin-gonic/gin"

type Controller interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}
