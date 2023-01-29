package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"employee-service/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	DB        *gorm.DB
	err       error
	employees = []models.Employee{
		{
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john@doe.com",
			Password:  HashPassword("johndoe123"),
			Birthday:  time.Date(1988, 12, 12, 0, 0, 0, 0, time.Local),
			Education: "Bachelor - Culinary Arts",
			JobType:   []string{"Kitchen"},
			Skills:    []string{"Organisation", "Creativity"},
		},
		{
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@doe.com",
			Password:  HashPassword("janedoe123"),
			Birthday:  time.Date(1995, 06, 04, 0, 0, 0, 0, time.Local),
			Education: "Bachelor - Banking and Finance",
			JobType:   []string{"Administration", "Economy"},
			Skills:    []string{"Accounting"},
		},
	}
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=employees port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB.Migrator().DropTable("employees")
	DB.AutoMigrate(&models.Employee{})
	for _, employee := range employees {
		DB.Create(&employee)
	}
}
