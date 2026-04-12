package task

import "github.com/gin-gonic/gin"

type Controller interface {
	List(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}
