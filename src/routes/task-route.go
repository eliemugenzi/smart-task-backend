package routes

import (
	"smart-task-backend/src/controllers"
	"smart-task-backend/src/middlewares"
	"smart-task-backend/src/repositories"
	"smart-task-backend/src/services"
	"smart-task-backend/src/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaskRoute(db *gorm.DB, taskRouter *gin.RouterGroup, logger *utils.Logger) {
	var (
		jwtService     services.JwtService        = services.NewJwtService()
		taskRepository repositories.TaskRepo      = repositories.NewTaskRepo(db)
		taskService    services.TaskService       = services.TaskService(taskRepository)
		taskController controllers.TaskController = controllers.NewTaskController(taskService, jwtService, logger)
	)

	taskRouter.GET("/", middlewares.AuthorizeJWT(), taskController.GetTasks)
	taskRouter.POST("/sync", middlewares.AuthorizeJWT(), taskController.SyncTasks)
}
