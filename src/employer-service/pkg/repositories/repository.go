package repositories

import (
	"employer-service/pkg/config"
	"employer-service/pkg/models"
)

func GetAll() ([]models.EmployerDTO, error) {
	var employers []models.Employer
	result := config.DB.Debug().Find(&employers)

	if result.Error != nil {
		return nil, result.Error
	}

	var dtos []models.EmployerDTO
	for _, employer := range employers {
		dtos = append(dtos, employer.EmployerToDTO())
	}
	return dtos, nil
}

func GetById(id int) (models.EmployerDTO, error) {
	var employer models.Employer
	result := config.DB.First(&employer, id)

	if result.Error != nil {
		return models.EmployerDTO{}, result.Error
	}

	return employer.EmployerToDTO(), nil
}
