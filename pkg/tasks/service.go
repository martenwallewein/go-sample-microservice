package tasks

import (
	"sync"

	"github.com/martenwallewein/go-sample-microservice/dbs"
	"github.com/martenwallewein/go-sample-microservice/models"
	"gorm.io/gorm"
)

type TaskService struct {
	db *gorm.DB
}

var taskService *TaskService
var initOnce sync.Once

func GetService() *TaskService {
	initOnce.Do(initService)
	return taskService
}

func initService() {
	db := dbs.GetDB()
	taskService = NewTaskService(db)
}

func NewTaskService(db *gorm.DB) *TaskService {
	return &TaskService{
		db: db,
	}
}

/**
 * Task operations
**/

func (dl *TaskService) GetAllTasks() ([]models.Task, error) {
	var tasks []models.Task
	result := dl.db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (dl *TaskService) CreateTask(task *models.Task) error {
	result := dl.db.Create(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *TaskService) EditTask(id uint, task models.Task) error {
	dbTask := models.Task{}
	dbTask.ID = id
	result := dl.db.Model(&dbTask).Updates(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (dl *TaskService) DeleteTask(id uint) error {
	dbTask := models.Task{}
	dbTask.ID = id
	result := dl.db.Delete(&dbTask)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
