package models

import (
	"time"
)

type EmployeeDTO struct {
	ID             uint
	Email          string
	FirstName      string
	LastName       string
	Birthday       time.Time
	Education      string
	JobType        []string
	Skills         []string
	ProfilePicture string
	CV             string
	Blocked        bool
	Deleted        bool
}
