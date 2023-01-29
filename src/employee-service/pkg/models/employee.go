package models

import (
	"time"

	"github.com/lib/pq"
)

type Employee struct {
	ID             int
	Email          string
	Password       string `json:"-"`
	FirstName      string
	LastName       string
	Birthday       time.Time
	Education      string
	JobType        pq.StringArray `gorm:"type:varchar(64)[]"`
	Skills         pq.StringArray `gorm:"type:varchar(64)[]"`
	ProfilePicture string
	CV             string
	Blocked        bool
	Deleted        bool
}
