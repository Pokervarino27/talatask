package domain

import "time"

type Employee struct {
	ID                string
	Name              string
	Skills            []Skill
	AvailabilityHours int
	AvailabilityDays  []time.Time
}

type Skill string
