package user

import "github.com/gin-gonic/gin"

type Controller interface {
	List(c *gin.Context)
}
