package repositories

import (
	"grade-service/pkg/config"
	"grade-service/pkg/models"
)

func GetAll() ([]models.Grade, error) {
	var grades []models.Grade
	result := config.DB.Find(&grades)

	if result.Error != nil {
		return nil, result.Error
	}
	return grades, nil
}

func GetByEmployerId(id int) ([]models.Grade, error) {
	var grades []models.Grade
	result := config.DB.Where("employer_id = ?", id).Find(&grades)

	if result.Error != nil {
		return nil, result.Error
	}
	return grades, nil
}
