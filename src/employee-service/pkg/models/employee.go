package models

import (
	"time"

	"github.com/lib/pq"
)

type Employee struct {
	ID             uint
	Email          string
	Password       string
	FirstName      string
	LastName       string
	Birthday       time.Time
	Education      string
	JobType        pq.StringArray `gorm:"type:varchar(64)[]"`
	RequierdSkills pq.StringArray `gorm:"type:varchar(64)[]"`
	Blocked        bool
}
