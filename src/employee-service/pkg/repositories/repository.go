package repositories

import (
	"employee-service/pkg/config"
	"employee-service/pkg/models"
	"errors"
)

func GetAll() ([]models.Employee, error) {
	var employees []models.Employee
	result := config.DB.Find(&employees)

	if result.Error != nil {
		return nil, result.Error
	}
	return employees, nil
}

func GetById(id int) (models.Employee, error) {
	var employee models.Employee
	result := config.DB.First(&employee, id)

	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func UniqueEmail(email string) bool {
	var employee models.Employee
	result := config.DB.Where("email = ?", email).First(&employee)
	return result.Error != nil
}

func Create(employee models.Employee) (models.Employee, error) {
	employee.Password = config.HashPassword(employee.Password)
	result := config.DB.Create(&employee)

	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func Update(employee models.Employee) (models.Employee, error) {
	employee.Password = config.HashPassword(employee.Password)
	result := config.DB.Save(&employee)

	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func GetByEmail(email string) (models.Employee, error) {
	var employee models.Employee
	result := config.DB.Where("email = ?", email).First(&employee)

	if result.Error != nil {
		return models.Employee{}, result.Error
	}
	return employee, nil
}

func Login(dto models.LoginDTO) (models.Employee, error) {
	employee, err := GetByEmail(dto.Email)
	if err != nil {
		return models.Employee{}, errors.New("email is not valid")
	}

	if !config.CheckPasswordHash(dto.Password, employee.Password) {
		return models.Employee{}, errors.New("password is not valid")
	}
	return employee, nil
}
