package services

import (
	"grade-service/pkg/models"
	"grade-service/pkg/repositories"
)

func GetAll() ([]models.Grade, error) {
	return repositories.GetAll()
}

func GetByEmployerId(id int) ([]models.Grade, error) {
	return repositories.GetByEmployerId(id)
}
