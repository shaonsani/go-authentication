package config

import (
	"fmt"
	"os"
	"strconv"

	"go-authentication/model"

	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, defaulting to system environment variables")
	}

	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		fmt.Println("Invalid DATABASE_PORT, defaulting to 5432")
		port = 5432
	}

	var (
		host     = os.Getenv("DATABASE_HOST")
		user     = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		dbName   = os.Getenv("DATABASE_NAME")
	)

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	fmt.Println("Database connected successfully!")

	// Automatically migrate the user model
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate the database", err)
	}

	log.Println("Database migrated successfully")

	return db
}
