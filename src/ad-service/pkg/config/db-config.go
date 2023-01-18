package config

import (
	"log"

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
			JobType:        []string{"Kitchen", "Full time"},
			RequierdSkills: []string{"Organisation", "Multitasking", "Creativity"},
		},
		{
			Name:           "Food Prep",
			EmployerID:     1,
			Description:    "Preparing ingredients for dishes to help the kitchen staff.",
			JobType:        []string{"Kitchen", "Part time", "Morning shift"},
			RequierdSkills: []string{"Organisation", "Multitasking", "Team work"},
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
