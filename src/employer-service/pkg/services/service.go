package services

import (
	"employer-service/pkg/models"
	"employer-service/pkg/repositories"
)

func GetAll() ([]models.EmployerDTO, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.EmployerDTO, error) {
	return repositories.GetById(id)
}
