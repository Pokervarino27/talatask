package usecases

import (
	"fmt"
	"sort"
	"time"

	"github.com/pokervarino27/talatask/internal/domain"
	"github.com/pokervarino27/talatask/internal/ports"
)

type TaskAssignmentService struct {
	employeeRepo ports.EmployeeRepository
	taskRepo     ports.TaskRepository
}

func NewTaskAssignmentService(er ports.EmployeeRepository, tr ports.TaskRepository) *TaskAssignmentService {
	return &TaskAssignmentService{
		employeeRepo: er,
		taskRepo:     tr,
	}
}

func (s *TaskAssignmentService) AssignTask() ([]domain.Assignment, error) {
	employees, err := s.employeeRepo.GetAll()
	if err != nil {
		return nil, err
	}
	fmt.Println("employees:", employees)
	tasks, err := s.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}
	fmt.Println("tasks", tasks)
	assignments := make([]domain.Assignment, 0)

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Date.Before(tasks[j].Date)
	})

	for _, task := range tasks {
		assigned := false
		for _, employee := range employees {
			if canAssign(employee, task) {
				assignments = append(assignments, domain.Assignment{
					EmployeeID: employee.ID,
					TaskID:     task.ID,
				})
				assigned = true
				break
			}
		}
		if !assigned {
			fmt.Println("cannot assigned task")
		}
	}
	return assignments, nil
}

func (s *TaskAssignmentService) GenerateReport(date time.Time) (*domain.AssignmentReport, error) {
	employees, err := s.employeeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	tasks, err := s.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	report := &domain.AssignmentReport{
		Date:      date,
		Employees: make([]domain.EmployeeReport, 0),
	}

	for _, employee := range employees {
		employeeReport := domain.EmployeeReport{
			ID:             employee.ID,
			Name:           employee.Name,
			AssignedTasks:  make([]domain.TaskReport, 0),
			TotalHours:     0,
			UsedSkills:     make([]domain.Skill, 0),
			RemainingHours: employee.AvailabilityHours,
		}

		for _, task := range tasks {
			if task.Date.Equal(date) && canAssign(employee, task) {
				taskReport := domain.TaskReport{
					ID:       task.ID,
					Title:    task.Title,
					Duration: task.Duration,
					Skills:   task.RequiredSkills,
				}
				employeeReport.AssignedTasks = append(employeeReport.AssignedTasks, taskReport)
				employeeReport.TotalHours += task.Duration
				employeeReport.RemainingHours -= task.Duration
				employeeReport.UsedSkills = append(employeeReport.UsedSkills, task.RequiredSkills...)
			}
		}

		report.Employees = append(report.Employees, employeeReport)
	}

	sort.Slice(report.Employees, func(i, j int) bool {
		return report.Employees[i].Name < report.Employees[j].Name
	})

	return report, nil

}

func canAssign(employee domain.Employee, task domain.Task) bool {
	availableOnDate := false
	for _, date := range employee.AvailabilityDays {
		if date.Equal(task.Date) {
			availableOnDate = true
			break
		}
	}
	if !availableOnDate {
		return false
	}

	for _, requiredSkill := range task.RequiredSkills {
		skill := false
		for _, employeeSkill := range employee.Skills {
			if employeeSkill == requiredSkill {
				skill = true
				break
			}
		}
		if !skill {
			return false
		}
	}
	return true
}
