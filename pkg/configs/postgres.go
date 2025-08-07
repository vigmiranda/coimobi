package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/vigmiranda/coimobi-service/internal/property/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	ginMode := os.Getenv("GIN_MODE")

	if ginMode == "release" {
		fmt.Println("Skipping .ENV loader")

		return
	}

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error trying to load .env")
		return
	}
}

func ConnectDatabase() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_DATABASE")
	dbPort := os.Getenv("DB_PORT")

	connectParameters := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbDatabase, dbPort)

	DB, err = gorm.Open(postgres.Open(connectParameters))
	if err != nil {
		log.Panic("Error connecting to configs")
	}
	DB.AutoMigrate(&model.Property{})

}
