package config

import (
	"time"

	"github.com/pokervarino27/talatask/internal/adapters/repositories"
	"github.com/pokervarino27/talatask/internal/domain"
)

type Database struct {
	EmployeeRepo *repositories.EmployeeRepositoryImpl
	TaskRepo     *repositories.TaskRepositoryImpl
}

func NewDatabase() *Database {
	db := &Database{
		EmployeeRepo: repositories.NewEmployeeRepository(),
		TaskRepo:     repositories.NewTaskRespository(),
	}

	db.seedData()

	return db

}

func (db *Database) seedData() {
	today := time.Now().Truncate(24 * time.Hour)
	tomorrow := today.AddDate(0, 0, 1)
	dayAfterTomorrow := today.AddDate(0, 0, 2)

	employees := []domain.Employee{
		{
			ID:   "1",
			Name: "Diego Comihual",
			Skills: []domain.Skill{
				"programming",
				"design",
				"testing",
			},
			AvailabilityHours: 8,
			AvailabilityDays:  []time.Time{today, tomorrow, dayAfterTomorrow},
		},
		{
			ID:   "2",
			Name: "Daniel Cortes",
			Skills: []domain.Skill{
				"analysis",
				"project management",
				"design",
			},
			AvailabilityHours: 6,
			AvailabilityDays:  []time.Time{today, tomorrow, dayAfterTomorrow},
		},
		{
			ID:   "3",
			Name: "Juan Valdes",
			Skills: []domain.Skill{
				"programming",
				"testing",
				"devops",
			},
			AvailabilityHours: 7,
			AvailabilityDays:  []time.Time{today, tomorrow, dayAfterTomorrow},
		},
	}

	for _, emp := range employees {
		err := db.EmployeeRepo.Create(emp)
		if err != nil {
			panic(err)
		}
	}

	tasks := []domain.Task{
		{
			ID:       "1",
			Title:    "Develop new feature",
			Date:     today,
			Duration: 6,
			RequiredSkills: []domain.Skill{
				"programming",
				"design",
			},
		},
		{
			ID:       "2",
			Title:    "Design user interface",
			Date:     tomorrow,
			Duration: 4,
			RequiredSkills: []domain.Skill{
				"design",
			},
		},
		{
			ID:       "3",
			Title:    "Test new feature",
			Date:     dayAfterTomorrow,
			Duration: 5,
			RequiredSkills: []domain.Skill{
				"testing",
				"programming",
			},
		},
		{
			ID:       "4",
			Title:    "Project planning",
			Date:     today,
			Duration: 3,
			RequiredSkills: []domain.Skill{
				"project management",
				"analysis",
			},
		},
		{
			ID:       "5",
			Title:    "Deploy application",
			Date:     dayAfterTomorrow,
			Duration: 4,
			RequiredSkills: []domain.Skill{
				"devops",
				"programming",
			},
		},
	}

	for _, task := range tasks {
		err := db.TaskRepo.Create(task)
		if err != nil {
			panic(err)
		}
	}
}
