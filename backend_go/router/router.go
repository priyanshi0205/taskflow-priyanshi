package router

import (
	authcontroller "backedn_go/controller/auth"
	projectcontroller "backedn_go/controller/project"
	taskcontroller "backedn_go/controller/task"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backedn_go/internal/shared/middleware"
)

type Dependencies struct {
	AuthController    authcontroller.Controller
	ProjectController projectcontroller.Controller
	TaskController    taskcontroller.Controller
	JWTSecret         string
}

func RegisterRoutes(engine *gin.Engine, deps Dependencies) {
	engine.Use(middleware.RequestLogger(), gin.Recovery())
	engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	authMiddleware := middleware.AuthRequired(deps.JWTSecret)

	engine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "TaskFlow API running")
	})
	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	authGroup := engine.Group("/auth")
	{
		authGroup.POST("/register", deps.AuthController.Register)
		authGroup.POST("/login", deps.AuthController.Login)
	}

	projectGroup := engine.Group("/projects")
	projectGroup.Use(authMiddleware)
	{
		projectGroup.GET("", deps.ProjectController.List)
		projectGroup.POST("", deps.ProjectController.Create)
		projectGroup.GET("/:id", deps.ProjectController.GetByID)
		projectGroup.GET("/:id/stats", deps.ProjectController.Stats)
		projectGroup.PATCH("/:id", deps.ProjectController.Update)
		projectGroup.DELETE("/:id", deps.ProjectController.Delete)
	}

	taskGroup := engine.Group("/")
	taskGroup.Use(authMiddleware)
	{
		taskGroup.GET("/projects/:id/tasks", deps.TaskController.List)
		taskGroup.POST("/projects/:id/tasks", deps.TaskController.Create)
		taskGroup.PATCH("/tasks/:id", deps.TaskController.Update)
		taskGroup.DELETE("/tasks/:id", deps.TaskController.Delete)
	}
}
