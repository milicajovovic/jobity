package services

import (
	"admin-service/pkg/models"
	"admin-service/pkg/repositories"
)

func GetAll() ([]models.Admin, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.Admin, error) {
	return repositories.GetById(id)
}

func Login(dto models.LoginDTO) (models.Admin, error) {
	return repositories.Login(dto)
}
