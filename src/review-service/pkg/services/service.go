package services

import (
	"review-service/pkg/models"
	"review-service/pkg/repositories"
)

func GetAll() ([]models.Review, error) {
	return repositories.GetAll()
}

func GetByEmployerId(id int) ([]models.Review, error) {
	return repositories.GetByEmployerId(id)
}

func Create(review models.Review) (models.Review, error) {
	return repositories.Create(review)
}
