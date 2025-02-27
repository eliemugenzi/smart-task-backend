package config

import (
	"fmt"
	"log"
	"os"
	"smart-task-backend/src/db/models"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DB *gorm.DB
}

func Configure() *gorm.DB {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			Colorful:                  true,
		},
	)
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Having trouble loading .env file...")
	}
	databaseHost := os.Getenv("DATABASE_HOST")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable Timezone=Africa/Kigali", databaseHost, databaseUser, databasePassword, databaseName, databasePort)

	DB, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: true,
			},
		),
		&gorm.Config{
			Logger: newLogger,
		},
	)

	if err != nil {
		log.Fatal("Having trouble connecting to the database: ")

		return DB
	}

	DB.AutoMigrate(&models.User{}, &models.Task{})

	return DB
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close the connection with Database")
	}

	dbSQL.Close()
}
