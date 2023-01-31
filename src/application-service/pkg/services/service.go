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

func GetByEmployerId(id int) ([]models.Application, error) {
	return repositories.GetByEmployerId(id)
}

func Apply(application models.Application) (models.Application, error) {
	if repositories.IsUnique(application.AdID, application.EmployeeID) {
		return repositories.Create(application)
	}
	return models.Application{}, errors.New("you already applied for this job")
}

func GetAccepted(employeeId int) ([]models.Application, error) {
	return repositories.GetAccepted(employeeId)
}

func Update(application models.Application) (models.Application, error) {
	return repositories.Update(application)
}

func GetInterviews(id int) ([]models.Application, error) {
	return repositories.GetInterviews(id)
}
