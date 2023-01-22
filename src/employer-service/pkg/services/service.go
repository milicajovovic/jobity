package services

import (
	"employer-service/pkg/models"
	"employer-service/pkg/repositories"
	"errors"
)

func GetAll() ([]models.EmployerDTO, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.EmployerDTO, error) {
	return repositories.GetById(id)
}

func Register(employer models.Employer) (models.EmployerDTO, error) {
	if repositories.UniqueEmail(employer.Email) {
		return repositories.Create(employer)
	}
	return models.EmployerDTO{}, errors.New("email must be unique")
}
