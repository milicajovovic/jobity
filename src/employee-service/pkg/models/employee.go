package models

import (
	"time"

	"github.com/lib/pq"
)

type Employee struct {
	ID        uint
	Email     string
	Password  string
	FirstName string
	LastName  string
	Birthday  time.Time
	Education string
	JobType   pq.StringArray `gorm:"type:varchar(64)[]"`
	Skills    pq.StringArray `gorm:"type:varchar(64)[]"`
	CV        string
	Blocked   bool
}

func (employee *Employee) EmployeeToDTO() EmployeeDTO {
	return EmployeeDTO{
		ID:        employee.ID,
		Email:     employee.Email,
		FirstName: employee.FirstName,
		LastName:  employee.LastName,
		Birthday:  employee.Birthday,
		Education: employee.Education,
		JobType:   employee.JobType,
		Skills:    employee.Skills,
		CV:        employee.CV,
		Blocked:   employee.Blocked,
	}
}
