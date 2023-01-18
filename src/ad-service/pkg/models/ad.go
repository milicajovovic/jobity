package models

import (
	"employer-service/pkg/models"

	"github.com/lib/pq"
)

type Ad struct {
	ID             int
	Name           string
	EmployerID     int
	Description    string
	JobType        pq.StringArray `gorm:"type:varchar(64)[]"`
	RequierdSkills pq.StringArray `gorm:"type:varchar(64)[]"`
}

func (ad *Ad) AdToDTO(employer models.EmployerDTO) AdDTO {
	return AdDTO{
		ID:             ad.ID,
		Name:           ad.Name,
		Employer:       employer,
		Description:    ad.Description,
		JobType:        ad.JobType,
		RequierdSkills: ad.RequierdSkills,
	}
}
