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

func GetByEmployerId(id int) ([]models.Application, error) {
	var applications []models.Application
	result := config.DB.Where("employer_id = ? AND status = ?", id, models.Applied).Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}

func IsUnique(adId int, employeeId int) bool {
	var application models.Application
	result := config.DB.Where("ad_id = ? AND employee_id = ?", adId, employeeId).First(&application)
	return result.Error != nil
}

func Create(application models.Application) (models.Application, error) {
	application.Status = models.Applied
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

func Update(application models.Application) (models.Application, error) {
	result := config.DB.Save(&application)

	if result.Error != nil {
		return models.Application{}, result.Error
	}
	return application, nil
}

func GetInterviews(id int) ([]models.Application, error) {
	var applications []models.Application
	result := config.DB.Where("employer_id = ? AND status = ?", id, models.ScheduledInterview).Find(&applications)

	if result.Error != nil {
		return nil, result.Error
	}
	return applications, nil
}
