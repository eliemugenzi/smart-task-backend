package routes

import (
	"smart-task-backend/src/controllers"
	"smart-task-backend/src/repositories"
	"smart-task-backend/src/services"
	"smart-task-backend/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoute(db *gorm.DB, authRouter *gin.RouterGroup, logger *utils.Logger) {
	var (
		jwtService     services.JwtService        = services.NewJwtService()
		authRepository repositories.AuthRepo      = repositories.NewAuthRepo(db)
		authService    services.AuthService       = services.NewAuthService(authRepository)
		authController controllers.AuthController = controllers.NewAuthController(authService, jwtService, logger)
	)

	authRouter.POST("/login", authController.Login)
	authRouter.POST("/signup", authController.Register)
	authRouter.POST("/token/verify", authController.VerifyToken)
	authRouter.POST("/token/refresh", authController.RefreshToken)
	authRouter.GET("/users", authController.GetUsers)
}
