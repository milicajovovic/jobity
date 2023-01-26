package models

import (
	"time"
)

type ApplicationStatus int

const (
	Applied ApplicationStatus = iota
	Rejected
	ScheduledInterview
	PassedInterview
)

type Application struct {
	ID         int
	AdID       int
	EmployeeID int
	Status     ApplicationStatus
	Interview  time.Time
	Deleted    bool
}
