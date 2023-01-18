package repositories

import (
	"ad-service/pkg/config"
	"ad-service/pkg/models"
	"strings"
)

func GetAll() ([]models.Ad, error) {
	var ads []models.Ad
	result := config.DB.Find(&ads)

	if result.Error != nil {
		return nil, result.Error
	}
	return ads, nil
}

func GetById(id int) (models.Ad, error) {
	var ad models.Ad
	result := config.DB.First(&ad, id)

	if result.Error != nil {
		return models.Ad{}, result.Error
	}
	return ad, nil
}

func SearchByName(name string) ([]models.Ad, error) {
	var ads []models.Ad
	result := config.DB.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(name)+"%").Find(&ads)

	if result.Error != nil {
		return nil, result.Error
	}
	return ads, nil
}

func SearchByDescription(description string) ([]models.Ad, error) {
	var ads []models.Ad
	result := config.DB.Where("LOWER(description) LIKE ?", "%"+strings.ToLower(description)+"%").Find(&ads)

	if result.Error != nil {
		return nil, result.Error
	}
	return ads, nil
}

func SearchByNameAndDescription(name string, description string) ([]models.Ad, error) {
	var ads []models.Ad
	result := config.DB.Where("LOWER(name) LIKE ? AND LOWER(description) LIKE ?", "%"+strings.ToLower(name)+"%", "%"+strings.ToLower(description)+"%").Find(&ads)

	if result.Error != nil {
		return nil, result.Error
	}
	return ads, nil
}
