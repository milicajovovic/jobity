package services

import (
	"errors"
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
	if repositories.IsUnique(review) {
		return repositories.Create(review)
	}
	return models.Review{}, errors.New("you already left a review for this employer")
}

func Appropriate(id int) (models.Review, error) {
	return repositories.Appropriate(id)
}

func Delete(id int) (models.Review, error) {
	return repositories.Delete(id)
}
