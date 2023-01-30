package services

import (
	"employer-service/pkg/models"
	"employer-service/pkg/repositories"
	"errors"
)

func GetAll() ([]models.Employer, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.Employer, error) {
	return repositories.GetById(id)
}

func Register(employer models.Employer) (models.Employer, error) {
	if repositories.UniqueEmail(employer.Email) {
		return repositories.Create(employer)
	}
	return models.Employer{}, errors.New("email must be unique")
}

func Login(dto models.LoginDTO) (models.Employer, error) {
	return repositories.Login(dto)
}

func Delete(id int) (models.Employer, error) {
	return repositories.Delete(id)
}
