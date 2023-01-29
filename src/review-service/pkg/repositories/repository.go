package repositories

import (
	"review-service/pkg/config"
	"review-service/pkg/models"
)

func GetAll() ([]models.Review, error) {
	var reviews []models.Review
	result := config.DB.Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}
	return reviews, nil
}

func GetByEmployerId(id int) ([]models.Review, error) {
	var reviews []models.Review
	result := config.DB.Where("employer_id = ?", id).Find(&reviews)

	if result.Error != nil {
		return nil, result.Error
	}
	return reviews, nil
}

func IsUnique(review models.Review) bool {
	var existingReview models.Review
	result := config.DB.Where("employer_id = ? AND employee_id = ?", review.EmployerID, review.EmployeeID).First(&existingReview)
	return result.Error != nil
}

func Create(review models.Review) (models.Review, error) {
	result := config.DB.Create(&review)

	if result.Error != nil {
		return models.Review{}, result.Error
	}
	return review, nil
}
