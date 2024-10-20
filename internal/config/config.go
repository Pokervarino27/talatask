package config

import (
	"github.com/pokervarino27/talatask/internal/adapters/handlers"
	"github.com/pokervarino27/talatask/internal/usecases"
)

type AppConfig struct {
	Handler  *handlers.Handler
	Database *Database
}

func NewAppConfig() *AppConfig {
	db := NewDatabase()
	taskAssignmentService := usecases.NewTaskAssignmentService(db.EmployeeRepo, db.TaskRepo)
	handler := handlers.NewHandler(taskAssignmentService)

	return &AppConfig{
		Handler:  handler,
		Database: db,
	}
}
