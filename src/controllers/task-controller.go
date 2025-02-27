package controllers

import (
	"net/http"
	"smart-task-backend/src/dto"
	"smart-task-backend/src/services"
	"smart-task-backend/src/utils"

	"github.com/gin-gonic/gin"
)

type TaskController interface {
	SyncTasks(context *gin.Context)
	GetTasks(context *gin.Context)
}

type taskController struct {
	taskService services.TaskService
	jwtService  services.JwtService
	logger      *utils.Logger
}

func NewTaskController(taskService services.TaskService, jwtService services.JwtService, logger *utils.Logger) *taskController {
	return &taskController{
		taskService,
		jwtService,
		logger,
	}
}

func (this_ *taskController) SyncTasks(context *gin.Context) {
	var tasksDto []dto.Task
	if err := context.ShouldBindJSON(&tasksDto); err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		this_.logger.Error(err.Error())
		return
	}

	this_.taskService.SyncTasks(tasksDto)
	context.JSON(http.StatusOK, utils.GetResponse(http.StatusOK, "Data in sync...", nil))
	return
}

func (this_ *taskController) GetTasks(context *gin.Context) {
	userId, _ := context.Get("user_id")
	_, foundTasks := this_.taskService.GetTasks(userId.(uint))
	context.JSON(http.StatusOK, utils.GetResponse(http.StatusOK, "", foundTasks))
	return
}
