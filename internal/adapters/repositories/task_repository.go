package repositories

import (
	"errors"
	"sync"

	"github.com/pokervarino27/talatask/internal/domain"
)

type TaskRepositoryImpl struct {
	tasks map[string]domain.Task
	mutex sync.RWMutex
}

func NewTaskRespository() *TaskRepositoryImpl {
	return &TaskRepositoryImpl{
		tasks: make(map[string]domain.Task),
	}
}

func (r *TaskRepositoryImpl) GetAll() ([]domain.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	tasks := make([]domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepositoryImpl) Create(task domain.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[task.ID]; exists {
		return errors.New("task already exists")
	}
	r.tasks[task.ID] = task
	return nil
}
