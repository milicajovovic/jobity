package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"admin-service/pkg/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	DB     *gorm.DB
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

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func InitDB() {
	dsn := "host=localhost user=postgres password=postgres dbname=admins port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	DB.Migrator().DropTable("admins")
	DB.AutoMigrate(&models.Admin{})
	for _, admin := range admins {
		DB.Create(&admin)
	}
}
