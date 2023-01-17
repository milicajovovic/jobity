package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"grade-service/pkg/models"
)

var (
	db     *gorm.DB
	err    error
	grades = []models.Grade{
		{
			EmployerID:    1,
			Grade:         5,
			Comment:       "Great working conditions!",
			Inappropriate: false,
		},
		{
			EmployerID:    1,
			Grade:         5,
			Comment:       "Perfect place to learn!",
			Inappropriate: false,
		},
		{
			EmployerID:    2,
			Grade:         3,
			Comment:       "Not worth it...",
			Inappropriate: false,
		},
	}
)

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=grades port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.Migrator().DropTable("grades")
	db.AutoMigrate(&models.Grade{})
	for _, grade := range grades {
		db.Create(&grade)
	}
}
