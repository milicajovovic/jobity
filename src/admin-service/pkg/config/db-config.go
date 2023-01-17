package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"admin-service/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	db     *gorm.DB
	err    error
	admins = []models.Admin{
		{
			Email:    "admin@gmail.com",
			Password: HashPassword("admin"),
		},
		{
			Email:    "smart@gmail.com",
			Password: HashPassword("smart123"),
		},
	}
)

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=admins port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.Migrator().DropTable("admins")
	db.AutoMigrate(&models.Admin{})
	for _, admin := range admins {
		db.Create(&admin)
	}
}
