package models

import (
	"employer-service/pkg/models"

	"github.com/lib/pq"
)

type AdDTO struct {
	ID             int
	Name           string
	Employer       models.EmployerDTO
	Description    string
	JobType        pq.StringArray
	RequierdSkills pq.StringArray
}
