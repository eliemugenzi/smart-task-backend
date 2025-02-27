package main

import (
	"net/http"
	"os"
	"smart-task-backend/src/db/config"
	"smart-task-backend/src/middlewares"
	"smart-task-backend/src/routes"
	"smart-task-backend/src/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

var db *gorm.DB = config.Configure()

func main() {
	godotenv.Load()
	defer config.CloseConnection(db)
	logger := utils.NewLogger()
	router := gin.Default()
	zLogger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	router.Use(middlewares.RequestLogger(&zLogger))

	routes.RootRoute(db, router, logger)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.Run(":" + os.Getenv("APP_PORT"))
}
