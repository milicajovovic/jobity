package repositories

import (
	"employee-service/pkg/config"
	"employee-service/pkg/models"
	"errors"
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

func Update(employee models.Employee) (models.EmployeeDTO, error) {
	employee.Password = config.HashPassword(employee.Password)
	result := config.DB.Save(&employee)

	if result.Error != nil {
		return models.EmployeeDTO{}, result.Error
	}
	return employee.EmployeeToDTO(), nil
}

func GetByEmail(email string) (models.Employee, error) {
	var employee models.Employee
	result := config.DB.Where("email = ?", email).First(&employee)

	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func Login(dto models.LoginDTO) (models.EmployeeDTO, error) {
	employee, err := GetByEmail(dto.Email)
	if err != nil {
		return models.EmployeeDTO{}, errors.New("email is not valid")
	}

	if !config.CheckPasswordHash(dto.Password, employee.Password) {
		return models.EmployeeDTO{}, errors.New("password is not valid")
	}
	return employee.EmployeeToDTO(), nil
}
