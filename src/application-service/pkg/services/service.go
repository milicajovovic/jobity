package services

import (
	"application-service/pkg/models"
	"application-service/pkg/repositories"
)

func GetAll() ([]models.Application, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.Application, error) {
	return repositories.GetById(id)
}

func Apply(adId int, employeeId int) (models.Application, error) {
	return repositories.Create(adId, employeeId)
}

func GetAccepted(employeeId int) ([]models.Application, error) {
	return repositories.GetAccepted(employeeId)
}
