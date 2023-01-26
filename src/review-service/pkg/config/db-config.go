package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"review-service/pkg/models"
)

var (
	DB      *gorm.DB
	err     error
	reviews = []models.Review{
		{
			EmployerID:    1,
			EmployeeID:    1,
			Grade:         5,
			Comment:       "Great working conditions!",
			Inappropriate: false,
		},
		{
			EmployerID:    2,
			EmployeeID:    2,
			Grade:         3,
			Comment:       "Not worth it...",
			Inappropriate: false,
		},
	}
)

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=reviews port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB.Migrator().DropTable("reviews")
	DB.AutoMigrate(&models.Review{})
	for _, review := range reviews {
		DB.Create(&review)
	}
}
