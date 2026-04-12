package app

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	authcontroller "backedn_go/controller/auth"
	projectcontroller "backedn_go/controller/project"
	taskcontroller "backedn_go/controller/task"
	"backedn_go/internal/shared/config"
	"backedn_go/internal/shared/database"
	"backedn_go/model"
	authrepo "backedn_go/repository/auth"
	projectrepo "backedn_go/repository/project"
	taskrepo "backedn_go/repository/task"
	"backedn_go/router"
	authusecase "backedn_go/usecase/auth"
	projectusecase "backedn_go/usecase/project"
	taskusecase "backedn_go/usecase/task"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	db, err := database.NewPostgres(cfg.DatabaseURL)
	if err != nil {
		return err
	}

	if err := database.PrepareSchema(db, &model.User{}, &model.Project{}, &model.Task{}); err != nil {
		return err
	}

	authRepository := authrepo.New(db)
	projectRepository := projectrepo.New(db)
	taskRepository := taskrepo.New(db)

	authUsecase := authusecase.New(authRepository, cfg.JWTSecret)
	projectUsecase := projectusecase.New(projectRepository)
	taskUsecase := taskusecase.New(taskRepository)

	authController := authcontroller.New(authUsecase)
	projectController := projectcontroller.New(projectUsecase)
	taskController := taskcontroller.New(taskUsecase)

	engine := gin.New()

	router.RegisterRoutes(engine, router.Dependencies{
		AuthController:    authController,
		ProjectController: projectController,
		TaskController:    taskController,
		JWTSecret:         cfg.JWTSecret,
	})

	server := &http.Server{
		Addr:              ":" + cfg.Port,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		slog.Info("backend server starting", "port", cfg.Port)
		if listenErr := server.ListenAndServe(); listenErr != nil && !errors.Is(listenErr, http.ErrServerClosed) {
			slog.Error("server failed", "error", listenErr)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	slog.Info("shutdown signal received, shutting down gracefully")
	return server.Shutdown(shutdownCtx)
}
