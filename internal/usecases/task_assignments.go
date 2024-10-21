package usecases

import (
	"sort"
	"time"

	"github.com/pokervarino27/talatask/internal/domain"
	"github.com/pokervarino27/talatask/internal/infraestructure/logger"
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
	logger.Infof("employees:", employees)
	tasks, err := s.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}
	logger.Infof("tasks", tasks)
	assignments := make([]domain.Assignment, 0)

	availabilityMap := make(map[time.Time][]domain.Employee)
	skillsMap := make(map[domain.Skill][]domain.Employee)

	for _, employee := range employees {
		for _, date := range employee.AvailabilityDays {
			availabilityMap[date] = append(availabilityMap[date], employee)
		}

		for _, skill := range employee.Skills {
			skillsMap[skill] = append(skillsMap[skill], employee)
		}
	}

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Date.Before(tasks[j].Date)
	})

	for _, task := range tasks {
		if task.IsAssigned {
			continue
		}

		availableEmployees := availabilityMap[task.Date]
		if len(availableEmployees) == 0 {
			logger.Warn("No hay empleados disponibles")
			continue
		}

		selectedEmployees := filterBySkills(availableEmployees, task.RequiredSkills, skillsMap)

		if len(selectedEmployees) == 0 {
			logger.Info("No se puede asignar tarea")
			continue
		}

		assignments = append(assignments, domain.Assignment{
			EmployeeID: selectedEmployees[0].ID,
			TaskID:     task.ID,
		})
		task.IsAssigned = true
	}
	return assignments, nil
}

func (s *TaskAssignmentService) GenerateReport(date time.Time) (*domain.AssignmentReport, error) {
	formattedDate := date.Format("2006-01-02")
	logger.Info(formattedDate)
	employees, err := s.employeeRepo.GetAll()
	if err != nil {
		return nil, err
	}

	tasks, err := s.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	logger.Infof("tasks", tasks)

	report := &domain.AssignmentReport{
		Date:      date,
		Employees: make([]domain.EmployeeReport, 0),
	}

	taskMap := make(map[string][]domain.Task)
	for _, task := range tasks {
		taskDate := task.Date.Format("2006-01-02")
		taskMap[taskDate] = append(taskMap[taskDate], task)
	}

	tasksForDate := taskMap[formattedDate]
	logger.Infof("tasksForDate", tasksForDate)

	for _, employee := range employees {
		employeeReport := domain.EmployeeReport{
			ID:             employee.ID,
			Name:           employee.Name,
			AssignedTasks:  make([]domain.TaskReport, 0),
			TotalHours:     0,
			UsedSkills:     make([]domain.Skill, 0),
			RemainingHours: employee.AvailabilityHours,
		}

		for _, task := range tasksForDate {
			if canAssign(employee, task) {
				taskReport := domain.TaskReport{
					ID:       task.ID,
					Title:    task.Title,
					Duration: task.Duration,
					Skills:   task.RequiredSkills,
				}
				employeeReport.AssignedTasks = append(employeeReport.AssignedTasks, taskReport)
				employeeReport.TotalHours += task.Duration
				employeeReport.RemainingHours -= task.Duration
				employeeReport.UsedSkills = append(employeeReport.UsedSkills, taskReport.Skills...)
			}
		}

		report.Employees = append(report.Employees, employeeReport)
	}
	return report, nil
}

func canAssign(employee domain.Employee, task domain.Task) bool {
	if task.IsAssigned {
		return false
	}
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

func filterBySkills(availableEmployees []domain.Employee, requiredSkills []domain.Skill, skillsMap map[domain.Skill][]domain.Employee) []domain.Employee {
	employeesSkillCount := make(map[string]int)
	for _, skill := range requiredSkills {
		for _, employee := range skillsMap[skill] {
			employeesSkillCount[employee.ID]++
		}
	}

	selectedEmployees := make([]domain.Employee, 0)
	for _, employee := range availableEmployees {
		if employeesSkillCount[employee.ID] == len(requiredSkills) {
			selectedEmployees = append(selectedEmployees, employee)
		}
	}
	return selectedEmployees
}
