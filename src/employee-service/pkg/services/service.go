package services

import (
	"employee-service/pkg/models"
	"employee-service/pkg/repositories"
	"errors"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"
)

func GetAll() ([]models.EmployeeDTO, error) {
	return repositories.GetAll()
}

func GetById(id int) (models.EmployeeDTO, error) {
	return repositories.GetById(id)
}

func RegisterForm(employee models.Employee) (models.EmployeeDTO, error) {
	if repositories.UniqueEmail(employee.Email) {
		return repositories.Create(employee)
	}
	return models.EmployeeDTO{}, errors.New("email must be unique")
}

func RegisterPdf(dto models.RegisterDTO) (models.EmployeeDTO, error) {
	content, err := ReadPdf("data/" + dto.PdfPath)
	if err != nil {
		return models.EmployeeDTO{}, err
	}

	employee := PdfToEmployee(content)
	employee.CV = dto.PdfPath
	employee.Password = dto.Passwod
	if repositories.UniqueEmail(employee.Email) {
		return repositories.Create(employee)
	}

	return models.EmployeeDTO{}, errors.New("email must be unique")
}

func ReadPdf(path string) ([]string, error) {
	f, r, err := pdf.Open(path)

	if err != nil {
		return nil, err
	}

	var content []string
	rows, _ := r.Page(1).GetTextByRow()
	for _, row := range rows {
		word := ""
		for _, char := range row.Content {
			word += char.S
		}
		content = append(content, word)
	}

	f.Close()
	return content, nil
}

func PdfToEmployee(content []string) models.Employee {
	var employee models.Employee
	for _, row := range content {
		rowData := strings.Split(row, ": ")
		switch strings.ToLower(rowData[0]) {
		case "first name":
			employee.FirstName = rowData[1]
		case "last name":
			employee.LastName = rowData[1]
		case "email", "e-mail":
			employee.Email = rowData[1]
		case "birthday":
			employee.Birthday, _ = time.Parse("DD-MM-YYYY", rowData[1])
		case "education":
			employee.Education = rowData[1]
		case "job type":
			employee.JobType = strings.Split(rowData[1], ", ")
		case "skills":
			employee.Skills = strings.Split(rowData[1], ", ")
		}
	}
	return employee
}

func Update(employee models.Employee) (models.EmployeeDTO, error) {
	if repositories.UniqueEmail(employee.Email) {
		return repositories.Update(employee)
	}
	return models.EmployeeDTO{}, errors.New("email must be unique")
}

func Login(dto models.LoginDTO) (models.EmployeeDTO, error) {
	return repositories.Login(dto)
}
