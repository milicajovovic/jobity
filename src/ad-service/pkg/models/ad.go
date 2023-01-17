package models

import "github.com/lib/pq"

type Ad struct {
	ID             uint
	Name           string
	EmployerID     uint
	Description    string
	JobType        pq.StringArray `gorm:"type:varchar(64)[]"`
	RequierdSkills pq.StringArray `gorm:"type:varchar(64)[]"`
}
