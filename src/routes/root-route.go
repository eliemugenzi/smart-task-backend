package routes

import (
	"smart-task-backend/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RootRoute(db *gorm.DB, router *gin.Engine, logger *utils.Logger) {
	apiRouter := router.Group("/api/v1")

	// Define routers here...

	// Auth router config
	authRouter := apiRouter.Group("/auth")
	AuthRoute(db, authRouter, logger)

}
