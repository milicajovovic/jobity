package services

import (
	"ad-service/pkg/models"
	"ad-service/pkg/repositories"
)

func GetAll() ([]models.Ad, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.Ad, error) {
	return repositories.GetById(id)
}

func Search(name string, description string) ([]models.Ad, error) {
	if name == "Name" && description == "Description" {
		return repositories.GetAll()
	} else if name == "Name" && description != "Description" {
		return repositories.SearchByDescription(description)
	} else if name != "Name" && description == "Description" {
		return repositories.SearchByName(name)
	} else {
		return repositories.SearchByNameAndDescription(name, description)
	}
}
