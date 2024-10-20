package domain

import "time"

type TaskReport struct {
	ID       string
	Title    string
	Duration int
	Skills   []Skill
}

type EmployeeReport struct {
	ID             string
	Name           string
	AssignedTasks  []TaskReport
	TotalHours     int
	UsedSkills     []Skill
	RemainingHours int
}

type AssignmentReport struct {
	Date      time.Time
	Employees []EmployeeReport
}
