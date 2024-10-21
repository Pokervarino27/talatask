package usecases

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/pokervarino27/talatask/internal/domain"
)

type MockEmployeeRepository struct {
	employees []domain.Employee
}

func (m *MockEmployeeRepository) GetAll() ([]domain.Employee, error) {
	return m.employees, nil
}

type MockTaskRepository struct {
	tasks []domain.Task
}

func (m *MockTaskRepository) GetAll() ([]domain.Task, error) {
	return m.tasks, nil
}

func TestAssignTasks(t *testing.T) {
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.AddDate(0, 0, 1)
	dayAfterTomorrow := today.AddDate(0, 0, 2)
	employees := []domain.Employee{
		{
			ID:                "1",
			Name:              "Diego Comihual",
			Skills:            []domain.Skill{"programming", "design"},
			AvailabilityHours: 8,
			AvailabilityDays:  []time.Time{today, tomorrow},
		}, {
			ID:                "2",
			Name:              "Daniel Cortes",
			Skills:            []domain.Skill{"testing", "analysis"},
			AvailabilityHours: 6,
			AvailabilityDays:  []time.Time{today, tomorrow, dayAfterTomorrow},
		},
	}

	tasks := []domain.Task{
		{
			ID:             "1",
			Title:          "Develop feature",
			Date:           today,
			Duration:       4,
			RequiredSkills: []domain.Skill{"programming"},
		},
		{
			ID:             "2",
			Title:          "Test feature",
			Date:           tomorrow,
			Duration:       3,
			RequiredSkills: []domain.Skill{"testing"},
		},
	}

	mockEmployeeRepo := &MockEmployeeRepository{employees: employees}
	mockTaskRepo := &MockTaskRepository{tasks: tasks}

	service := NewTaskAssignmentService(mockEmployeeRepo, mockTaskRepo)

	assignments, err := service.AssignTask()

	fmt.Println(assignments)

	if err != nil {
		t.Errorf("Error al asignar tareas: %v", err)
	}

	if len(assignments) != len(tasks) {
		t.Errorf("Se esperaban %d asignaciones, se obtuvieron %d", len(tasks), len(assignments))
	}

	expectedAssignments := []domain.Assignment{
		{EmployeeID: "1", TaskID: "1"},
		{EmployeeID: "2", TaskID: "2"},
	}

	if !reflect.DeepEqual(assignments, expectedAssignments) {
		t.Errorf("Las asignaciones no son las esperadas. Resultado: %v, Expected: %v", assignments, expectedAssignments)
	}
}

func TestGenerateReport(t *testing.T) {
	now := time.Now()
	employees := []domain.Employee{
		{
			ID:                "1",
			Name:              "Diego Comihual",
			Skills:            []domain.Skill{"programming", "design"},
			AvailabilityHours: 8,
			AvailabilityDays:  []time.Time{now, now.AddDate(0, 0, 1)},
		},
		{
			ID:                "2",
			Name:              "Daniel Cortes",
			Skills:            []domain.Skill{"testing", "analysis"},
			AvailabilityHours: 6,
			AvailabilityDays:  []time.Time{now, now.AddDate(0, 0, 1)},
		},
	}

	tasks := []domain.Task{
		{
			ID:             "1",
			Title:          "Develop feature",
			Date:           now,
			Duration:       4,
			RequiredSkills: []domain.Skill{"programming"},
		},
		{
			ID:             "2",
			Title:          "Test feature",
			Date:           now,
			Duration:       3,
			RequiredSkills: []domain.Skill{"testing"},
		},
	}

	mockEmployeeRepo := &MockEmployeeRepository{employees: employees}
	mockTaskRepo := &MockTaskRepository{tasks: tasks}

	service := NewTaskAssignmentService(mockEmployeeRepo, mockTaskRepo)

	report, err := service.GenerateReport(now)

	if err != nil {
		t.Errorf("Error al generar el reporte: %v", err)
	}

	if !report.Date.Equal(now) {
		t.Errorf("La fecha del reporte no es correcta. Obtenido")
	}

	if len(report.Employees) != len(employees) {
		t.Errorf("employees expected: %d, Result: %d", len(employees), len(report.Employees))
	}

	for _, emplReport := range report.Employees {
		switch emplReport.ID {
		case "1":
			if len(emplReport.AssignedTasks) != 1 {
				t.Errorf("expected: 1, ID: 1, result:%d", len(emplReport.AssignedTasks))
			}
			if emplReport.TotalHours != 4 {
				t.Errorf("expected: 4, ID: 1, result:%d", emplReport.TotalHours)
			}
			if emplReport.RemainingHours != 4 {
				t.Errorf("expected: 4, ID: 1, result:%d", emplReport.RemainingHours)
			}
		case "2":
			if len(emplReport.AssignedTasks) != 1 {
				t.Errorf("expected: 1, ID: 2, result:%d", len(emplReport.AssignedTasks))
			}
			if emplReport.TotalHours != 3 {
				t.Errorf("expected: 3, ID: 2, result:%d", emplReport.TotalHours)
			}
			if emplReport.RemainingHours != 3 {
				t.Errorf("expected: 3, ID: 2, result:%d", emplReport.RemainingHours)
			}
		default:
			t.Errorf("unexpected result %s", emplReport.ID)
		}
	}
}
