package repositories

import (
	"errors"
	"sync"

	"github.com/pokervarino27/talatask/internal/domain"
)

type EmployeeRepositoryImpl struct {
	employees map[string]domain.Employee
	mutex     sync.RWMutex
}

func NewEmployeeRepository() *EmployeeRepositoryImpl {
	return &EmployeeRepositoryImpl{
		employees: make(map[string]domain.Employee),
	}
}

func (r *EmployeeRepositoryImpl) GetAll() ([]domain.Employee, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	employees := make([]domain.Employee, 0, len(r.employees))
	for _, empl := range r.employees {
		employees = append(employees, empl)
	}
	return employees, nil
}

func (r *EmployeeRepositoryImpl) Create(employee domain.Employee) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.employees[employee.ID]; exists {
		return errors.New("employee already exists")
	}

	r.employees[employee.ID] = employee
	return nil
}
