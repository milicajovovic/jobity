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
	db        *gorm.DB
	err       error
	employees = []models.Employee{
		{
			FirstName:      "John",
			LastName:       "Doe",
			Email:          "john@doe.com",
			Password:       HashPassword("johndoe123"),
			Birthday:       time.Date(1988, 12, 12, 0, 0, 0, 0, time.Local),
			Education:      "Bachelor - Culinary Arts",
			JobType:        []string{"Kitchen"},
			RequierdSkills: []string{"Organisation", "Creativity"},
			Blocked:        false,
		},
	}
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=employees port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.Migrator().DropTable("employees")
	db.AutoMigrate(&models.Employee{})
	for _, employee := range employees {
		db.Create(&employee)
	}
}
