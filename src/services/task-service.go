package services

import (
	"smart-task-backend/src/db/models"
	"smart-task-backend/src/dto"
	"smart-task-backend/src/repositories"

	"gorm.io/gorm"
)

type TaskService interface {
	SyncTasks(tasks []dto.Task) (*gorm.DB, []models.Task)
	GetTasks(userId uint) (*gorm.DB, []models.Task)
}

type taskService struct {
	taskRepo repositories.TaskRepo
}

func NewTaskService(taskRepo repositories.TaskRepo) *taskService {
	return &taskService{
		taskRepo,
	}
}

func (this_ *taskService) SyncTasks(tasks []dto.Task) (*gorm.DB, []models.Task) {
	return this_.taskRepo.SyncTasks(tasks)
}

func (this_ *taskService) GetTasks(userId uint) []models.Task {
	_, results := this_.taskRepo.GetTasks(userId)

	return results
}
