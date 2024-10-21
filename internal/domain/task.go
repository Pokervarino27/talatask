package domain

import "time"

type Task struct {
	ID             string
	Title          string
	Date           time.Time
	Duration       int //en horas
	RequiredSkills []Skill
	IsAssigned     bool
}
