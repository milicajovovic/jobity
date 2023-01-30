package repositories

import (
	"admin-service/pkg/config"
	"admin-service/pkg/models"
	"errors"
)

func GetAll() ([]models.Admin, error) {
	var admins []models.Admin
	result := config.DB.Find(&admins)

	if result.Error != nil {
		return nil, result.Error
	}
	return admins, nil
}

func GetById(id int) (models.Admin, error) {
	var admin models.Admin
	result := config.DB.First(&admin, id)

	if result.Error != nil {
		return models.Admin{}, result.Error
	}
	return admin, nil
}

func GetByEmail(email string) (models.Admin, error) {
	var admin models.Admin
	result := config.DB.Where("email = ?", email).First(&admin)

	if result.Error != nil {
		return models.Admin{}, result.Error
	}
	return admin, nil
}

func Login(dto models.LoginDTO) (models.Admin, error) {
	admin, err := GetByEmail(dto.Email)
	if err != nil {
		return models.Admin{}, errors.New("email is not valid")
	}

	if !config.CheckPasswordHash(dto.Password, admin.Password) {
		return models.Admin{}, errors.New("password is not valid")
	}
	return admin, nil
}
