package repositories

import (
	"employer-service/pkg/config"
	"employer-service/pkg/models"
	"errors"
)

func GetAll() ([]models.Employer, error) {
	var employers []models.Employer
	result := config.DB.Find(&employers)

	if result.Error != nil {
		return nil, result.Error
	}
	return employers, nil
}

func GetById(id int) (models.Employer, error) {
	var employer models.Employer
	result := config.DB.First(&employer, id)

	if result.Error != nil {
		return models.Employer{}, result.Error
	}
	return employer, nil
}

func UniqueEmail(email string) bool {
	var employer models.Employer
	result := config.DB.Where("email = ?", email).First(&employer)
	return result.Error != nil
}

func Create(employer models.Employer) (models.Employer, error) {
	employer.Password = config.HashPassword(employer.Password)
	result := config.DB.Create(&employer)

	if result.Error != nil {
		return models.Employer{}, result.Error
	}
	return employer, nil
}

func GetByEmail(email string) (models.Employer, error) {
	var employer models.Employer
	result := config.DB.Where("email = ?", email).First(&employer)

	if result.Error != nil {
		return models.Employer{}, result.Error
	}
	return employer, nil
}

func Login(dto models.LoginDTO) (models.Employer, error) {
	employer, err := GetByEmail(dto.Email)
	if err != nil {
		return models.Employer{}, errors.New("email is not valid")
	}

	if !config.CheckPasswordHash(dto.Password, employer.Password) {
		return models.Employer{}, errors.New("password is not valid")
	}
	return employer, nil
}

func Delete(id int) (models.Employer, error) {
	var employer models.Employer
	result := config.DB.First(&employer, id).Update("deleted", true)

	if result.Error != nil {
		return models.Employer{}, result.Error
	}
	return employer, nil
}
