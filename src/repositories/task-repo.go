package repositories

import (
	"smart-task-backend/src/db/models"
	"smart-task-backend/src/dto"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type TaskRepo interface {
	SyncTasks(tasks []dto.Task) (*gorm.DB, []models.Task)
	GetTasks(userId uint) (*gorm.DB, []models.Task)
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepo {
	return &taskRepo{db}
}

func (this_ *taskRepo) SyncTasks(tasks []dto.Task) (*gorm.DB, []models.Task) {

	var tasksModel []models.Task

	for _, task := range tasks {
		var taskModel models.Task
		if err := smapping.FillStruct(&taskModel, smapping.MapFields(&task)); err != nil {
			panic(err)
		}
		tasksModel = append(tasksModel, taskModel)
	}
	result := this_.db.Create(&tasksModel)
	return result, tasksModel
}

func (this_ *taskRepo) GetTasks(userId uint) (*gorm.DB, []models.Task) {
	var foundTasks []models.Task
	result := this_.db.Where("user_id", userId).Take(&foundTasks)

	return result, foundTasks
}
