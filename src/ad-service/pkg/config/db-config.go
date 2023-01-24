package config

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ad-service/pkg/models"
)

var (
	DB  *gorm.DB
	err error
	ads = []models.Ad{
		{
			Name:           "Chef",
			EmployerID:     1,
			Description:    "Culinary professional trained in all aspects of food preparation.",
			Posted:         time.Date(2022, 11, 12, 0, 0, 0, 0, time.Local),
			JobType:        []string{"Kitchen", "Full time"},
			RequiredSkills: []string{"Organisation", "Multitasking", "Creativity"},
		},
		{
			Name:           "Food Prep",
			EmployerID:     1,
			Description:    "Preparing ingredients for dishes to help the kitchen staff.",
			Posted:         time.Date(2023, 01, 12, 0, 0, 0, 0, time.Local),
			JobType:        []string{"Kitchen", "Part time", "Morning shift"},
			RequiredSkills: []string{"Organisation", "Multitasking", "Team work"},
		},
	}
)

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=ads port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB.Migrator().DropTable("ads")
	DB.AutoMigrate(&models.Ad{})
	for _, ad := range ads {
		DB.Create(&ad)
	}
}
