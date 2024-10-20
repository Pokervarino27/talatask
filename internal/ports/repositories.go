package ports

import "github.com/pokervarino27/talatask/internal/domain"

type EmployeeRepository interface {
	GetAll() ([]domain.Employee, error)
}

type TaskRepository interface {
	GetAll() ([]domain.Task, error)
}
