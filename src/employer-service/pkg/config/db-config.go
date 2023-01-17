package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"employer-service/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	db        *gorm.DB
	err       error
	employers = []models.Employer{
		{
			Email:    "foody@gmail.com",
			Password: HashPassword("foody123"),
			Name:     "Foody",
			Address:  "West Street 5",
		},
		{
			Email:    "abank@gmail.com",
			Password: HashPassword("abank321"),
			Name:     "A-Bank",
			Address:  "North Street 15",
		},
		{
			Email:    "itsoft@gmail.com",
			Password: HashPassword("itsoft123"),
			Name:     "IT Soft",
			Address:  "East Street 49",
		},
	}
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=employers port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.Migrator().DropTable("employers")
	db.AutoMigrate(&models.Employer{})
	for _, employer := range employers {
		db.Create(&employer)
	}
}