package repositories

import (
	"application-service/pkg/config"
	"application-service/pkg/models"
)

func GetAll() ([]models.Application, error) {
	var applications []models.Application
	result := config.DB.Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}

func GetById(id int) (models.Application, error) {
	var application models.Application
	result := config.DB.First(&application, id)

	if result.Error != nil {
		return models.Application{}, result.Error
	}
	return application, nil
}

func Create(adId int, employeeId int) (models.Application, error) {
	application := models.Application{
		AdID:       adId,
		EmployeeID: employeeId,
		Status:     models.Applied,
	}
	result := config.DB.Create(&application)

	if result.Error != nil {
		return models.Application{}, result.Error
	}
	return application, nil
}

func GetAccepted(employeeId int) ([]models.Application, error) {
	var applications []models.Application
	result := config.DB.Where("employee_id = ? AND status = ?", employeeId, models.PassedInterview).Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}
