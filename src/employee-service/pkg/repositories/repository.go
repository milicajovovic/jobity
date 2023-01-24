package repositories

import (
	"employee-service/pkg/config"
	"employee-service/pkg/models"
)

func GetAll() ([]models.EmployeeDTO, error) {
	var employees []models.Employee
	result := config.DB.Find(&employees)

	if result.Error != nil {
		return nil, result.Error
	}

	var dtos []models.EmployeeDTO
	for _, employer := range employees {
		dtos = append(dtos, employer.EmployeeToDTO())
	}
	return dtos, nil
}

func GetById(id int) (models.EmployeeDTO, error) {
	var employee models.Employee
	result := config.DB.First(&employee, id)

	if result.Error != nil {
		return models.EmployeeDTO{}, result.Error
	}
	return employee.EmployeeToDTO(), nil
}

func UniqueEmail(email string) bool {
	var employee models.Employee
	result := config.DB.Where("email = ?", email).First(&employee)
	return result.Error != nil
}

func Create(employee models.Employee) (models.EmployeeDTO, error) {
	employee.Password = config.HashPassword(employee.Password)
	result := config.DB.Create(&employee)

	if result.Error != nil {
		return models.EmployeeDTO{}, result.Error
	}
	return employee.EmployeeToDTO(), nil
}
