package models

import (
	"time"

	"github.com/lib/pq"
)

type Ad struct {
	ID             int
	Name           string
	EmployerID     int
	Description    string
	Posted         time.Time
	JobType        pq.StringArray `gorm:"type:varchar(64)[]"`
	RequiredSkills pq.StringArray `gorm:"type:varchar(64)[]"`
	Deleted        bool
}
