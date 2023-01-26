package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"application-service/pkg/models"
)

var (
	DB           *gorm.DB
	err          error
	applications = []models.Application{
		{
			AdID:       1,
			EmployeeID: 1,
			Status:     models.PassedInterview,
		},
	}
)

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=applications port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB.Migrator().DropTable("applications")
	DB.AutoMigrate(&models.Application{})
	for _, ad := range applications {
		DB.Create(&ad)
	}
}
