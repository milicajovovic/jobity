package services

import (
	"application-service/pkg/models"
	"application-service/pkg/repositories"
	"errors"
)

func GetAll() ([]models.Application, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.Application, error) {
	return repositories.GetById(id)
}

func Apply(adId int, employeeId int) (models.Application, error) {
	if repositories.IsUnique(adId, employeeId) {
		return repositories.Create(adId, employeeId)
	}
	return models.Application{}, errors.New("you already applied for this job")
}

func GetAccepted(employeeId int) ([]models.Application, error) {
	return repositories.GetAccepted(employeeId)
}
