package ports

import (
	"time"

	"github.com/pokervarino27/talatask/internal/domain"
)

type TaskAssignmentService interface {
	AssignTask() ([]domain.Assignment, error)
	GenerateReport(date time.Time) (*domain.AssignmentReport, error)
}
